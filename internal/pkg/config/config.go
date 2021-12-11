package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("config not loaded")
	}
}

func GetConnectionString() (string, error) {
	key, flag := os.LookupEnv("DATABASE_URL")
	if !flag || key == "" {
		return "", errors.New("connection string not found")
	}
	return key, nil
}
