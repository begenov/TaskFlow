package storage

import (
	"database/sql"

	"github.com/begenov/TaskFlow/internal/pkg/lib/e"
	"github.com/begenov/TaskFlow/internal/storage/mysql"
)

type Storage struct {
	*mysql.MySql
}

func NewStorage() (*Storage, error) {
	db, err := sql.Open("", "")
	if err != nil {
		return nil, e.Wrap("", err)
	}
	if err := db.Ping(); err != nil {
		return nil, e.Wrap("", err)
	}

	return &Storage{mysql.NewMySql(db)}, nil
}
