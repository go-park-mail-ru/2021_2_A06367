package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all .\internal\models\film.go

type Film struct {
	Id       uuid.UUID   `json:"id"       example:"768eb570-2e0e-11ec-8d3d-0242ac130004"`
	Title    string      `json:"title"    example:"Достать ножи"`
	Genres   []string    `json:"genres"   example:"{'Боевик', 'Драма'}"`
	Year     int         `json:"year"     example:"2010"`
	Director []string    `json:"director" example:"{'Райан Джонсон', 'Леопольд Хьюз'}"`
	Authors  []string    `json:"authors"  example:"{'Райан Джонсон', 'Рэм Бергман'}"`
	Actors   []uuid.UUID `json:"actors"   example:"{'768eb570-2e0e-11ec-8d3d-0242ac130004', '9ebe8b02-30e2-11ec-8d3d-0242ac130003', 'a62bdb60-30e2-11ec-8d3d-0242ac130003'}"`
	Release  time.Time   `json:"release"  example:"2019-10-28"`
	Duration int         `json:"duration" example:"130"`
	Language string      `json:"language" example:"RU"`
}
