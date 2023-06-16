package main

import (
	"fmt"
	"os/user"
)

func main() {
	user.Register(user.User{
		Name:     "Dwi",
		Email:    "dwih02@gmail.com",
		Password: "satu",
	})

	for _, data := range user.Get() {
		fmt.Printf("Name: %s, Email: %s, Password: %s\n", data.Nama, data.Email, data.Password)
	}
}
