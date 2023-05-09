package storage

import (
	"context"
	"database/sql"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/begenov/TaskFlow/task-app/internal/storage/postgresql"
)

type Storage struct {
	Task Task
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Task: postgresql.NewTask(db),
	}
}

type Task interface {
	CreateTask(context.Context, models.Todo) error
	AllTask(context.Context) ([]models.Todo, error)
	TaskByID(context.Context, int) (models.Todo, error)
	UpdateTask(context.Context, models.Todo) error
	DeleteTask(context.Context, models.Todo) error
}
