package serviceuser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/begenov/TaskFlow/pkg/auth"
	"github.com/begenov/TaskFlow/user-app/internal/config"
	"github.com/begenov/TaskFlow/user-app/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type userProvider interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByEmail(ctx context.Context, email string) (models.User, error)
	UserByID(ctx context.Context, id int) (models.User, error)
	SetSession(ctx context.Context, userID int, session models.Session) error
}

type UserService struct {
	user         userProvider
	cfg          config.Config
	tokenManager auth.TokenManager
}

func NewUserService(user userProvider, cfg *config.Config, tokenManager auth.TokenManager) *UserService {
	return &UserService{
		user:         user,
		cfg:          *cfg,
		tokenManager: tokenManager,
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

func (u *UserService) UserByID(ctx context.Context, id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, sql.ErrNoRows
	}

	return u.user.UserByID(ctx, id)

}

func (u *UserService) User(ctx context.Context, email string, password string) (models.Tokens, error) {
	if err := emptyUserPass(email, password); err != nil {
		return models.Tokens{}, err
	}
	user, err := u.user.UserByEmail(ctx, email)
	if err != nil {
		return models.Tokens{}, err
	}

	if ok := checkUserPass(user.Password, password); !ok {
		return models.Tokens{}, fmt.Errorf("incorrect: check email or password")
	}

	return u.createSession(ctx, user.ID)
}

func (u *UserService) ParseToken(accessToken string) (int, error) {
	return u.tokenManager.Parse(accessToken)
}

func (u *UserService) createSession(ctx context.Context, userID int) (models.Tokens, error) {
	var (
		res models.Tokens
		err error
	)
	res.AccessToken, err = u.tokenManager.NewJWT(userID, u.cfg.JWT.AccessTokenTTL)
	if err != nil {

		return res, err
	}
	res.RefreshToken, err = u.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := models.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(u.cfg.JWT.RefreshTokenTTL),
	}
	fmt.Println(session)
	if err := u.user.SetSession(ctx, userID, session); err != nil {
		return res, err
	}
	return res, nil
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

/*
func (u *UserService) genereteJWToken(userID int) (string, error) {
	claims := models.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(u.cfg.JWT.AccessTokenTTL).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokemstr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokemstr, nil

}*/
