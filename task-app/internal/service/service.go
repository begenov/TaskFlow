package service

import (
	"context"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/begenov/TaskFlow/task-app/internal/storage"
)

type Tasks interface {
	CreateTask(ctx context.Context, task models.Todo) error
	TaskByID(ctx context.Context, id int) (models.Todo, error)
	AllTasks(ctx context.Context) ([]models.Todo, error)
	UpdateTask(ctx context.Context, task models.Todo, taskID int, userID int) error
	DeleteTask(ctx context.Context, id int) error
}

type Service struct {
	Task Tasks
}

func NewService(s storage.Storage) *Service {
	return &Service{
		Task: NewTaskService(s.Task),
	}
}
