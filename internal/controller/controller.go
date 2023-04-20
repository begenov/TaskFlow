package controller

import (
	usercontroller "github.com/begenov/TaskFlow/internal/controller/user-controller"
	"github.com/begenov/TaskFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type userProvider interface {
	SignUp(ctx *gin.Context)
	SignUpForm(ctx *gin.Context)
}

type Controller interface {
	Router() *gin.Engine
}

type controller struct {
	user userProvider
}

func NewController(service service.Service) Controller {
	return &controller{
		user: usercontroller.NewUserController(service.User),
	}
}
