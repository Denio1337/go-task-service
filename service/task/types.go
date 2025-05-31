package task

import (
	"app/storage"
	"errors"
	"time"
)

type (
	// Task structure
	Task struct {
		ID          uint
		Name        string
		Description string
		StartTime   time.Time
		EndTime     time.Time
	}

	// Parameters for getting tasks
	GetTasksParams struct {
		Page     uint
		PageSize uint
		Date     time.Time
	}

	// Result for getting tasks
	GetTasksResult struct {
		Tasks      []*Task
		TotalPages uint
	}

	// Parameters for getting task dates
	GetDatesParams struct {
		Page     uint
		PageSize uint
		Date     time.Time
	}

	// Result for getting task dates
	GetDatesResult struct {
		Dates      []time.Time
		TotalPages uint
	}
)

// Errors

var (
	ErrExceedsPageCount = errors.New("current page exceeds total pages")
)

type TaskNotFoundError = storage.TaskNotFoundError
