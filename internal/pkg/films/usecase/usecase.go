package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/google/uuid"
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

func (u *FilmsUsecase) GetFilm(id uint) (models.Film, models.StatusCode) {
	return u.repo.GetFilmById(id)
}

func (u *FilmsUsecase) GetFilmsOfActor(actor_id uuid.UUID) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByActor(actor_id)
}

func (u *FilmsUsecase) GetCompilationForUser(user_id uuid.UUID) ([]models.Film, models.StatusCode) {
	return u.repo.GetFilmsByUser(user_id)
}
