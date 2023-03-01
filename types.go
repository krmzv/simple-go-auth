package main

import (
	"time"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(firstName, lastName, email string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
}
