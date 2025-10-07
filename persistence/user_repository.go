package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IUserRepository interface {
	GetAllUser() []domain.User
	AddUser(user domain.User) error
}

type UserRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.User]
}

func NewUserRepository(dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanUser),
	}
}

func (userRepository *UserRepository) GetAllUser() []domain.User {
	ctx := context.Background()
	users, err := userRepository.scanner.QueryAndScan(ctx, "SELECT * FROM users")
	if err != nil {
		log.Error(err)
		return nil
	}
	return users
}

func (userRepository *UserRepository) AddUser(user domain.User) error {
	ctx := context.Background()
	err := userRepository.scanner.ExecuteExec(ctx, "insert into users (first_name,last_name,email,password,created_at) values ($1,$2,$3,$4,$5)",
		user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
