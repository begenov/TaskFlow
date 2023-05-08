package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	NewJWT(userID int, ttl time.Duration) (string, error)
	Parse(accessToken string) (int, error)
	NewRefreshToken() (string, error)
}

type claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type manager struct {
	signinKey string
}

func NewManager(signinKey string) (TokenManager, error) {
	if signinKey == "" {
		return nil, errors.New("emoty signing key")
	}
	return &manager{
		signinKey: signinKey,
	}, nil
}

func (m *manager) NewJWT(userID int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	})
	return token.SignedString([]byte(m.signinKey))
}

func (m *manager) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(m.signinKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	userIDFloat64, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id is not a number")
	}
	userID := int(userIDFloat64)
	return userID, nil

}

func (m *manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
