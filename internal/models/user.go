package models

import (
	"github.com/google/uuid"
	"time"
)

//TODO: сгенерировать easyjson файл

type User struct {
	Id                uuid.UUID `json:"-"`
	Login             string    `json:"login"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"password"`
	CreatedAt         time.Time `json:"-"`
}
