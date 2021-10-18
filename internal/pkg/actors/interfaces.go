package actors

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=actors

type ActorsUsecase interface {
	GetById(id string) (models.Actors, models.StatusCode)
}

type ActorsRepository interface {
	GetActorById(id string) (models.Actors, models.StatusCode)
}