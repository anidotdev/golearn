package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Username    string
	Email       string
	SignInCount int
	IsActive    bool
}

// value receiver
func (u User) PrintInfo() {
	fmt.Printf("Username: %s, Email: %s, SignInCount: %d, IsActive: %t\n", u.Username, u.Email, u.SignInCount, u.IsActive)
}

// pointer receiver
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// interface
type UserInfo interface {
	PrintInfo()
}

func ShowInfo(ui UserInfo) {
	ui.PrintInfo()
}

// memory alignment
type Example struct {
	b int32
	a int8
	c int8
}

// nested struct
type Address struct {
	City    string
	Country string
}

type NestedUser struct {
	Username string
	Email    string
	Address  Address // Nested struct
}

// struct composition

type Profile struct {
	Age int
	Bio string
}

type ComposedUser struct {
	Username string
	Email    string
	Address  Address // struct composition
	Profile  Profile // struct composition
}

func main() {

	user1 := User{
		Username:    "john_doe",
		Email:       "john@example.com",
		SignInCount: 10,
		IsActive:    true,
	}
	ShowInfo(user1)
	user1.UpdateEmail("new@example.com")
	ShowInfo(user1)

	user := NestedUser{
		Username: "john_doe",
		Email:    "john@example.com",
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}

	fmt.Println(unsafe.Sizeof(Example{}))
	fmt.Println("User: %s, Email: %s, City: %s, Country: %s\n", user.Username, user.Email, user.Address.City, user.Address.Country)

	user2 := ComposedUser{
		Username: "john_doe",
		Email:    "john@example.com",
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		Profile: Profile{
			Age: 30,
			Bio: "Software engineer",
		},
	}
	// Access fields of the composed struct
	fmt.Printf("User: %s, Email: %s, City: %s, Age: %d, Bio: %s\n", user2.Username, user2.Email, user2.Address.City, user2.Profile.Age, user2.Profile.Bio)
}
