package main

import "math/rand"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func NewUser(firstName, lastName, email string) *User {
	return &User{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
