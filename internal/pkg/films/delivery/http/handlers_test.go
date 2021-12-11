package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	usecase2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/mocks"
	mocks2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/mocks"
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
	"time"
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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmByGenre(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)

	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		t.Error("wrong")
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar,usecase, use2)

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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmBySelection(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmsByActor(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	uid := uuid.New()

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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	//usecase.EXPECT().FilmsByActor(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmsByActor(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Status: 1,
	}, nil)

	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)


	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"actor_id": uuid.New().String(),
	})
	w := httptest.NewRecorder()

	handler.FilmByActor(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmById(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Film{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	uid := uuid.New()
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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	uid := uuid.New()

	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FilmById(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmById3(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)
	uid := uuid.New()

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

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	//usecase.EXPECT().FilmsByUser(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))

	w := httptest.NewRecorder()

	enc := usecase2.NewTokenator()

	str := enc.GetToken(models.User{Id: uuid.New(), Login: "hi"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value: str,
		Path:   "/",
		Domain: "a06367.ru",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}
	r.AddCookie(SSCookie)

	handler.FilmsByUser(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmStartSelection(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	//usecase.EXPECT().FilmStartSelection(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{}, nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			t.Error(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler.FilmStartSelection(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}
