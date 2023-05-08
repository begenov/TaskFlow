package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/begenov/TaskFlow/user-app/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) CreateUser(ctx context.Context, user models.User) error {
	stmt := `INSERT INTO "user" (username, email, password) VALUES ($1, $2, $3)`
	if _, err := u.db.ExecContext(ctx, stmt, &user.Username, &user.Email, &user.Password); err != nil {
		fmt.Println("oraz")
		return fmt.Errorf("error %w", err)
	}
	return nil
}

func (u *UserStorage) UserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	stmt := `SELECT id, username, email, password, created_at FROM "user" WHERE email = $1`
	row := u.db.QueryRowContext(ctx, stmt, email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		log.Println("error posgresql", err)
		return user, err
	}
	return user, nil
}

func (u *UserStorage) UserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User

	stmt := `SELECT id, username, email, password, created_at FROM "user" WHERE id = $1`
	row := u.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStorage) SetSession(ctx context.Context, userID int, session models.Session) error {
	stmt, err := u.db.PrepareContext(ctx, "UPDATE \"user\" SET session=$1, last_visit_at=$2  WHERE id=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Printf("session.ExpiresAt: %v\n", session.ExpiresAt)
	_, err = stmt.ExecContext(ctx, session.RefreshToken, session.ExpiresAt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStorage) GetByRefreshToken(ctx context.Context, refreshToken string) (models.User, error) {
	var user models.User
	err := u.db.QueryRowContext(ctx, `SELECT id, username, email, session FROM "user" WHERE session = $1 AND last_visit_at :: timestamptz > $2`, refreshToken, time.Now()).Scan(&user.ID, &user.Username, &user.Email, &user.TokenStr)
	if err != nil {
		log.Println("model.user", err)
		return models.User{}, err
	}
	return user, nil
}
