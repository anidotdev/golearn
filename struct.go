package main

import (
	"fmt"
)

type User struct {
	Username    string
	Email       string
	SignInCount int
	IsActive    bool
}

func main() {
	user := User{
		Username:    "john_doe",
		Email:       "john@example.com",
		SignInCount: 10,
		IsActive:    true,
	}
	fmt.Println(user)
}
