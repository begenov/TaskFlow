package service

import (
	"context"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	serviceuser "github.com/begenov/TaskFlow/user-app/internal/service/service-user"
	"github.com/begenov/TaskFlow/user-app/internal/storage"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	User(ctx context.Context, email string, password string) (models.User, error)
}

type Service struct {
	User userProvider
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		User: serviceuser.NewUserService(storage.User),
	}
}