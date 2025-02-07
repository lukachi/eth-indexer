package config

import (
	"os"
)

type Config struct {
	RPCUrl     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func Load() Config {
	return Config{
		RPCUrl:     os.Getenv("RPCUrl"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
