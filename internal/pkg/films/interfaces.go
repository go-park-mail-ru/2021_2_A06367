package films

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=films

type FilmsUsecase interface {
	GetCompilation(topic string) ([]models.Film, models.StatusCode)
	GetSelection(selection string) ([]models.Film, models.StatusCode)
	GetByKeyword(keyword string) ([]models.Film, models.StatusCode)
}

type FilmsRepository interface {
	GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode)
	GetHottestFilms() ([]models.Film, models.StatusCode)
	GetNewestFilms() ([]models.Film, models.StatusCode)
}
