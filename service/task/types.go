package task

import "time"

type (
	// Parameters for getting tasks
	GetTasksParams struct {
		Page     uint
		PageSize uint
		Date     time.Time
	}

	// Task structure
	Task struct {
		ID          uint
		Name        string
		Description string
		StartTime   time.Time
		EndTime     time.Time
	}
)
