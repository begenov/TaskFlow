package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Secret = []byte("secret")

const Userkey = "user"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"password" binding:"required"`
	TokenStr  string
}

type Data struct {
	Todos    []Todo
	UserName string
	Todo     Todo
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Session struct {
	RefreshToken string
	ExpiresAt    time.Time
}
