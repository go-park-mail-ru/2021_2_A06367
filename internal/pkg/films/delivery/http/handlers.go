package http

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FilmsHandler struct {
	logger *zap.SugaredLogger
	client grpc.FilmsServiceClient
}

func NewFilmsHandler(logger *zap.SugaredLogger, cl grpc.FilmsServiceClient) *FilmsHandler {
	return &FilmsHandler{
		logger: logger,
		client: cl,
	}
}

// FilmByGenre godoc
// @Summary Get films of genre
// @Description Get films of genre
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} []models.Film
// @Param genre path string true "Боевик"
// @Failure 400,404 {string} 1
// @Router /films/genre/{genre} [get]
func (h FilmsHandler) FilmByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genres, found := vars["genre"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	films, err := h.client.FilmByGenre(context.Background(), &grpc.KeyWord{Word: genres})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}

// FilmBySelection godoc
// @Summary Get details of films
// @Description Get details of films
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} []models.Film
// @Param selection path string true "КАКАЯ-ТО СТРОКА"
// @Failure 400,404 {string} 1
// @Router /films/selection/{selection} [get]
// @Router /films/selection/{selection} [options]
func (h FilmsHandler) FilmBySelection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	selection, found := vars["selection"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	films, err := h.client.FilmBySelection(context.Background(), &grpc.KeyWord{Word: selection})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}

// FilmByActor godoc
// @Summary Get details of actor from selection
// @Description Get details of actor from selection
// @Tags Film
// @Accept json
// @Produce json
// @Param actor_id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {object} []models.Film
// @Failure 400,404 {string} 1
// @Router /films/selection/actor/{actor_id} [get]
func (h FilmsHandler) FilmByActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, found := vars["actor_id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	idActor, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	films, err := h.client.FilmsByActor(context.Background(), &grpc.UUID{Id: idActor.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}

// FilmById godoc
// @Summary Get details of film
// @Description Get details of film
// @Tags Film
// @Accept json
// @Produce json
// @Param film_id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {object} models.Film
// @Failure 400,404 {string} 1
// @Router /films/film/{film_id} [get]
func (h FilmsHandler) FilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, found := vars["film_id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film, err := h.client.FilmById(context.Background(), &grpc.UUID{Id: id.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmToModel(film)
	util.Response(w, models.Okey, filmSet)

}

// FilmsByUser godoc
// @Summary Get details of personal film
// @Description Get details of personal film
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} models.Film
// @Failure 400,404 {string} 1
// @Router /films/selection/user/personal [get]
func (h FilmsHandler) FilmsByUser(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	films, err := h.client.FilmsByUser(context.Background(), &grpc.UUID{Id: access.Id.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}

func (h FilmsHandler) FilmStartSelection(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)

	if access == nil {
		access = &models.AccessDetails{}
	}
	selection, err := h.client.FilmStartSelection(context.Background(), &grpc.UUID{Id: access.Id.String()})
	if err != nil {
		return
	}
	film := h.FilmsToModels(*selection)
	util.Response(w, models.Okey, film)
}

func (h FilmsHandler) AddStarred(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	_, err = h.client.AddStarred(context.Background(), &grpc.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.Okey, nil)
}

func (h FilmsHandler) RemoveStarred(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	_, err = h.client.RemoveStarred(context.Background(), &grpc.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.Okey, nil)
}

func (h FilmsHandler) AddWatchlist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	_, err = h.client.AddWatchList(context.Background(), &grpc.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.Okey, nil)
}

func (h FilmsHandler) RemoveWatchlist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	_, err = h.client.RemoveWatchList(context.Background(), &grpc.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.Okey, nil)
}

func (h FilmsHandler) GetStarred(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	films, err := h.client.Starred(context.Background(), &grpc.UUID{
		Id: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}
	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}
func (h FilmsHandler) GetWatchlist(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}
	films, err := h.client.WatchList(context.Background(), &grpc.UUID{
		Id: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}
	filmSet := h.FilmsToModels(*films)
	util.Response(w, models.Okey, filmSet)
}

func (h FilmsHandler) RandomFilm(w http.ResponseWriter, r *http.Request) {

	film, _ := h.client.Random(context.Background(), &grpc.Nothing{})
	util.Response(w, models.Okey, film)

}

func (h FilmsHandler) FilmToModel(film *grpc.Film) models.Film {
	id, _ := uuid.Parse(film.Id)
	releaseru, _ := time.Parse("", film.ReleaseRus)
	release, _ := time.Parse("", film.Release)

	var actors []uuid.UUID
	for i := 0; i < len(film.Actors); i++ {
		id, _ := uuid.Parse(film.Actors[i])
		actors = append(actors, id)
	}

	return models.Film{
		Id:          id,
		Title:       film.Title,
		Genres:      film.Genres,
		Country:     film.Country,
		ReleaseRus:  releaseru,
		Year:        int(film.Year),
		Director:    film.Director,
		Authors:     film.Authors,
		Actors:      actors,
		Release:     release,
		Duration:    int(film.Duration),
		Budget:      film.Budget,
		Age:         int(film.Age),
		Pic:         film.Pic,
		Src:         film.Src,
		Description: film.Description,
		IsSeries:    film.IsSeries,
		Seasons:     nil,
	}
}

func (h FilmsHandler) FilmsToModels(films grpc.Films) []models.Film {
	var result []models.Film
	for i := 0; i < len(films.Data); i++ {
		film := h.FilmToModel(films.Data[i])
		result = append(result, film)
	}
	return result
}
