package infrastructure

import (
	"context"
	"fmt"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var userRepository persistence.IUserRepository
var user_ctx context.Context
var user_dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	user_ctx = context.Background()
	user_dbPool = postgresql.GetConnectionPool(user_ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Username:              "postgres",
		Password:              "123456",
		DbName:                "ecommerce",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})

	userRepository = persistence.NewUserRepository(user_dbPool)

	fmt.Println("Before All Tests - User")
	exitCode := m.Run()
	fmt.Println("After All Tests - User")
	os.Exit(exitCode)
}

func setup_userData(ctx context.Context, dbPool *pgxpool.Pool) {
	UserTestDataInitialize(ctx, dbPool)
}

func clear_userData(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateUserTestData(ctx, dbPool)
}

func TestGetAllUsers(t *testing.T) {
	setup_userData(user_ctx, user_dbPool)
	t.Run("Get All Users", func(t *testing.T) {
		actualUsers := userRepository.GetAllUser()
		fmt.Println(actualUsers)
		assert.Equal(t, 2, len(actualUsers))
	})
	clear_userData(user_ctx, user_dbPool)
}

func TestAddUser(t *testing.T) {
	setup_userData(user_ctx, user_dbPool)
	t.Run("Add User", func(t *testing.T) {
		userRepository.AddUser(domain.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "jhondoe",
			Password:  "123456",
			CreatedAt: time.Now(),
		})
		actualProducts := userRepository.GetAllUser()
		assert.Equal(t, 3, len(actualProducts))
	})
	clear_userData(user_ctx, user_dbPool)
}
