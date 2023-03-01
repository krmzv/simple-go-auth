package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role struct {
	Slug string `json:"slug"`
}

func (r Role) String() string {
	return r.Slug
}

var (
	Unknown   = Role{""}
	Company   = Role{"company"}
	Developer = Role{"developer"}
	Admin     = Role{"admin"}
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	Role      Role      `json:"role"`
}

func NewUser(name, email, password string) (*User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:      name,
		Email:     email,
		Password:  string(pw),
		CreatedAt: time.Now().UTC(),
	}, nil
}
