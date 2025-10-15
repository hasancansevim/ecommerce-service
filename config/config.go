package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host                  string `envconfig:"DB_HOST" default:"localhost"`
	Port                  string `envconfig:"DB_PORT" default:"6432"`
	Username              string `envconfig:"DB_USERNAME" default:"postgres"`
	Password              string `envconfig:"DB_PASSWORD" default:"123456"`
	Name                  string `envconfig:"DB_NAME" default:"ecommerce"`
	MaxConnections        int    `envconfig:"DB_MAX_CONNECTIONS" default:"10"`
	MaxConnectionIdleTime string `envconfig:"DB_MAX_CONNECTION_IDLE_TIME" default:"10s"`
}

type AuthConfig struct {
	JWTSecret   string `envconfig:"JWT_SECRET" default:"akaimpkminik3"`
	JWTDuration string `envconfig:"JWT_DURATION" default:"24h"`
}

type ServerConfig struct {
	Port        string `envconfig:"SERVER_PORT" default:"8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
