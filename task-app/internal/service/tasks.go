package service

import (
	"context"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/begenov/TaskFlow/task-app/internal/storage"
)

type tasksService struct {
	task storage.Task
}

func NewTaskService(task storage.Task) Tasks {
	return &tasksService{
		task: task,
	}
}

func (t *tasksService) CreateTask(ctx context.Context, task models.Todo) error {
	return nil
}
