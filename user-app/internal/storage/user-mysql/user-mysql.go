package usermysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/begenov/TaskFlow/user-app/internal/models"
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

func (u *UserStorage) UserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	stmt := `SELECT id, username, email, password, created_at FROM user WHERE email = ?`
	row := u.db.QueryRowContext(ctx, stmt, email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStorage) UserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User

	stmt := `SELECT id, username, email, password, created_at FROM user WHERE id = ?`
	row := u.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStorage) SetSession(ctx context.Context, userID int, session models.Session) error {
	stmt, err := u.db.PrepareContext(ctx, "UPDATE user SET session=?, lastVisitAt=? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, session.RefreshToken, session.ExpiresAt, userID)
	if err != nil {
		return err
	}
	return nil
}
