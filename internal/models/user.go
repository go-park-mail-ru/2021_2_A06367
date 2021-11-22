package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all ./internal/models/user.go

type User struct {
	Id                uuid.UUID `json:"id"`
	Login             string    `json:"login"    example:"Kalim1248"`
	EncryptedPassword string    `json:"password" example:"dd81d9f0a8c5e7904931b8c9ccbf429b"`
	CreatedAt         time.Time `json:"-"`
}

type LoginUser struct {
	Login             string `json:"login"    example:"Kalim1248"`
	EncryptedPassword string `json:"password" example:"dd81d9f0a8c5e7904931b8c9ccbf429b"`
}

type Profile struct {
	Id            uuid.UUID `json:"id"`
	Login         string    `json:"login"`
	About         string    `json:"about"`
	Avatar        string    `json:"avatar"`
	Subscriptions uint      `json:"subscriptions"`
	Subscribers   uint      `json:"subscribers"`
}

type PassUpdate struct {
	Password string `json:"password"`
}

type BioUpdate struct {
	About string `json:"about"`
}
