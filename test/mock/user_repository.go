package mock

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUser() []domain.User {
	args := m.Called()
	if args.Get(0) == nil {
		return []domain.User{}
	}
	return args.Get(0).([]domain.User)
}

func (m *MockUserRepository) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id int64) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

var _ persistence.IUserRepository = (*MockUserRepository)(nil)
