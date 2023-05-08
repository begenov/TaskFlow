package storage

import (
	"context"
	"database/sql"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	userpostgresql "github.com/begenov/TaskFlow/user-app/internal/storage/user-postgresql"
)

type User interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByEmail(ctx context.Context, email string) (models.User, error)
	UserByID(ctx context.Context, id int) (models.User, error)
	SetSession(ctx context.Context, userID int, session models.Session) error
}

type Storage struct {
	User User
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		User: userpostgresql.NewUserStorage(db),
	}
}
