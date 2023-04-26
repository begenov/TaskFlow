package storage

import (
	"context"
	"database/sql"

	usermysql "github.com/begenov/TaskFlow/internal/storage/user-mysql"
	"github.com/begenov/TaskFlow/models"
)

type User interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByEmail(ctx context.Context, email string) (models.User, error)
}

type Storage struct {
	User User
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		User: usermysql.NewUserStorage(db),
	}
}
