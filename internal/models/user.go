package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
}
