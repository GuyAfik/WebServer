// Package entites defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entities

// UserEntity -.
type UserEntity struct {
	Username string `json:"username" example:"username123" binding:"required"`
	Password string `json:"password" example:"abdjf841235" binding:"required"`
	Email    string `json:"email" example:"test@gmail.com" binding:"required"`
}
