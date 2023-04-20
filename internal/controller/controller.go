package controller

import (
	usercontroller "github.com/begenov/TaskFlow/internal/controller/user-controller"
	"github.com/begenov/TaskFlow/internal/service"
)

type userProvider interface{}

type Controller interface{}

type controller struct {
	user userProvider
}

func NewController(service service.Service) Controller {
	return &controller{
		user: usercontroller.NewUserController(service.User),
	}
}
