package config

import "os"

type Config struct {
	NatsURL            string
	Subject            string
	DBConnectionString string
	RedisURL           string
}

func NewConfig() *Config {
	return &Config{
		NatsURL:            getEnv("NATS_URL", "nats://localhost:4222"),
		Subject:            getEnv("NATS_SUBJECT", "orders"),
		DBConnectionString: getEnv("DB_CONNECTION_STRING", "postgres://postgres:root@localhost:5433/natsdb?sslmode=disable"),
		RedisURL:           getEnv("REDIS_URL", "redis://localhost:6379"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
