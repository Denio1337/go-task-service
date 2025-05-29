package task

import "time"

type (
	// DTO for GET /tasks
	GetDTO struct {
		Page     int `query:"page" validate:"required,min=1"`              // Page number, must be >= 1
		PageSize int `query:"page_size" validate:"required,min=1,max=100"` // Number of items per page, must be between 1 and 100
	}

	// Response for GET /tasks
	Response struct {
		Tasks []*Task `json:"tasks"` // List of tasks
		Meta  Meta    `json:"meta"`  // Pagination metadata
	}

	// Task structure
	Task struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
	}

	// Metadata for pagination in GET /tasks response
	Meta struct {
		Page       int `json:"page"`        // Current page number
		PageSize   int `json:"page_size"`   // Number of items per page
		Total      int `json:"total"`       // Total number of items
		TotalPages int `json:"total_pages"` // Total number of pages
	}
)

const (
	PageName     = "page"      // Query parameter for page number
	PageSizeName = "page_size" // Query parameter for page size

	DefaultPage     = "1"  // Default page number if not specified
	DefaultPageSize = "10" // Default page size if not specified
)
