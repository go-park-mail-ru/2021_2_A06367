package films

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

type FilmsUsecase interface {
	GetCompilation(topic string) ([]models.Film, models.StatusCode)
}

type FilmsRepository interface {
	GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode)
}
