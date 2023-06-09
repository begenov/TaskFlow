package service

import (
	"context"

	"github.com/begenov/TaskFlow/pkg/auth"
	"github.com/begenov/TaskFlow/user-app/internal/config"
	"github.com/begenov/TaskFlow/user-app/internal/models"
	serviceuser "github.com/begenov/TaskFlow/user-app/internal/service/service-user"
	"github.com/begenov/TaskFlow/user-app/internal/storage"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	User(ctx context.Context, email string, password string) (models.Tokens, error)
	UserByID(ctx context.Context, id int) (models.User, error)
	ParseToken(accessToken string) (int, error)
	RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error)
}

type Service struct {
	User userProvider
}

func NewService(storage storage.Storage, cfg *config.Config, tokenManager auth.TokenManager) *Service {
	return &Service{
		User: serviceuser.NewUserService(storage.User, cfg, tokenManager),
	}
}
