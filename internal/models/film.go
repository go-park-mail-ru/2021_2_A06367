package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all .\internal\models\film.go

type Film struct {
	Id           uuid.UUID   `json:"id"`
	Title        string      `json:"title"`
	Genres       []string    `json:"genres"`
	Country      string      `json:"country"`
	ReleaseRus   time.Time   `json:"release_rus"`
	Year         int         `json:"year"`
	Director     []string    `json:"director"`
	Authors      []string    `json:"authors"`
	Actors       []uuid.UUID `json:"actors"`
	Release      time.Time   `json:"release"`
	Duration     int         `json:"duration"`
	Language     string      `json:"language"`
	Budget       string      `json:"budget"`
	Age          int         `json:"age"`
	Pic          []string    `json:"pic"`
	Src          []string    `json:"src"`
	Description  string      `json:"description"`
	IsSeries     bool        `json:"is_series"`
	Seasons      *[]Season   `json:"seasons,omitempty"`
	Rating       float64     `json:"rating"`
	NeedsPayment bool        `json:"-"`         //платный или нет
	IsAvailable  bool        `json:"available"` //может ли пользователь смотреть
}

type Season struct {
	Num  int      //номер сезона
	Src  []string //список серий
	Pics []string //список превью у серий
}
