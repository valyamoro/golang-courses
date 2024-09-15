package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"user_id"`
}
