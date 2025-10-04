package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
)

type FakeUserRepository struct {
	users []domain.User
}

func NewFakeUserRepository(initialUsers []domain.User) persistence.IUserRepository {
	return &FakeUserRepository{
		users: initialUsers,
	}
}

func (fakeUserRepository FakeUserRepository) GetAllUser() []domain.User {
	return fakeUserRepository.users
}

func (fakeUserRepository FakeUserRepository) AddUser(user domain.User) error {
	fakeUserRepository.users = append(fakeUserRepository.users, domain.User{
		Id:        int64(len(fakeUserRepository.users) + 1),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	})
	return nil
}
