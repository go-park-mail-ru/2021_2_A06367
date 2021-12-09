package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"go.uber.org/zap"
	"net/http"
)

type ActorHandler struct {
	uc     actors.ActorsUsecase
	logger *zap.SugaredLogger
}

func NewActorsHandler(uc actors.ActorsUsecase, logger *zap.SugaredLogger) *ActorHandler {
	return &ActorHandler{
		uc:     uc,
		logger: logger,
	}
}

// ActorsById godoc
// @Summary Get details of actor
// @Description Get details of actor
// @Tags Actors
// @Accept json
// @Produce json
// @Param id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {array} models.Actors
// @Failure 400,404 {string} 1
// @Router /actors/actor{id} [get]
func (h ActorHandler) ActorsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	idActor, err := uuid.Parse(id)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	actor := models.Actors{Id: idActor}
	actor, status := h.uc.GetById(actor)
	util.Response(w, status, actor)
}

func (h ActorHandler) FetchActors(w http.ResponseWriter, r *http.Request) {
	var actorsArr []models.Actors

	err := json.NewDecoder(r.Body).Decode(&actorsArr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	out, status := h.uc.GetByActors(actorsArr)
	util.Response(w, status, out)
}
