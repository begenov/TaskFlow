package mysql

import "database/sql"

type MySql struct {
	TodoMySql
	UserMySql
}

func NewMySql(db *sql.DB) *MySql {
	return &MySql{
		TodoMySql: *NewTodoMySql(db),
		UserMySql: *NewUserMySql(db),
	}
}

func (m *MySql) User() error {
	return nil
}
