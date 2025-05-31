package task

import (
	"app/router/types/response"
	"app/service/task"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// DELETE /tasks/{id}
func DeleteTask(c *fiber.Ctx) error {
	// Extract ID from URL parameters
	id, err := c.ParamsInt(IdName)
	if err != nil {
		return ErrInvalidId
	}

	// Route to service
	err = task.DeleteTask(uint(id))

	// Handle specific error for task not found
	var errTaskNotFound *task.TaskNotFoundError
	if errors.As(err, &errTaskNotFound) {
		return ErrTaskNotFound
	}

	// Handle other errors
	if err != nil {
		return err
	}

	return c.JSON(response.SuccessResponse(err))
}
