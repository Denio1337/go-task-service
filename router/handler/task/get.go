package task

import (
	cerror "app/router/types/error"
	"app/router/types/response"
	"app/router/validator"
	"app/service/task"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /tasks
func GetTasks(c *fiber.Ctx) error {
	// TODO: Is it possible to validate query parameters with Fiber's built-in validation?

	// Get pagination parameters
	page := uint(c.QueryInt(PageName, DefaultPage))
	pageSize := uint(c.QueryInt(PageSizeName, DefaultPageSize))

	// Extract date from query parameters
	date, err := getDateParam(c, DateName, time.Time{})
	if err != nil {
		return ErrInvalidDate
	}

	// Validate query parameters
	dto := &GetTasksDTO{
		Page:     page,
		PageSize: pageSize,
		Date:     date,
	}
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := task.GetTasks(&task.GetTasksParams{
		Page:     dto.Page, // Convert to zero-based index
		PageSize: dto.PageSize,
		Date:     dto.Date,
	})

	// Handle error from service
	if errors.Is(err, task.ErrExceedsPageCount) {
		return ErrExceedsPageCount
	}
	if err != nil {
		return err
	}

	// Transform result to DTO
	tasks := transformServiceTasks(result.Tasks)

	// TODO: Transform result to correct DTO with meta information about pagination
	return c.JSON(response.SuccessResponse(&TasksResponse{
		Tasks: tasks,
		Meta: Meta{
			Page:       dto.Page,
			PageSize:   dto.PageSize,
			TotalPages: result.TotalPages,
		},
	}))
}
