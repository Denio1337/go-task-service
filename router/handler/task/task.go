package task

import (
	cerror "app/router/types/error"
	"app/router/types/response"
	"app/router/validator"
	"app/service/task"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GET /tasks
func GetTasks(c *fiber.Ctx) error {
	// TODO: Is it possible to validate query parameters with Fiber's built-in validation?
	// TODO: Additional date query paremeter

	// Parse page from query parameters
	page, err := strconv.Atoi(c.Query(PageName, DefaultPage))
	if err != nil {
		return cerror.NewError(
			fiber.ErrBadRequest.Code,
			"page must be a number",
		)
	}

	// Parse page size from query parameters
	pageSize, err := strconv.Atoi(c.Query(PageSizeName, DefaultPageSize))
	if err != nil {
		return cerror.NewError(
			fiber.ErrBadRequest.Code,
			"page_size must be a number",
		)
	}

	// Validate query parameters
	dto := &GetDTO{
		Page:     page,
		PageSize: pageSize,
	}
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := task.GetTasks(&task.GetTasksParams{
		Page:     uint(dto.Page), // Convert to zero-based index
		PageSize: uint(dto.PageSize),
	})
	_ = result

	// Handle error from service
	// TODO: Handle specific errors with custom error messages
	if err != nil {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = err.Error()
		return &cerr
	}

	// Transform result to DTO
	tasks := make([]*Task, 0, len(result))
	for _, t := range result {
		tasks = append(tasks, transformTask(t))
	}

	// TODO: Transform result to correct DTO with meta information about pagination
	return c.JSON(response.SuccessResponse(&Response{
		Tasks: tasks,
		Meta: Meta{
			Page:       dto.Page,
			PageSize:   dto.PageSize,
			Total:      len(tasks),
			TotalPages: 0,
		},
	}))
}

// POST /tasks
// TODO: Implement this endpoint
func AddTask(c *fiber.Ctx) error {
	// TODO: Is it possible to validate query parameters with Fiber's built-in validation?
	// TODO: Additional date query paremeter

	// Parse page from query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return cerror.NewError(
			fiber.ErrBadRequest.Code,
			"page must be a number",
		)
	}

	// Parse page size from query parameters
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return cerror.NewError(
			fiber.ErrBadRequest.Code,
			"page_size must be a number",
		)
	}

	// Validate query parameters
	dto := &GetDTO{
		Page:     page,
		PageSize: pageSize,
	}
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := task.GetTasks(&task.GetTasksParams{
		Page:     uint(dto.Page), // Convert to zero-based index
		PageSize: uint(dto.PageSize),
	})
	_ = result

	// Handle error from service
	// TODO: Handle specific errors with custom error messages
	if err != nil {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = err.Error()
		return &cerr
	}

	// Transform result to DTO
	tasks := make([]*Task, 0, len(result))
	for _, t := range result {
		tasks = append(tasks, transformTask(t))
	}

	// TODO: Transform result to correct DTO with meta information about pagination
	return c.JSON(response.SuccessResponse(&Response{
		Tasks: tasks,
		Meta: Meta{
			Page:       dto.Page,
			PageSize:   dto.PageSize,
			Total:      len(tasks),
			TotalPages: 0,
		},
	}))
}

func transformTask(task *task.Task) *Task {
	return &Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}
