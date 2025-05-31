package task

import (
	cerror "app/router/types/error"
	"app/router/types/response"
	"app/router/validator"
	"app/service/task"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// PATCH /tasks/{id}
func UpdateTask(c *fiber.Ctx) error {
	// Extract ID from URL parameters
	id, err := c.ParamsInt(IdName)
	if err != nil {
		return ErrInvalidId
	}

	// Validate body
	var dto UpdateTaskDTO
	if err := c.BodyParser(&dto); err != nil {
		return ErrInvalidBody
	}

	// Validate input
	if errs := validator.Validate(&dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Call service to update task
	updated, err := task.UpdateTask(&task.Task{
		ID:          uint(id),
		Name:        dto.Name,
		Description: dto.Description,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
	})
	var errTaskNotFound *task.TaskNotFoundError
	if errors.As(err, &errTaskNotFound) {
		return ErrTaskNotFound
	}
	if err != nil {
		return err
	}

	return c.JSON(response.SuccessResponse(transformServiceTask(updated)))
}
