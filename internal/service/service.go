package service

import (
	serviceuser "github.com/begenov/TaskFlow/internal/service/service-user"
	"github.com/begenov/TaskFlow/internal/storage"
	"github.com/begenov/TaskFlow/models"
)

type userProvider interface {
	CreateUser(user models.User) error
}

type Service struct {
	User userProvider
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		User: serviceuser.NewUserService(storage.User),
	}
}
