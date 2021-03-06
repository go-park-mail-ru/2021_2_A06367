package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"go.uber.org/zap"
)

type FilmsUsecase struct {
	repo   films.FilmsRepository
	logger *zap.SugaredLogger
}

func NewFilmsUsecase(repo films.FilmsRepository, logger *zap.SugaredLogger) *FilmsUsecase {
	return &FilmsUsecase{
		repo:   repo,
		logger: logger,
	}
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

func (u *FilmsUsecase) GetStartSelections(authorized bool, user models.User) ([]models.Film, models.StatusCode) {
	if authorized {
		return u.GetCompilationForUser(user)
	}
	return u.GetSelection("hottest")
}

func (u *FilmsUsecase) AddStarred(film models.Film, user models.User) models.StatusCode {
	return u.repo.InsertStarred(film, user)
}

func (u *FilmsUsecase) GetIfStarred(film models.Film, user models.User) models.StatusCode {
	return u.repo.IfStarred(film, user)
}

func (u *FilmsUsecase) GetIfWatchlist(film models.Film, user models.User) models.StatusCode {
	return u.repo.IfWatchList(film, user)
}

func (u *FilmsUsecase) RemoveStarred(film models.Film, user models.User) models.StatusCode {
	return u.repo.DeleteStarred(film, user)
}

func (u *FilmsUsecase) AddWatchlist(film models.Film, user models.User) models.StatusCode {
	return u.repo.InsertWatchlist(film, user)
}

func (u *FilmsUsecase) RemoveWatchlist(film models.Film, user models.User) models.StatusCode {
	return u.repo.DeleteWatchlist(film, user)
}

func (u FilmsUsecase) GetStarred(user models.User) ([]models.Film, models.StatusCode) {
	return u.repo.GetStarredFilms(user)
}

func (u FilmsUsecase) GetWatchlist(user models.User) ([]models.Film, models.StatusCode) {
	return u.repo.GetWatchlistFilms(user)
}

func (u FilmsUsecase) Randomize() (models.Film, models.StatusCode) {
	return u.repo.GetRandom()
}

func (u FilmsUsecase) GetRating(film models.Film) (models.Film, models.StatusCode) {
	return u.repo.GetRating(film)
}

func (u FilmsUsecase) SetRating(film models.Film, user models.User, rating float64) models.StatusCode {
	return u.repo.SetRating(film, user, rating)
}

func (u FilmsUsecase) GetIdBySlug(slug string) (models.Film, models.StatusCode) {
	return u.repo.GetIdBySlug(slug)
}

func (u FilmsUsecase) GetRatingByUser(film models.Film, user models.User) (models.Film, models.StatusCode) {
	return u.repo.GetRatingByUser(film, user)
}
