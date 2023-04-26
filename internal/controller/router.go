package controller

import (
	"github.com/begenov/TaskFlow/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func (c *controller) Router() *gin.Engine {

	mux := gin.New()
	session := sessions.Default(&gin.Context{})
	mux.Use(gin.Recovery(), gin.Logger())
	store, _ := redis.NewStore(10, "tcp", "192.168.1.246:6379", "", []byte("secret"))
	session.Set(models.Userkey, store)
	mux.Use(sessions.Sessions("mysession", store))
	user := mux.Group("/user")
	{
		user.POST("/sign-up", c.user.SignUp)
		user.POST("/sign-in", c.user.SignIn)
	}

	return mux
}
