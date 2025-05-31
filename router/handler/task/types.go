package task

import (
	cerror "app/router/types/error"
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	// Task structure
	Task struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
	}

	// DTO for GET /tasks
	GetTasksDTO struct {
		Page     uint      `query:"page" validate:"required,min=1"`              // Page number, must be >= 1
		PageSize uint      `query:"page_size" validate:"required,min=1,max=100"` // Number of items per page, must be between 1 and 100
		Date     time.Time `query:"date"`                                        // Date to filter tasks, optional
	}

	// DTO for POST /tasks
	AddTaskDTO struct {
		Name        string    `json:"name" validate:"required,min=3,max=100"` // Task name, must be between 3 and 100 characters
		Description string    `json:"description" validate:"max=255"`         // Task description, max 255 characters
		StartTime   time.Time `json:"start_time" validate:"omitempty"`        // Start time of the task, must be provided
		EndTime     time.Time `json:"end_time"`                               // End time of the task, optional
	}

	// DTO for PATCH /tasks/{id}
	UpdateTaskDTO struct {
		Name        string    `json:"name" validate:"omitempty,min=3,max=100"` // Task name, must be between 3 and 100 characters
		Description string    `json:"description" validate:"max=255"`          // Task description, max 255 characters
		StartTime   time.Time `json:"start_time"`                              // Start time of the task, must be provided
		EndTime     time.Time `json:"end_time"`                                // End time of the task, optional
	}

	// DTO for GET /tasks/dates
	GetDatesDTO = GetTasksDTO

	// Metadata for pagination in GET /tasks response
	Meta struct {
		Page       uint `json:"page"`        // Current page number
		PageSize   uint `json:"page_size"`   // Number of items per page
		TotalPages uint `json:"total_pages"` // Total number of pages available
	}

	// Response for GET /tasks
	TasksResponse struct {
		Tasks []*Task `json:"tasks"` // List of tasks
		Meta  Meta    `json:"meta"`  // Pagination metadata
	}

	// Response for GET /tasks/dates
	TaskDatesResponse struct {
		Dates []time.Time `json:"dates"` // List of dates from tasks
		Meta  Meta        `json:"meta"`  // Pagination metadata
	}

	// Parameters for pagination
	PaginationParams struct {
		Page     uint
		PageSize uint
	}
)

const (
	PageName     = "page"      // Query parameter name for page number
	PageSizeName = "page_size" // Query parameter name for page size
	DateName     = "date"      // Query parameter name for date
	IdName       = "id"        // Path parameter name for task ID

	DefaultPage     = 1  // Default page number value
	DefaultPageSize = 10 // Default page size value
)

var (
	// ErrExceedsPageCount is returned when the requested page exceeds the total number of pages
	ErrExceedsPageCount = cerror.New(
		fiber.ErrBadRequest.Code,
		"current page exceeds total pages",
	)

	// Returned when the page is not a valid number
	ErrInvalidPage = cerror.New(
		fiber.ErrBadRequest.Code,
		"page must be a number",
	)

	// Returned when the page size is not a valid number
	ErrInvalidPageSize = cerror.New(
		fiber.ErrBadRequest.Code,
		"page_size must be a number",
	)

	// Returned when the date is not in ISO 8601 format (RFC3339)
	ErrInvalidDate = cerror.New(
		fiber.ErrBadRequest.Code,
		"date must be in ISO 8601 format (RFC3339)",
	)

	// Returned when the ID is not a positive number greater than 0
	ErrInvalidId = cerror.New(
		fiber.ErrBadRequest.Code,
		"id must be a positive number greater than 0",
	)

	// Returned when the request body is invalid or missing required fields
	ErrInvalidBody = cerror.New(
		fiber.ErrBadRequest.Code,
		"request body is invalid or missing required fields",
	)

	// Returned when the end time is incorrect
	ErrInvalidEndTime = cerror.New(
		fiber.ErrBadRequest.Code,
		"end_time must be after start_time and in ISO 8601 format (RFC3339)",
	)

	// Returned when the task with the specified ID is not found
	ErrTaskNotFound = cerror.New(
		fiber.ErrNotFound.Code,
		"task not found",
	)
)
