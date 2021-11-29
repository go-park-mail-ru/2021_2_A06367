package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestActorById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	usecase := mocks.NewMockActorsUsecase(ctl)
	usecase.EXPECT().GetById(models.Actors{Id: uid}).Times(1).Return(models.Actors{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.ActorsById(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestActorByIdNotFound(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	usecase := mocks.NewMockActorsUsecase(ctl)
	usecase.EXPECT().GetById(models.Actors{Id: uid}).Times(1).Return(models.Actors{}, models.NotFound)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error(err.Error())
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)

		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.ActorsById(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestActorHandler_FetchActors(t *testing.T) {

}
