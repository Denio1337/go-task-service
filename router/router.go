package router

import (
	"app/router/handler/task"
	"app/router/types/response"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Create and configure Fiber application
func New() *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork:      true,              // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Task Service", // Set "Server" HTTP-header
		AppName:      "Go Task Service",
		ErrorHandler: handleError,
	})

	// Configure endpoints
	setupRoutes(router)

	return router
}

// Set router api
func setupRoutes(app *fiber.App) {
	// Middleware
	apiGroup := app.Group(ApiBasePath, logger.New())

	// Tasks
	taskGroup := apiGroup.Group(TasksPath)
	taskGroup.Get("/", task.GetTasks)
}

// Handle error response
func handleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(response.ErrorResponse(err.Error()))
}

const (
	ApiBasePath = "/api"

	TasksPath = "/tasks"
)
