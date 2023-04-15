package mysql

import (
	"database/sql"

	"github.com/begenov/TaskFlow/models"
)

type UserMySql struct {
	db *sql.DB
}

func NewUserMySql(db *sql.DB) *UserMySql {
	return &UserMySql{db: db}
}

func (u *UserMySql) CreateUser(user models.User) error {
	return nil
}

func UserByID(id int) (models.User, error) {
	return models.User{}, nil
}

func UserByEmail(email string) (models.User, error) {
	return models.User{}, nil
}

func UserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}
