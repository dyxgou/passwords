package models

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Passwords int    `json:"passwords"`
	Points    int    `json:"points"`
}
