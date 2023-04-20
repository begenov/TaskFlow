package serviceuser

import (
	"github.com/begenov/TaskFlow/models"
)

type userProvider interface {
	CreateUser(user models.User) error
}

type UserService struct {
	user userProvider
}

func NewUserService(user userProvider) *UserService {
	return &UserService{
		user: user,
	}
}

func (u *UserService) CreateUser(user models.User) error {
	return u.user.CreateUser(user)
}
