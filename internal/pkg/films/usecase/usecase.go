package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
)

type FilmsUsecase struct {
	repo films.FilmsRepository
}

func NewFilmsUsecase(repo films.FilmsRepository) *FilmsUsecase {
	return &FilmsUsecase{repo: repo}
}

func (u FilmsUsecase) GetCompilation(topic string) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByTopic(topic)
}

func (u FilmsUsecase) GetSelection(selection string) ([]models.Film, models.StatusCode) {

	switch selection {
	case "hottest":
		return u.repo.GetHottestFilms()
	default:
		return u.repo.GetNewestFilms()
	}
}

func (u FilmsUsecase) GetByKeyword(keyword string) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByKeyword(keyword)
}

func (u *FilmsUsecase) GetFilm(film models.Film) (models.Film, models.StatusCode) {
	return u.repo.GetFilmById(film)
}

func (u *FilmsUsecase) GetFilmsOfActor(actor models.Actors) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByActor(actor)
}

func (u *FilmsUsecase) GetCompilationForUser(user models.User) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByUser(user)
}
