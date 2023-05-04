package controller

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (c *controller) Router() *gin.Engine {

	mux := gin.Default()

	user := mux.Group("/user")
	{
		user.POST("/sign-up", c.user.SignUp)
		user.POST("/sign-in", c.user.SignIn)
	}

	home := mux.Group("/")

	{
		home.GET("/home", authMiddleware(), homepage)
	}

	return mux
}

func homepage(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"info": "ok",
	})

}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"ERROR": "Status Unauthorized",
			})
			return
		}

		token, err := jwt.Parse(authHeader[7:], func(t *jwt.Token) (interface{}, error) {
			return []byte("Oraz"), nil
		})

		if err != nil {
			log.Println("error")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		ctx.Set("user_id", userID)

		ctx.Next()
	}

}

// test
func test(t testing.M) {
	r := 0
	helloalem := "gfnhfg"
}
