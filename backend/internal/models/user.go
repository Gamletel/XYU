package models

import "github.com/go-playground/validator/v10"

type User struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Login    string `json:"login,omitempty" validate:"required,min=5,max=20"`
	Password string `json:"-" validate:"required,min=5"`
	Name     string `json:"name,omitempty" validate:"min=1"`
	Surname  string `json:"surname,omitempty" validate:"min=1"`
	Avatar   string `json:"avatar,omitempty" validate:"min=1"`
}

func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}
