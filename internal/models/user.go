package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all ./internal/models/user.go

type User struct {
	Id                uuid.UUID `json:"-"`
	Login             string    `json:"login"    example:"Kalim1248"`
	Email             string    `json:"email"    example:"random@mail.ru"`
	EncryptedPassword string    `json:"password" example:"dd81d9f0a8c5e7904931b8c9ccbf429b"`
	CreatedAt         time.Time `json:"-"`
}

type LoginUser struct {
	Login             string `json:"login"    example:"Kalim1248"`
	EncryptedPassword string `json:"password" example:"dd81d9f0a8c5e7904931b8c9ccbf429b"`
}

type Profile struct {
	Id            uuid.UUID
	Login         string
	About         string
	Avatar        string
	Subscriptions uint
	Subscribers   uint
}
