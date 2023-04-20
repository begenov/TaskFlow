package usermysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/begenov/TaskFlow/models"
	"github.com/go-sql-driver/mysql"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) CreateUser(ctx context.Context, user models.User) error {
	stmt := `INSERT INTO user (username, email, password) VALUES (?, ?, ?)`
	if _, err := u.db.ExecContext(ctx, stmt, &user.Username, &user.Email, &user.Password); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("email already exists %v", err)
		}
		return fmt.Errorf("error %w", err)
	}
	return nil
}

//INSERT INTO users(email, password) VALUES(?, ?)
