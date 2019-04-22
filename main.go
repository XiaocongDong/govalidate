package main

import (
	"fmt"
	"github.com/XiaocongDong/govalidate/validate"
)

type User struct {
	FirstName string `validate:"required"`
	LastName string `validate:"required"`
}

func main () {
	user := User{FirstName: "Taolang"}
	validator := validate.New()

	errors := validator.Struct(user)

	if errors != nil {
		for _, error := range errors {
			fmt.Println(error.Tag())
		}
	}
}