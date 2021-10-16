package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
)

type ActorsUsecase struct {
	repo actors.ActorsRepository
}

func NewActorsUsecase(repo actors.ActorsRepository) *ActorsUsecase {
	return &ActorsUsecase{repo: repo}
}

func (u ActorsUsecase) GetById(topic string) (models.Actors, models.StatusCode) {
	return u.repo.GetActorById(topic)
}