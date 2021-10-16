package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	http2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type ActorHandler struct {
	uc     actors.ActorsUsecase
	logger *zap.SugaredLogger
}

func NewActorsHandler(uc actors.ActorsUsecase) *ActorHandler {
	return &ActorHandler{uc: uc}
}

func (h ActorHandler) ActorsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		http2.Response(w, models.NotFound, nil)
	}

	actor, status := h.uc.GetById(id)
	http2.Response(w, status, actor)
}