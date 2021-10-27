package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all .\internal\models\film.go

type Film struct {
	Id       uuid.UUID   `json:"id"`
	Title    string      `json:"title"`
	Genres   []string    `json:"genres"`
	Year     int         `json:"year"`
	Director []string    `json:"director"`
	Authors  []string    `json:"authors"`
	Actors   []uuid.UUID `json:"actors"`
	Release  time.Time   `json:"release"`
	Duration int         `json:"duration"`
	Language string      `json:"language"`
	Src      []string    `json:"src"`
}
