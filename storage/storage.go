package storage

import (
	"app/storage/contract"
	"app/storage/impl"
	"app/storage/model"
)

// Global storage instance
var instance contract.Storage

// Create DB Connection with current implementation
func init() {
	instance = impl.Impl
}

// Interface

func GetTasks(params *GetTasksParams) (*GetTasksResult, error) {
	return instance.GetTasks(params)
}

func CreateTask(task *model.Task) (*model.Task, error) {
	return instance.CreateTask(task)
}

func UpdateTask(params *UpdateTaskParams) (*model.Task, error) {
	return instance.UpdateTask(params)
}

func DeleteTask(id uint) error {
	return instance.DeleteTask(id)
}

func GetDates(params *GetDatesParams) (*GetDatesResult, error) {
	return instance.GetDates(params)
}

// Types

type GetTasksParams = contract.GetTasksParams
type GetTasksResult = contract.GetTasksResult
type GetDatesParams = contract.GetDatesParams
type GetDatesResult = contract.GetDatesResult
type UpdateTaskParams = contract.UpdateTaskParams

// Errors

type TaskNotFoundError = contract.TaskNotFoundError
