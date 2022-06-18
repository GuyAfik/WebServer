// Package entites defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entities

// User -.
type User struct {
	UserName string `json:"username" example:"username123"`
	Password string `json:"password" example:"abdjf841235"`
	Email    string `json:"email" example:"test@gmail.com"`
}
