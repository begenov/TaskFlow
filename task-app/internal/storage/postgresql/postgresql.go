package postgresql

import (
	"database/sql"

	"github.com/begenov/TaskFlow/pkg/models"
)

type task struct {
	db *sql.DB
}

type Task interface {
	CreateTask(models.Todo) error
	AllTask() ([]models.Todo, error)
	TaskByID(id int) (models.Todo, error)
	UpdateTask(models.Todo) error
	RemoveTask(models.Todo) error
}

func NewTask(db *sql.DB) Task {
	return &task{}
}

func (t *task) CreateTask(task models.Todo) error {
	return nil
}

func (t *task) AllTask() ([]models.Todo, error) {
	return nil, nil
}

func (t *task) TaskByID(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (t *task) UpdateTask(models.Todo) error {
	return nil
}

func (t *task) RemoveTask(models.Todo) error {
	return nil
}
