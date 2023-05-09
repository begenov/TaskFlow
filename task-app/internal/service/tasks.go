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

func (t *tasksService) TaskByID(ctx context.Context, id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (t *tasksService) AllTasks(ctx context.Context) ([]models.Todo, error) {
	return nil, nil
}

func (t *tasksService) UpdateTask(ctx context.Context, task models.Todo) error {
	return nil
}

func (t *tasksService) DeleteTask(ctx context.Context, id int) error {
	return nil
}
