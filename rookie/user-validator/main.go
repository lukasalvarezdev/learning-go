package userValidator

import fieldsValidator "rookie/user-fields-validator"

func ValidateUser(age int, email string) (error, error) {
	ageError := fieldsValidator.ValidateUserAge(age)
	emailError := fieldsValidator.ValidateUserEmail(email)

	return ageError, emailError
}
