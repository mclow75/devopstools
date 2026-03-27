package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

// Глобальный инстанс
// Важно использовать WithRequiredStructEnabled
var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func main() {

	type User struct {
		Name     string `validate:"required"`
		Age      uint   `validate:"gte=0,lte=140"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8"`
	}

	user := User{
		Name:     "Иван Бровкин",
		Age:      19,
		Email:    "btovkin_i@home.local",
		Password: "Alsdl-vfEr!5+",
	}

	err := Validate.Struct(user)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			fmt.Println("Validation Errors:")
			for _, err := range validationErrors {
				fmt.Printf("Field: %s, Tag: %s, Value: %v\n", err.Field(), err.Tag(), err.Value())
			}
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Struct validated successfully!")
	}

	bad := User{
		Name:  "Анна Шульцман",
		Age:   150,             // Invalid age
		Email: "invalid-email", // Invalid email
		// none password
	}

	err = Validate.Struct(bad)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			fmt.Println("\nInvalid User Errors:")
			for _, err := range validationErrors {
				fmt.Printf("Field: %s, Tag: %s, Value: %v\n", err.Field(), err.Tag(), err.Value())
			}
		}
	}
}
