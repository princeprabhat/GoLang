package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	IsActive  bool
}

type NewUser struct {
	Users []User
	Db    *gorm.DB
}

type SignUpSchema struct {
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
