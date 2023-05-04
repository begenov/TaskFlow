package controller

import (
	"log"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func (c *controller) Router() *gin.Engine {

	mux := gin.New()

	store, _ := redis.NewStore(10, "tcp", "192.168.1.246:6379", "", []byte("secret"))
	log.Println("error")
	mux.Use(sessions.Sessions("mysession", store))
	log.Println("ok")
	session := sessions.Default(&gin.Context{})
	mux.Use(gin.Recovery(), gin.Logger())
	session.Set(models.Userkey, store)
	mux.Use(sessions.Sessions("mysession", store))
	user := mux.Group("/user")
	{
		user.POST("/sign-up", c.user.SignUp)
		user.POST("/sign-in", c.user.SignIn)
	}

	return mux
}
