package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *controller) Router() *gin.Engine {

	mux := gin.Default()

	user := mux.Group("/user")
	{
		user.POST("/sign-up", c.user.SignUp)
		user.POST("/sign-in", c.user.SignIn)
		user.POST("/auth/refresh", c.user.UserRefresh)

	}

	return mux
}
