package usercontroller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	"github.com/dgrijalva/jwt-go"
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

	ctx.JSON(http.StatusCreated, "OK")
}

/*
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
*/

func (u *UserController) SignIn(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": fmt.Sprintf("error %v", err),
		})
		return
	}
	log.Println(user)
	if err := checkUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"content": fmt.Sprintf("check pls email and password: %v", err),
		})
		return
	}

	token, err := GenereteJWToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"content": fmt.Sprintf("error jwt token %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("successful: %v", token),
	})

}

func checkUser(user models.User) error {
	if user.Email != Test.Email || user.Password != Test.Password {
		return fmt.Errorf("error")
	}
	return nil
}

// func GenerateJWT(email, role string) (string, error) {
// 	var mySigningKey = []byte(secretkey)
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["email"] = email
// 	claims["role"] = role
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil

// }

func GenereteJWToken() (string, error) {

	// var mySigningKey = []byte(sampleSecretKey)
	exe := time.Now().Add(10 * time.Minute)

	claims := models.Claims{
		Username: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exe.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokemstr, err := token.SignedString([]byte("Oraz"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokemstr, nil

}

func ValidateJWToken(tokenString string) (string, error) {
	// Load the ECDSA private key for verification

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the public key for verification
		return "", nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Get the username from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	username, ok := claims["Username"].(string)
	if !ok {
		return "", errors.New("missing username in token claims")
	}

	return username, nil
}

var Test = models.User{
	Email:    "test@gmail.com",
	Password: "test",
}
