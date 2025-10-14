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
	CreateUser(user domain.User) error
	GetUserByEmail(email string) (domain.User, error)
	GetUserById(id int64) (domain.User, error)
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
		return []domain.User{}
	}
	return users
}

func (userRepository *UserRepository) GetUserById(id int64) (domain.User, error) {
	ctx := context.Background()
	user, err := userRepository.scanner.QueryRowAndScan(ctx, "select * from users where id=$1", id)
	if err != nil {
		log.Error(err)
		return domain.User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	ctx := context.Background()
	user, err := userRepository.scanner.QueryRowAndScan(ctx, "select * from users where email=$1", email)
	if err != nil {
		log.Error(err)
		return domain.User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) CreateUser(user domain.User) error {
	ctx := context.Background()
	err := userRepository.scanner.ExecuteExec(ctx, "insert into users (first_name,last_name,email,password_hash,created_at) values ($1,$2,$3,$4,$5)",
		user.FirstName, user.LastName, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
