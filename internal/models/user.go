package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all .\internal\models\user.go

type User struct {
	Id                uuid.UUID `json:"-"`
	Login             string    `json:"login"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"password"`
	CreatedAt         time.Time `json:"-"`
}

type LoginUser struct {
	Login string `json:"login"`
	EncryptedPassword string `json:"password"`
}