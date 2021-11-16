package actors

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=actors

type ActorsUsecase interface {
	GetById(actor models.Actors) (models.Actors, models.StatusCode)

	GetByActors(actor []models.Actors) ([]models.Actors, models.StatusCode)

	GetByKeyword(keyword string) ([]models.Actors, models.StatusCode)
}

type ActorsRepository interface {
	GetActorById(actor models.Actors) (models.Actors, models.StatusCode)

	GetActors(actor []models.Actors) ([]models.Actors, models.StatusCode)

	GetActorsByKeyword(keyword string) ([]models.Actors, models.StatusCode)
}
