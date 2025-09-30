package app

import "go-ecommerce-service/common/postgresql"

type ConfigurationManager struct {
	PostgreSQLConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgresqlconfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSQLConfig: postgresqlconfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Username:              "postgres",
		Password:              "123456",
		DbName:                "ecommerce",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	}
}
