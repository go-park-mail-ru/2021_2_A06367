package models

import (
	"github.com/google/uuid"
	"time"
)

// easyjson -all ./internal/models/actors.go

type Actors struct {
	Id          uuid.UUID `json:"id"            example:"768eb570-2e0e-11ec-8d3d-0242ac130004"`
	Name        string    `json:"name,omitempty"          example:"Дэниел"`
	Surname     string    `json:"surname,omitempty"       example:"Крейг"`
	Avatar      string    `json:"avatar,omitempty"        example:"./cmd/local/y.png"`
	Height      float32   `json:"height,omitempty"        example:"1.78"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty" example:"1968-03-03T02:18:00Z"`
	Description string    `json:"description,omitempty"   example:"Английский актёр, наиболее известный по роли Джеймса Бонда."`
	Genres      []string  `json:"genres,omitempty"        example:"{'Достать ножи', 'Казино рояль', 'Девушка с татуировкой дракона'}"`
}
