package models

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Data struct {
	Todos    []Todo
	UserName string
	Todo     Todo
}
