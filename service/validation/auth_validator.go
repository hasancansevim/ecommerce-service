package validation

import "go-ecommerce-service/service/model"

func ValidateRegisterModel(registerModel model.RegisterCreate) error {
	return NewValidator().
		MinLength(registerModel.Password, "password", 6).
		Error()
}
