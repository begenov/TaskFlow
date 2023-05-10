package models

type Todo struct {
	ID          int `json:"id"`
	UserID      int
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string
}

type Data struct {
	Todos    []Todo
	UserName string
	Todo     Todo
}
