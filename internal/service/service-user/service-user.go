package serviceuser

import (
	"context"
	"fmt"
	"strings"

	"github.com/begenov/TaskFlow/models"
	"golang.org/x/crypto/bcrypt"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByEmail(ctx context.Context, email string) (models.User, error)
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

func (u *UserService) User(ctx context.Context, email string, password string) (models.User, error) {
	if ok := emptyUserPass(email, password); !ok {
		return models.User{}, fmt.Errorf("empty user pass")
	}

	user, err := u.user.UserByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	if ok := checkUserPass(user.Password, password); !ok {
		return user, fmt.Errorf("check user password incorrect")
	}

	return user, nil

}

func checkUserPass(password_hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}

func emptyUserPass(email string, password string) bool {
	return strings.Trim(email, " ") == "" || strings.Trim(password, " ") == ""

}

func generatePasswordAndHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
