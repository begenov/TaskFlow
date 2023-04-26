package usercontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/begenov/TaskFlow/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	User(ctx context.Context, email string, password string) (models.User, error)
}

type UserController struct {
	user userProvider
}

func NewUserController(user userProvider) *UserController {
	return &UserController{
		user: user,
	}
}

func (c *UserController) SignUpForm(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"susses": "ok",
	})
}

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
func (u *UserController) SignIn(ctx *gin.Context) {
	var err error
	session := sessions.Default(ctx)
	log.Println("error user sign in")
	user := session.Get(models.Userkey)
	if user != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": "Please logout first",
		})
		return
	}
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	user, err = u.user.User(context.Background(), email, password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": fmt.Sprintf("check pls email and password %v", err),
		})
	}
	session.Set(models.Userkey, user)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"content": fmt.Sprintf("failed to save session %v", err),
		})
	}
	ctx.JSON(http.StatusOK, session.Get(models.Userkey))
}
