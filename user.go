package main

import (
	"fmt"

	"github.com/larsverp/go-picnic"
)

type User struct {
	Email    string
	Password string
	client   *picnic.Client
}

func (u *User) StartQuestions() {
	fmt.Println("What's your Picnic email? :")
	fmt.Scanln(&u.Email)

	fmt.Println("What's your picnic password? :")
	fmt.Scanln(&u.Password)
}

func (u *User) Login() {
	var err error
	u.client, err = picnic.NewClient(picnic.NewUser(u.Email, u.Password))
	if err != nil {
		fmt.Println("Error logging in: ", err)
	}
}
