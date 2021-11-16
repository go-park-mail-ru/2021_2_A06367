package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type FilmsHandler struct {
	uc     films.FilmsUsecase
	logger *zap.SugaredLogger
}

func NewFilmsHandler(uc films.FilmsUsecase, logger *zap.SugaredLogger) *FilmsHandler {
	return &FilmsHandler{
		uc:     uc,
		logger: logger,
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

	filmSet, status := h.uc.GetCompilation(genres)
	util.Response(w, status, filmSet)
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

	filmSet, status := h.uc.GetSelection(selection)
	util.Response(w, status, filmSet)
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
	actor := models.Actors{Id: idActor}
	filmSet, status := h.uc.GetFilmsOfActor(actor)
	util.Response(w, status, filmSet)
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

	filmReq := models.Film{Id: id}
	film, status := h.uc.GetFilm(filmReq)
	util.Response(w, status, film)
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
	user := models.User{Id: access.Id}
	film, status := h.uc.GetCompilationForUser(user)
	util.Response(w, status, film)
}

func (h FilmsHandler) FilmStartSelection(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		film, status := h.uc.GetStartSelections(false, models.User{})
		util.Response(w, status, film)
		return
	}
	user := models.User{Id: access.Id}
	film, status := h.uc.GetStartSelections(true, user)
	util.Response(w, status, film)
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

	status := h.uc.AddStarred(film, user)
	util.Response(w, status, film)
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

	status := h.uc.RemoveStarred(film, user)
	util.Response(w, status, film)
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

	status := h.uc.AddWatchlist(film, user)
	util.Response(w, status, film)
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

	status := h.uc.RemoveWatchlist(film, user)
	util.Response(w, status, film)
}

func (h FilmsHandler) GetStarred(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	films, status := h.uc.GetStarred(user)
	util.Response(w, status, films)
}
func (h FilmsHandler) GetWatchlist(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}
	films, status := h.uc.GetWatchlist(user)
	util.Response(w, status, films)
}
