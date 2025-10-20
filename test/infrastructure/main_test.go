package infrastructure

import (
	"context"
	"fmt"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/config"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/test/infrastructure/testdata"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	dbPool            *pgxpool.Pool
	ctx               context.Context
	productRepository persistence.IProductRepository

	testDataRegistry *testdata.TestDataRegistry
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	dbPool, _ := postgresql.GetConnectionPool(ctx, config.DatabaseConfig{})
	productRepository = persistence.NewProductRepository(dbPool)

	testDataRegistry = testdata.NewTestDataRegistry()
	testDataRegistry.Register(testdata.NewProductTestDataManager())

	fmt.Println("before all tests")
	exitCode := m.Run()
	fmt.Println("after all tests")
	dbPool.Close()
	os.Exit(exitCode)
}

func setupTestData(dataTypes ...string) {
	if err := testDataRegistry.InitializeSpecific(ctx, dbPool, dataTypes...); err != nil {
		panic(fmt.Sprintf("❌ Failed to setup test data: %v", err))
	}
}

func cleanUpTestData(dataTypes ...string) {
	if err := testDataRegistry.CleanupSpecific(ctx, dbPool, dataTypes...); err != nil {
		panic(fmt.Sprintf("❌ Failed to cleanup test data: %v", err))
	}
}

func setupProductTestData() {
	cleanUpTestData(testdata.ProductTestDataManager{}.GetName())
	setupTestData(testdata.ProductTestDataManager{}.GetName())
}

func cleanupProductTestData() {
	cleanUpTestData(testdata.ProductTestDataManager{}.GetName())
}
