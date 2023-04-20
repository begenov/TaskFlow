package serviceuser

import (
	"context"
	"fmt"

	"github.com/begenov/TaskFlow/models"
	"golang.org/x/crypto/bcrypt"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
}

type UserService struct {
	user userProvider
}

func NewUserService(user userProvider) *UserService {
	return &UserService{
		user: user,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user models.User) error {
	hash, err := generatePasswordAndHash([]byte(user.Password))
	if err != nil {
		return fmt.Errorf("error in generate password and hash: %w", err)
	}
	user.Password = hash
	err = u.user.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error create user %w", err)
	}
	return nil
}

func generatePasswordAndHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
