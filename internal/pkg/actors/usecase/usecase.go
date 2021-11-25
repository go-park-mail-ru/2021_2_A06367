package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	"go.uber.org/zap"
)

type ActorsUsecase struct {
	repo   actors.ActorsRepository
	logger *zap.SugaredLogger
}

func NewActorsUsecase(repo actors.ActorsRepository, logger *zap.SugaredLogger) *ActorsUsecase {
	return &ActorsUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (u ActorsUsecase) GetById(actor models.Actors) (models.Actors, models.StatusCode) {
	return u.repo.GetActorById(actor)
}

func (u ActorsUsecase) GetByActors(actors []models.Actors) ([]models.Actors, models.StatusCode) {
	return u.repo.GetActors(actors)
}

func (u ActorsUsecase) GetByKeyword(keyword string) ([]models.Actors, models.StatusCode) {
	return u.repo.GetActorsByKeyword(keyword)
}
