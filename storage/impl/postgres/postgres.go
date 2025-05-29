package postgres

import (
	"app/config"
	"app/storage/contract"
	"app/storage/model"
	"fmt"
	"strconv"

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

func (s *PostgresStorage) GetTasks(params *contract.GetTasksParams) ([]*model.Task, error) {
	var tasks []*model.Task

	result := s.db.
		Offset(int(params.Offset)).
		Limit(int(params.Limit)).
		Where("start_time >= ?", params.Date).
		Order("start_time ASC").
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
