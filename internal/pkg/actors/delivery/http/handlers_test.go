package http

import (
	"encoding/json"
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
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
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

func TestActorNoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	usecase := mocks.NewMockActorsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.ActorsById(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestActorBadId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockActorsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"id": "abcd",
	})
	w := httptest.NewRecorder()

	handler.ActorsById(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
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
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
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
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	usecase := mocks.NewMockActorsUsecase(ctl)
	usecase.EXPECT().GetByActors([]models.Actors{{Id: uid}}).Times(1).Return([]models.Actors{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	js, err := json.Marshal([]models.Actors{{Id: uid}})
	if err != nil {
		t.Skip()
	}
	r := httptest.NewRequest("GET", "/film", strings.NewReader(string(js)))
	w := httptest.NewRecorder()

	handler.FetchActors(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestActorHandler_BadFetchActors(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	usecase := mocks.NewMockActorsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewActorsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/film", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FetchActors(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}
