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
	log.Println("===========", header)
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

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ERROR": err.Error(),
		})

	}

	ctx.Set("user_id", userID)
}
