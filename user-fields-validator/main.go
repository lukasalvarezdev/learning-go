package userFieldsValidator

import "errors"

func ValidateUserAge(age int) error {
	if age < 18 {
		return errors.New("user is too young to use our service")
	}

	return nil
}

func ValidateUserEmail(email string) error {
	if email == "" {
		return errors.New("user email is empty")
	}

	return nil
}
