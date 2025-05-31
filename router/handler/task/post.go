package task

import (
	cerror "app/router/types/error"
	"app/router/types/response"
	"app/router/validator"
	"app/service/task"
	"time"

	"github.com/gofiber/fiber/v2"
)

// POST /tasks
func AddTask(c *fiber.Ctx) error {
	var dto AddTaskDTO
	if err := c.BodyParser(&dto); err != nil {
		return ErrInvalidBody
	}

	// Validate input
	if errs := validator.Validate(&dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Default value for StartTime
	if dto.StartTime.IsZero() {
		dto.StartTime = time.Now()
	}

	// Check if EndTime is before StartTime
	if !dto.EndTime.IsZero() && dto.EndTime.Before(dto.StartTime) {
		return ErrInvalidEndTime
	}

	// Route to service
	created, err := task.CreateTask(&task.Task{
		Name:        dto.Name,
		Description: dto.Description,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
	})
	if err != nil {
		cerr := *cerror.ErrInternalServer
		cerr.Message = err.Error()
		return &cerr
	}

	return c.JSON(response.SuccessResponse(transformServiceTask(created)))
}
