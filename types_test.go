package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "john@pm.me", "verysecurepw")
	assert.Nil(t, err)

	fmt.Println(user)
}
