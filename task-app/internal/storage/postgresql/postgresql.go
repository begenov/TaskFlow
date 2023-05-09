package postgresql

import (
	"context"
	"database/sql"

	"github.com/begenov/TaskFlow/pkg/models"
)

type Task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *Task {
	return &Task{db: db}
}

func (t *Task) CreateTask(ctx context.Context, task models.Todo) error {
	return nil
}

func (t *Task) AllTask(ctx context.Context) ([]models.Todo, error) {
	return nil, nil
}

func (t *Task) TaskByID(ctx context.Context, id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (t *Task) UpdateTask(ctx context.Context, task models.Todo) error {
	return nil
}

func (t *Task) DeleteTask(ctx context.Context, task models.Todo) error {
	return nil
}
