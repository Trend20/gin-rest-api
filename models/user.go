package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Role     string `json:"role"`
}
