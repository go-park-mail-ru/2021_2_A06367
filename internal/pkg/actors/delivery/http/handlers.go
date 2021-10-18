package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type ActorHandler struct {
	uc     actors.ActorsUsecase
	logger *zap.SugaredLogger
}

func NewActorsHandler(uc actors.ActorsUsecase, logger *zap.SugaredLogger) *ActorHandler {
	return &ActorHandler{
		uc: uc,
		logger: logger,
	}
}

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