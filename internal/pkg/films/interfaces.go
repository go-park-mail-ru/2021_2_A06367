package films

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=films

type FilmsUsecase interface {
	GetCompilation(topic string) ([]models.Film, models.StatusCode)
	GetSelection(selection string) ([]models.Film, models.StatusCode)
	GetByKeyword(keyword string) ([]models.Film, models.StatusCode)
	GetFilm(film models.Film) (models.Film, models.StatusCode)
	GetFilmsOfActor(actor models.Actors) ([]models.Film, models.StatusCode)
	GetCompilationForUser(user models.User) ([]models.Film, models.StatusCode)
	GetStartSelections(authorized bool, user models.User) ([]models.Film, models.StatusCode)

	GetStarred(user models.User) ([]models.Film, models.StatusCode)
	GetIfStarred(film models.Film, user models.User) models.StatusCode
	GetIfWatchlist(film models.Film, user models.User) models.StatusCode
	AddStarred(film models.Film, user models.User) models.StatusCode
	RemoveStarred(film models.Film, user models.User) models.StatusCode

	GetWatchlist(user models.User) ([]models.Film, models.StatusCode)
	AddWatchlist(film models.Film, user models.User) models.StatusCode
	RemoveWatchlist(film models.Film, user models.User) models.StatusCode

	Randomize() (models.Film, models.StatusCode)
	GetRating(film models.Film)  (models.Film, models.StatusCode)
	SetRating(film models.Film, user models.User, rating float64) models.StatusCode
	GetIdBySlug(slug string) (models.Film, models.StatusCode)
}

type FilmsRepository interface {
	GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode)
	GetHottestFilms() ([]models.Film, models.StatusCode)
	GetNewestFilms() ([]models.Film, models.StatusCode)
	GetFilmsByKeyword(keyword string) ([]models.Film, models.StatusCode)
	GetFilmById(film models.Film) (models.Film, models.StatusCode)
	GetFilmsByActor(actor models.Actors) ([]models.Film, models.StatusCode)
	GetFilmsByUser(user models.User) ([]models.Film, models.StatusCode)

	GetStarredFilms(user models.User) ([]models.Film, models.StatusCode)
	InsertStarred(film models.Film, user models.User) models.StatusCode
	DeleteStarred(film models.Film, user models.User) models.StatusCode
	IfStarred(film models.Film, user models.User) models.StatusCode
	IfWatchList(film models.Film, user models.User) models.StatusCode

	GetWatchlistFilms(user models.User) ([]models.Film, models.StatusCode)
	InsertWatchlist(film models.Film, user models.User) models.StatusCode
	DeleteWatchlist(film models.Film, user models.User) models.StatusCode

	GetRandom() (models.Film, models.StatusCode)
	SetRating(film models.Film, user models.User, rating float64) models.StatusCode
	GetRating(film models.Film) (models.Film, models.StatusCode)
	GetIdBySlug(slug string) (models.Film, models.StatusCode)
}
