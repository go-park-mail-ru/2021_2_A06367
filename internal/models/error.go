package models

type CustomError string

const (
	ErrNoToken      = "no token"
	ErrInvalidToken = "invalid token"
	ErrExpiredToken = "token expired"
	ErrNoAuthData   = "no auth data"
	ErrNoSecretKey  = "no secret key"
)
