package http

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	usecase2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/mocks"
	generated2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
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

func prepare() *generated.Film {
	return &generated.Film{
		Id:                 uuid.NewString(),
		Title:              "",
		Genres:             []string{},
		Country:            "",
		Year:               0,
		ReleaseRus:         "2006-01-02",
		Director:           []string{},
		Authors:            []string{},
		Actors:             []string{},
		Release:            "2006-01-02",
		Duration:           0,
		ReleaseRusLanguage: "",
		Budget:             "",
		Age:                0,
		Pic:                []string{},
		Src:                []string{},
		Description:        "",
		IsSeries:           false,
		Seasons: []*generated.Season{
			&generated.Season{
				Num:  0,
				Src:  []string{},
				Pics: []string{},
			},
		},
		Status:       0,
		Rating:       0,
		NeedsPayment: true,
		Slug:         "",
	}
}
func TestFilmByGenre(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmByGenre(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
	}, nil)

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
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	r := httptest.NewRequest("GET", "/films/genres", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"genre": "topic",
	})
	w := httptest.NewRecorder()

	handler.FilmByGenre(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmByGenre2(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)


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
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	r := httptest.NewRequest("GET", "/films/genres", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
	})
	w := httptest.NewRecorder()

	handler.FilmByGenre(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmByGenre3(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmByGenre(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
	}, errors.New(""))

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
	handler := NewFilmsHandler(zapSugar, usecase, use2)

	r := httptest.NewRequest("GET", "/films/genres", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"genre": "topic",
	})
	w := httptest.NewRecorder()

	handler.FilmByGenre(w, r)
	require.Equal(t, w.Code, http.StatusNotFound)
	log.Print(w.Body.String())
}

func TestFilmBySelection(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmBySelection(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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
	usecase.EXPECT().FilmsByActor(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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
		Data: []*generated.Film{prepare()}, Status: 1,
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
	usecase.EXPECT().FilmById(gomock.Any(), gomock.Any()).Times(1).Return(prepare(), nil)
	use2 := mocks2.NewMockSubsServiceClient(ctl)
	use2.EXPECT().GetLicense(gomock.Any(), gomock.Any()).Return(&generated2.License{
		Status: 0, ExpiresDate: time.Now().Add(time.Hour).String(),
	}, nil)
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
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
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
		Name:     "SSID",
		Value:    str,
		Path:     "/",
		Domain:   "a06367.ru",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}
	r.AddCookie(SSCookie)

	handler.FilmsByUser(w, r)
	require.Equal(t, w.Code, http.StatusBadRequest)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmByUser2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmsByUser(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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

	uid := uuid.New()
	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"film_id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.FilmsByUser(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_FilmStartSelection(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().FilmStartSelection(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.FilmStartSelection(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_SetRatingById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().SetRating(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	r.URL.Query().Set("rating", "2")

	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.SetRating(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_RatingById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().GetRating(gomock.Any(), gomock.Any()).Times(1).Return(prepare(), nil)
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
		"id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.GetRating(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_Random(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().Random(gomock.Any(), gomock.Any()).Times(1).Return(prepare(), nil)
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
		"id": uid.String(),
	})
	w := httptest.NewRecorder()

	handler.RandomFilm(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_AddStarred(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().AddStarred(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.AddStarred(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_AddWl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().AddWatchList(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.AddWatchlist(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_RemoveStarred(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().RemoveStarred(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.RemoveStarred(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_RemoveWl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().RemoveWatchList(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.RemoveWatchlist(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_IfStarred(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().IfStarred(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.IfStarred(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_IfWl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().IfWatchList(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Nothing{}, nil)
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
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.IfWl(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_IfStarred2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().Starred(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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

	uid := uuid.New()
	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.GetStarred(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}

func TestFilmsHandler_IfWl2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	usecase := mocks.NewMockFilmsServiceClient(ctl)
	usecase.EXPECT().WatchList(gomock.Any(), gomock.Any()).Times(1).Return(&generated.Films{
		Data: []*generated.Film{prepare()},
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

	uid := uuid.New()
	r := httptest.NewRequest("GET", "/film/"+uid.String(), strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"id": uid.String(),
	})
	t.Setenv("SECRET", "salt")
	enc := usecase2.NewTokenator()
	str := enc.GetToken(models.User{Id: uuid.New(), Login: "WTF"})
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  str,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()

	handler.GetWatchlist(w, r)
	require.Equal(t, w.Code, http.StatusOK)
	log.Print(w.Body.String())
}
