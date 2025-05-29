package task

import (
	"app/storage"
	"app/storage/model"
)

func GetTasks(params *GetTasksParams) ([]*Task, error) {
	// Calculate offset and limit for pagination
	offset := (params.Page - 1) * params.PageSize
	limit := params.PageSize

	// Get tasks from storage
	// TODO: Handle storage errors with custom errors
	tasks, err := storage.GetTasks(&storage.GetTasksParams{
		Offset: offset,
		Limit:  limit,
		Date:   params.Date,
	})
	if err != nil {
		return nil, err
	}

	// Transform model.Task to Task
	// TODO: Is there a better way to do this?
	transformedTasks := make([]*Task, 0, len(tasks))
	for _, task := range tasks {
		transformedTasks = append(transformedTasks, transformTask(task))
	}

	return transformedTasks, nil
}

func transformTask(task *model.Task) *Task {
	return &Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}
