package main

import (
	"time"
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

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Role      Role      `json:"role"`
}

func NewUser(name, email string) *User {

	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
}
