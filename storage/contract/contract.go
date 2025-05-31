package contract

import (
	"app/storage/model"
	"fmt"
	"time"
)

// Storage interface
type Storage interface {
	// Tasks operations
	GetTasks(*GetTasksParams) (*GetTasksResult, error)
	CreateTask(*model.Task) (*model.Task, error)
	UpdateTask(*UpdateTaskParams) (*model.Task, error)
	DeleteTask(id uint) error

	// Dates operations
	GetDates(*GetDatesParams) (*GetDatesResult, error)
}

type (
	GetTasksParams struct {
		Offset uint
		Limit  uint
		Date   time.Time
	}

	GetTasksResult struct {
		Tasks []*model.Task
		Total uint
	}

	GetDatesParams struct {
		Offset uint
		Limit  uint
		Date   time.Time
	}

	GetDatesResult struct {
		Dates []time.Time
		Total uint
	}

	UpdateTaskParams struct {
		ID   uint
		Task *model.Task
	}
)

// TaskNotFoundError represents an error when a task is not found in the storage.

type TaskNotFoundError struct {
	ID uint
}

func NewTaskNotFoundError(id uint) *TaskNotFoundError {
	return &TaskNotFoundError{ID: id}
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("task with id %d not found", e.ID)
}
