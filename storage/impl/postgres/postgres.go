package postgres

import (
	"app/config"
	"app/storage/contract"
	"app/storage/model"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage implementation structure
type PostgresStorage struct {
	db *gorm.DB
}

// Create new PostgreSQL storage implementation
func New() (contract.Storage, error) {
	// Parse port from environment
	p := config.Get(config.EnvDBPort)
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	// Define data source name
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Get(config.EnvDBHost),
		port,
		config.Get(config.EnvDBUser),
		config.Get(config.EnvDBPassword),
		config.Get(config.EnvDBName),
	)

	// Try to connect with default gorm config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schemas to database
	db.AutoMigrate(&model.Task{})

	return &PostgresStorage{db: db}, nil
}

// Interface

func (s *PostgresStorage) GetTasks(params *contract.GetTasksParams) (*contract.GetTasksResult, error) {
	var tasks []*model.Task
	var total int64

	// Count total tasks matching the filter
	countResult := s.db.Model(&model.Task{}).Where("start_time >= ?", params.Date).Count(&total)
	if countResult.Error != nil {
		return nil, countResult.Error
	}

	result := s.db.
		Where("start_time >= ?", params.Date).
		Order("start_time ASC").
		Offset(int(params.Offset)).
		Limit(int(params.Limit)).
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return &contract.GetTasksResult{
		Tasks: tasks,
		Total: uint(total),
	}, nil
}

func (s *PostgresStorage) CreateTask(task *model.Task) (*model.Task, error) {
	result := s.db.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (s *PostgresStorage) UpdateTask(params *contract.UpdateTaskParams) (*model.Task, error) {
	// Ensure the task exists
	var existing model.Task
	if err := s.db.First(&existing, params.ID).Error; err != nil {
		return nil, contract.NewTaskNotFoundError(params.ID)
	}

	// Update only the fields that are set in the input task
	result := s.db.Model(&existing).Updates(params.Task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &existing, nil
}

func (s *PostgresStorage) DeleteTask(id uint) error {
	result := s.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return contract.NewTaskNotFoundError(id)
	}
	return nil
}

func (s *PostgresStorage) GetDates(params *contract.GetDatesParams) (*contract.GetDatesResult, error) {
	var dates []time.Time
	var total int64

	// Count total unique dates (without pagination)
	countQuery := s.db.
		Model(&model.Task{}).
		Where("start_time >= ?", params.Date).
		Select("DISTINCT DATE(start_time) AS date")

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	// Get unique dates with pagination
	rows, err := s.db.
		Model(&model.Task{}).
		Where("start_time IS NOT NULL AND start_time >= ?", params.Date).
		Select("DISTINCT DATE(start_time) AS date").
		Order("date ASC").
		Offset(int(params.Offset)).
		Limit(int(params.Limit)).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var date time.Time
		if err := rows.Scan(&date); err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}

	return &contract.GetDatesResult{
		Dates: dates,
		Total: uint(total),
	}, nil
}
