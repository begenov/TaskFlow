package storage

import (
	"database/sql"

	"github.com/begenov/TaskFlow/task-app/internal/storage/postgresql"
)

type Storage struct {
	Task postgresql.Task
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Task: postgresql.NewTask(db),
	}
}
