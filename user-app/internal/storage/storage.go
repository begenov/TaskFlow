package storage

import (
	"context"
	"database/sql"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	usermysql "github.com/begenov/TaskFlow/user-app/internal/storage/user-mysql"
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
