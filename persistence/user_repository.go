package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IUserRepository interface {
	GetAllUser() []domain.User
	AddUser(user domain.User) error
}

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{
		dbPool: dbPool,
	}
}

func (userRepository *UserRepository) GetAllUser() []domain.User {
	ctx := context.Background()
	getAllUserSql := "select * from users"
	userRows, queryRowErr := userRepository.dbPool.Query(ctx, getAllUserSql)
	if queryRowErr != nil {
		log.Error(queryRowErr.Error())
		return []domain.User{}
	}
	usersFromRows := extractUsersFromRows(userRows)
	return usersFromRows
}

func (userRepository *UserRepository) AddUser(user domain.User) error {
	ctx := context.Background()
	addUserSql := `insert into users (first_name,last_name,email,password,created_at) values ($1,$2,$3,$4,$5)`
	addNewUser, execErr := userRepository.dbPool.Exec(ctx, addUserSql,
		user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)

	if execErr != nil {
		log.Error(execErr.Error())
		return execErr
	}
	log.Info("Added new user %v : ", addNewUser)
	return nil
}

func extractUsersFromRows(rows pgx.Rows) []domain.User {
	var users = []domain.User{}
	var id int64
	var fist_name string
	var last_name string
	var email string
	var password string
	var created_at time.Time

	for rows.Next() {
		rows.Scan(&id, &fist_name, &last_name, &email, &password, &created_at)
		users = append(users, domain.User{
			Id:        id,
			FirstName: fist_name,
			LastName:  last_name,
			Email:     email,
			Password:  password,
			CreatedAt: created_at,
		})
	}
	return users
}
