package validation

import "go-ecommerce-service/service/model"

func ValidateUserCreate(userCreate model.UserCreate) error {
	return NewValidator().
		RequiredString(userCreate.FirstName, "first_name").
		MinLength(userCreate.FirstName, "first_name", 2).
		RequiredString(userCreate.LastName, "last_name").
		MinLength(userCreate.LastName, "last_name", 2).
		RequiredString(userCreate.Email, "email").
		MinLength(userCreate.Email, "e_mail", 11).
		RequiredString(userCreate.PasswordHash, "password").
		MinLength(userCreate.PasswordHash, "password", 6).
		Error()
}
