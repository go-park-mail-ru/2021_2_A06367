package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
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

//var testUsers = []models.User{
//	{Id: uuid.New(), Login: "User A"},
//	{Id: uuid.New(), Login: "User B"},
//}
//
//var testActors = []models.Actors{
//	{Id: uuid.New(), Name: "Burunov"},
//	{Id: uuid.New(), Name: "Petrov"},
//}
//
//var testFilms = []models.Film{
//	{Id: uuid.New(), Title: "Policeman from Rublevka", Genres: []string{"Comedy"}, Year: 2015, Director: []string{"Director"}, Authors: []string{"author"}},
//	{Id: uuid.New(), Title: "Mission Impossible", Genres: []string{"Triller"}},
//}

func TestFilmByGenre(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := films.NewMockFilmsUsecase(ctl)
	usecase.EXPECT().GetCompilation("topic").Times(1).Return([]models.Film{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase)

	r := httptest.NewRequest("GET", "/films/genres", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"genre": "topic",
	})
	w := httptest.NewRecorder()

	handler.FilmByGenre(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmBySelection(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := films.NewMockFilmsUsecase(ctl)
	usecase.EXPECT().GetSelection("hottest").Times(1).Return([]models.Film{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/films/hottest", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"selection": "hottest",
	})
	w := httptest.NewRecorder()

	handler.FilmBySelection(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmByActor(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	usecase := films.NewMockFilmsUsecase(ctl)
	usecase.EXPECT().GetFilmsOfActor(models.Actors{Id: uid}).Times(1).Return([]models.Film{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error(err.Error())
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"actor_id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.FilmByActor(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmByActor2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := films.NewMockFilmsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error(err.Error())
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"actor_id": "uid.String()",
	})
	w := httptest.NewRecorder()

	handler.FilmByActor(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmByActor3(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := films.NewMockFilmsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error(err.Error())
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FilmByActor(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	usecase := films.NewMockFilmsUsecase(ctl)
	usecase.EXPECT().GetFilm(models.Film{Id: uid}).Times(1).Return(models.Film{}, models.Okey)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"film_id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.FilmById(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmById2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	uid := uuid.New()
	usecase := films.NewMockFilmsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FilmById(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmById3(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	uid := uuid.New()
	usecase := films.NewMockFilmsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"film_id": "uid.String()",
	})
	w := httptest.NewRecorder()

	handler.FilmById(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmsByUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	usecase := films.NewMockFilmsUsecase(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))

	w := httptest.NewRecorder()

	handler.FilmsByUser(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmStartSelection(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	usecase := films.NewMockFilmsUsecase(ctl)
	usecase.EXPECT().GetStartSelections(false, models.User{}).Return([]models.Film{}, models.Okey)
	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(usecase, zapSugar)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FilmStartSelection(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}
