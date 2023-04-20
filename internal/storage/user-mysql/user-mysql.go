package usermysql

import (
	"database/sql"

	"github.com/begenov/TaskFlow/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) CreateUser(user models.User) error {
	return nil
}
