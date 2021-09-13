package models

import (
	"github.com/google/uuid"
	"time"
)

type Film struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Genres []string `json:"genres"`
	Year int `json:"year"`
	Director []string `json:"director"`
	Authors []uuid.UUID `json:"authors"`
	Release time.Time `json:"release"`
	Duration int  `json:"duration"`
	Language string `json:"language"`
}
