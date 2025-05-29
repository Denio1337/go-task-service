package contract

import (
	"app/storage/model"
	"time"
)

// Storage interface
type Storage interface {
	// Tasks operations
	GetTasks(*GetTasksParams) ([]*model.Task, error)
}

type GetTasksParams struct {
	Offset uint
	Limit  uint
	Date   time.Time
}
