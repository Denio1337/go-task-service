package task

import (
	"app/storage"
	"math"
	"time"
)

func GetTasks(params *GetTasksParams) (*GetTasksResult, error) {
	// Calculate offset and limit for pagination
	offset := (params.Page - 1) * params.PageSize
	limit := params.PageSize

	// Get tasks from storage
	// TODO: Handle storage errors with custom errors
	result, err := storage.GetTasks(&storage.GetTasksParams{
		Offset: offset,
		Limit:  limit,
		Date:   params.Date.Truncate(24 * time.Hour),
	})
	if err != nil {
		return nil, err
	}

	// Transform model.Task to Task
	// TODO: Is there a better way to do this?
	transformedTasks := make([]*Task, 0, len(result.Tasks))
	for _, task := range result.Tasks {
		transformedTasks = append(transformedTasks, transformModelTask(task))
	}

	// Calculate total pages based on total tasks and page size
	var totalPages uint
	switch result.Total {
	case 0:
		totalPages = 1 // If no tasks, we still return one page
	default:
		totalPages = uint(math.Ceil(float64(result.Total) / float64(params.PageSize)))
	}

	// Check if requested page exceeds total pages
	if params.Page > totalPages {
		return nil, ErrExceedsPageCount
	}

	return &GetTasksResult{
		Tasks:      transformedTasks,
		TotalPages: totalPages,
	}, nil
}

func CreateTask(task *Task) (*Task, error) {
	// Transform Task to model.Task
	modelTask := transformTask(task)

	// Create task in storage
	createdTask, err := storage.CreateTask(modelTask)
	if err != nil {
		return nil, err
	}

	return transformModelTask(createdTask), nil
}

func UpdateTask(task *Task) (*Task, error) {
	// Transform Task to model.Task
	modelTask := transformTask(task)

	// Update task in storage
	updatedTask, err := storage.UpdateTask(&storage.UpdateTaskParams{
		Task: modelTask,
		ID:   task.ID,
	})
	if err != nil {
		return nil, err
	}

	return transformModelTask(updatedTask), nil
}

func DeleteTask(id uint) error {
	// Delete task in storage
	err := storage.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
