package models

import (
	"time"

	"github.com/golang-jwt/jwt"
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
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
