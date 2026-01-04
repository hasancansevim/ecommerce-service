package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
	Server   ServerConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
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

type RedisConfig struct {
	Host string `envconfig:"REDIS_HOST" default:"localhost"`
	Port string `envconfig:"REDIS_PORT" default:"6379"`
}

type RabbitMQConfig struct {
	Host     string `envconfig:"RABBITMQ_HOST" default:"localhost"`
	Port     string `envconfig:"RABBITMQ_PORT" default:"5672"`
	User     string `envconfig:"RABBITMQ_USER" default:"guest"`
	Password string `envconfig:"RABBITMQ_PASSWORD" default:"guest"`
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
