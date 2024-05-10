package domain

import "gopkg.in/go-playground/validator.v9"

type User struct {
	Id       int    `json:"id" validate:"required,gte=0"`
	Username string `json:"username" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

func (u User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
