package main

import (
	"fmt"
	"github.com/XiaocongDong/govalidate/validate"
)

type User struct {
	FirstName string `validate:"required"`
}

func main () {
	user := User{"Taolang"}
	validator := validate.New()

	errors := validator.Struct(user)

	if errors != nil {
		fmt.Println(errors)
	}
}