package task

import (
	"app/service/task"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Get date parameter from query with a default value
func getDateParam(c *fiber.Ctx, name string, defaultValue time.Time) (time.Time, error) {
	// Extract date from query parameters
	dateStr := c.Query(name, defaultValue.Format(time.RFC3339))
	if dateStr == "" {
		return defaultValue, nil
	}

	// Parse the date string
	parsedDate, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return defaultValue, err
	}
	return parsedDate, nil
}

// Transform task.Task to Task
func transformServiceTask(task *task.Task) *Task {
	return &Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}

// Transform []task.Task to []Task
func transformServiceTasks(tasks []*task.Task) []*Task {
	result := make([]*Task, len(tasks))
	for i, task := range tasks {
		result[i] = transformServiceTask(task)
	}
	return result
}
