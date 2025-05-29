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

func GetTasks(params *contract.GetTasksParams) ([]*model.Task, error) {
	return instance.GetTasks(params)
}

// Types

type GetTasksParams = contract.GetTasksParams
