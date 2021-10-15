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

func NewFilmsHandler(uc films.FilmsUsecase) *FilmsHandler {
	return &FilmsHandler{uc: uc}
}

func (h FilmsHandler) FilmByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genres, found := vars["genre"]
	if !found {
		util.Response(w, models.NotFound, nil)
	}

	filmSet, status := h.uc.GetCompilation(genres)
	util.Response(w, status, filmSet)
}

func (h FilmsHandler) FilmBySelection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	selection, found := vars["selection"]
	if !found {
		util.Response(w, models.NotFound, nil)
	}

	filmSet, status := h.uc.GetSelection(selection)
	util.Response(w, status, filmSet)
}

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
	filmSet, status := h.uc.GetFilmsOfActor(idActor)
	util.Response(w, status, filmSet)
}

func (h FilmsHandler) FilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, found := vars["film_id"]
	if !found {
		util.Response(w, models.NotFound, nil)
	}

	idFilm, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	film, status := h.uc.GetFilm(idFilm)
	util.Response(w, status, film)
}

func (h FilmsHandler) FilmsByUser(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	id := access.Id
	film, status := h.uc.GetCompilationForUser(id)
	util.Response(w, status, film)
}
