package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	if birthdate == "" {
		birthdate = "01/01/2001"
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "admin",
			lastName: "admin",
			birthdate: "01/01/2001",
			createdAt: time.Now(),
		},
	}
}

func (u User) OutPutUserData() {
	fmt.Printf("Created at: %s\n", u.createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserNames() {
	u.firstName = ""
	u.lastName = ""
}