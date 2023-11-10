package config

import "os"

type Config struct {
	NatsURL string
	Subject string
}

func NewConfig() *Config {
	return &Config{
		NatsURL: getEnv("NATS_URL", "nats://localhost:4222"),
		Subject: getEnv("NATS_SUBJECT", "orders"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
