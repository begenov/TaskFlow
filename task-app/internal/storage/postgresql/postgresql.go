package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/begenov/TaskFlow/pkg/models"
)

type Task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *Task {
	return &Task{db: db}
}

func (t *Task) CreateTask(ctx context.Context, task models.Todo) error {
	stmt := `INSERT INTO "task" (title, description, user_id, created_at) VALUES ($1, $2, $3, $4)`
	log.Println(task.UserID)
	if _, err := t.db.ExecContext(ctx, stmt, task.Title, task.Description, task.UserID, &task.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (t *Task) AllTask(ctx context.Context) ([]models.Todo, error) {
	var tasks []models.Todo
	stmt := `SELECT * FROM "task"`
	row, err := t.db.QueryContext(ctx, stmt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	for row.Next() {
		var task models.Todo
		if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}

func (t *Task) TaskByID(ctx context.Context, id int) (models.Todo, error) {
	var task models.Todo
	stmt := `SELECT * FROM "task" WHERE id = $1`
	row := t.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt); err != nil {
		return task, err
	}
	return task, nil
}

func (t *Task) UpdateTask(ctx context.Context, task models.Todo) error {
	stmt := `UPDATE "task" SET title = $1, description = $2, created_at = $3 WHERE id = $4`
	_, err := t.db.ExecContext(ctx, stmt, task.Title, task.Description, task.CreatedAt, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) DeleteTask(ctx context.Context, id int) error {
	stmt := `DELETE FROM "task" WHERE id = $1`
	_, err := t.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
