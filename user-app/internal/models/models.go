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
	Password  string    `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
