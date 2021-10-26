package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all .\internal\models\film.go

type Film struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Genres   []string  `json:"genres"`
	Year     int       `json:"year"`
	Director []string  `json:"director"`
	Authors  []string  `json:"authors"`
	Actors   []uuid.UUID `json:"actors"   example:"{'768eb570-2e0e-11ec-8d3d-0242ac130004', '9ebe8b02-30e2-11ec-8d3d-0242ac130003', 'a62bdb60-30e2-11ec-8d3d-0242ac130003'}"`
	Release  time.Time `json:"release"`
	Duration int       `json:"duration"`
	Language string    `json:"language"`
	Src		 []string  `json:"src"`
}
