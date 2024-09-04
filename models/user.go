package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Role     string `json:"role"`
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Location string `json:"location" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
