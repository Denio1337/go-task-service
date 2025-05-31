package task

import "app/storage/model"

// model.Task to Task
func transformModelTask(task *model.Task) *Task {
	return &Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}

// Task to model.Task
func transformTask(task *Task) *model.Task {
	return &model.Task{
		Name:        task.Name,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}
