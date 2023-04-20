package controller

import "github.com/gin-gonic/gin"

func (c *controller) Router() *gin.Engine {

	mux := gin.New()

	mux.Use(gin.Recovery(), gin.Logger())

	user := mux.Group("/user")
	{
		user.GET("/sign-up", c.user.SignUpForm)
		user.POST("/sign-up", c.user.SignUp)
	}

	return mux
}
