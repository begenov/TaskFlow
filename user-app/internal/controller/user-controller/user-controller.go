package usercontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	"github.com/gin-gonic/gin"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	User(ctx context.Context, email string, password string) (models.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error)
	UserByID(ctx context.Context, id int) (models.User, error)
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

	ctx.JSON(http.StatusCreated, "OK")
}

func (u *UserController) SignIn(ctx *gin.Context) {

	var user models.User

	if err := ctx.BindJSON(&user); err != nil {

		panic(err)

	}

	log.Println("-------------")
	tokens, err := u.user.User(context.Background(), user.Email, user.Password)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": fmt.Sprint(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": tokens,
	})

}

type refreshToken struct {
	Token string `json:"token" binding:"required"`
}

func (c *UserController) UserRefresh(ctx *gin.Context) {
	var inp refreshToken
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": fmt.Sprint(err),
		})
		return
	}
	res, err := c.user.RefreshToken(context.Background(), inp.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"content": fmt.Sprint(err),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Tokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}
