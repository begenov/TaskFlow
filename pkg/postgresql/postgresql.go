package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("can't init db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connection db: %w", err)
	}
	return db, nil
}
