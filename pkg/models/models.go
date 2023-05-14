package models

import "time"

type Todo struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string
	CreatedAt   time.Time
}

type Data struct {
	Todos    []Todo
	UserName string
	Todo     Todo
}
