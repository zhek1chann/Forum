package models

import (
	"forum/pkg/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Status         int
}

type UserSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (u UserSignupForm) FromToUser() User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	return User{
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: hashedPassword,
	}
}
