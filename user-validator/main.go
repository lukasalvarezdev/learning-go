package userValidator

import fieldsValidator "learning-go/user-fields-validator"

func ValidateUser(age int, email string) (error, error) {
	ageError := fieldsValidator.ValidateUserAge(age)
	emailError := fieldsValidator.ValidateUserEmail(email)

	return ageError, emailError
}
