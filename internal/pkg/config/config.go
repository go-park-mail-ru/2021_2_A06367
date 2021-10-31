package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load(".env")
}

func GetConnectionString() (string, error) {
	key, flag := os.LookupEnv("DATABASE_URL")
	if !flag {
		return "", errors.New("connection string not found")
	}
	return key, nil
}

func GetCsrfToken() ([]byte, error) {
	key, flag := os.LookupEnv("CSRF")
	if !flag {
		return nil, errors.New("CSRF string not found")
	}

	check := []byte(key)
	if len(check) != 32 {
		return nil, errors.New("invalid CSRF length")
	}
	return check, nil
}
