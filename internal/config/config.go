package config

import (
	"github.com/joho/godotenv"
	"log"

	"os"
)

type Config struct {
	RPCUrl     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Continue   string // expect "true" or "false" (case-insensitive)
	FromBlock  string // if provided, indexing starts from this block number
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		RPCUrl:     os.Getenv("RPCUrl"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		Continue:   os.Getenv("CONTINUE"),   // new
		FromBlock:  os.Getenv("FROM_BLOCK"), // new
	}
}
