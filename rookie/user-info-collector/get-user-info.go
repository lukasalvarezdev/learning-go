package main

import (
	"fmt"
	userValidator "rookie/user-validator"
)

func GetUserInfo() {
	var userName string
	var userAge int
	var userEmail string

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&userAge)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	ageError, emailError := userValidator.ValidateUser(userAge, userEmail)

	if ageError != nil || emailError != nil {
		fmt.Println("Error: ", ageError, emailError)
	} else {
		fmt.Printf("User: %v is %v years old and his email is %v \n", userName, userAge, userEmail)
	}

}
