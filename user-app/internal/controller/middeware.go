package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (c *controller) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if strings.TrimSpace(header) == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ERROR": "invalid auth header",
		})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ERROR": "invalid auth header",
		})
	}

	userID, err := c.service.User.ParseToken(headerParts[1])

	if err != nil {
		log.Println("error")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ERROR": err.Error(),
		})
	}

	ctx.Set("user_id", userID)

	// parse token
}

/*
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"ERROR": "Status Unauthorized 1",
			})
			return
		}

		token, err := jwt.Parse(authHeader[7:], func(t *jwt.Token) (interface{}, error) {
			return []byte("Oraz"), nil
		})

		fmt.Println(token)

		if err != nil {
			log.Println("error")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !token.Valid {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["user_id"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unouthorized",
			})
		}

		ctx.Set("user_id", userID)

		ctx.Next()
	}

*/
