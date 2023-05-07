package serviceuser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/begenov/TaskFlow/user-app/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	key = "Oraz"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByEmail(ctx context.Context, email string) (models.User, error)
	UserByID(ctx context.Context, id int) (models.User, error)
}

type UserService struct {
	user userProvider
}

func NewUserService(user userProvider) *UserService {
	return &UserService{
		user: user,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user models.User) error {
	hash, err := generatePasswordAndHash([]byte(user.Password))
	if err != nil {
		return fmt.Errorf("error in generate password and hash: %w", err)
	}
	user.Password = hash
	err = u.user.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error create user %w", err)
	}
	return nil
}

func (u *UserService) User(ctx context.Context, email string, password string) (models.User, error) {
	if err := emptyUserPass(email, password); err != nil {
		return models.User{}, err
	}
	user, err := u.user.UserByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	if ok := checkUserPass(user.Password, password); !ok {
		return user, fmt.Errorf("incorrect: check email or password")
	}

	user.TokenStr, err = genereteJWToken(user.ID)
	if err != nil {
		return user, fmt.Errorf("inctrrect: check email or password ")
	}

	return user, nil

}

func (u *UserService) UserByID(ctx context.Context, id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, sql.ErrNoRows
	}

	return u.user.UserByID(ctx, id)

}

func genereteJWToken(userID int) (string, error) {

	exe := time.Now().Add(10 * time.Minute)
	fmt.Println(userID)
	claims := models.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exe.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokemstr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokemstr, nil

}

func checkUserPass(password_hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}

func emptyUserPass(email string, password string) error {
	fmt.Println(email, password)
	if strings.TrimSpace(email) == "" {
		return fmt.Errorf("error empty email")
	}

	if strings.TrimSpace(password) == "" {
		return fmt.Errorf("error empty password")
	}
	return nil

}

func generatePasswordAndHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
