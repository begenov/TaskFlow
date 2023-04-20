package storage

import (
	"database/sql"

	usermysql "github.com/begenov/TaskFlow/internal/storage/user-mysql"
	"github.com/begenov/TaskFlow/models"
)

type User interface {
	CreateUser(user models.User) error
}

type Storage struct {
	User User
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		User: usermysql.NewUserStorage(db),
	}
}
