package usercontroller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/begenov/TaskFlow/models"
	"github.com/gin-gonic/gin"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
}

type UserController struct {
	user userProvider
}

func NewUserController(user userProvider) *UserController {
	return &UserController{
		user: user,
	}
}

// GET Sign-up
func (c *UserController) SignUpForm(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"susses": "ok",
	})
}

// POST Sign-up
func (c *UserController) SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		msg := fmt.Sprintf("%v", fmt.Errorf("invalid request body %w", err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	user.CreatedAt = time.Now()

	if user.Email == "" || user.Password == "" || user.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	if err := c.user.CreateUser(context.Background(), user); err != nil {
		msg := fmt.Sprintf("%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": msg,
		})
		return
	}

	ctx.JSON(http.StatusCreated, "Great job")
}
