package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateTestUser() domain.User {
	return domain.User{
		Id:           1,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$hashed_password123",
	}
}

func CreateTestUserCreate() model.RegisterCreate {
	return model.RegisterCreate{
		FirstName: "New",
		LastName:  "User",
		Email:     "new@example.com",
		Password:  "password123",
	}
}

func CreateLoginRequest() model.LoginCreate {
	return model.LoginCreate{
		Email:    "test@example.com",
		Password: "password123",
	}
}

func CreateUserWithHashedPassword() domain.User {
	return domain.User{
		Id:           1,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "test@example.com",
		PasswordHash: "$2a$10$hashed_correct_password_123",
	}
}
