package usecase

import "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"

type FilmsUsecase struct {
	repo films.FilmsRepository
}

func NewFilmsUsecase(repo films.FilmsRepository) *FilmsUsecase {
	return &FilmsUsecase{repo: repo}
}
