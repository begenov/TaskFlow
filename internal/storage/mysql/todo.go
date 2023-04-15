package mysql

import (
	"database/sql"

	"github.com/begenov/TaskFlow/models"
)

type TodoMySql struct {
	db *sql.DB
}

func NewTodoMySql(db *sql.DB) *TodoMySql {
	return &TodoMySql{db: db}
}

func (t *TodoMySql) CreateTodo(todo models.Todo) error {
	return nil
}

func AllTodo() ([]models.Todo, error) {
	return nil, nil
}

func TodoById(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (t *TodoMySql) DeleteTodo(id int) error {
	return nil
}

func (t *TodoMySql) UpdateTodo(todo models.Todo) error {
	return nil
}
