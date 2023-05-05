package controller

import (
	usercontroller "github.com/begenov/TaskFlow/user-app/internal/controller/user-controller"
	"github.com/begenov/TaskFlow/user-app/internal/service"
	"github.com/gin-gonic/gin"
)

type userProvider interface {
	SignUp(ctx *gin.Context)
	SignUpForm(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type Controller interface {
	Router() *gin.Engine
}

type controller struct {
	user    userProvider
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		user:    usercontroller.NewUserController(service.User),
		service: service,
	}
}
