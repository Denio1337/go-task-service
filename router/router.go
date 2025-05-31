package router

import (
	"app/config"
	"app/router/handler/task"
	"app/router/types/consts"
	"app/router/types/response"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Create and configure Fiber application
func New() *fiber.App {
	router := fiber.New(fiber.Config{
		// Spawn multiple Go processes listening on the same port
		Prefork:      config.GetBool(config.EnvMultipleProcesses),
		ServerHeader: consts.ServerHeader, // Set "Server" HTTP-header
		AppName:      consts.ServerHeader,
		ErrorHandler: handleError,
	})

	// Configure endpoints
	setupRoutes(router)

	return router
}

// Set router api
func setupRoutes(app *fiber.App) {
	// Base API path
	apiGroup := app.Group(consts.ApiBasePath, logger.New())

	// Tasks
	taskGroup := apiGroup.Group(consts.TasksPath)
	taskGroup.Get("/", task.GetTasks)
	taskGroup.Post("/", task.AddTask)

	idPath := fmt.Sprintf("/:%s", task.IdName)
	taskGroup.Patch(idPath, task.UpdateTask)
	taskGroup.Delete(idPath, task.DeleteTask)

	// Tasks dates
	taskGroup.Get(consts.TaskDatesPath, task.GetDates)
}

// Handle error response
func handleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := consts.ErrorStatusCode

	// Retrieve custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(response.ErrorResponse(err.Error()))
}
