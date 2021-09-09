package delivery

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
)

type FilmsHandler struct {
	uc films.FilmsUsecase
}

func NewFilmsHandler(uc films.FilmsUsecase) *FilmsHandler {
	return &FilmsHandler{uc: uc}
}
