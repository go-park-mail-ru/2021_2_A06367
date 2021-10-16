package http
//
//import (
//	"fmt"
//	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
//	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
//	"github.com/golang/mock/gomock"
//	"github.com/gorilla/mux"
//	"github.com/stretchr/testify/require"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
//func TestFilmByGenre(t *testing.T) {
//
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	usecase := films.NewMockFilmsUsecase(ctl)
//	usecase.EXPECT().GetCompilation("topic").Times(1).Return([]models.Film{}, models.Okey)
//
//	handler := NewFilmsHandler(usecase)
//
//	r := httptest.NewRequest("GET", "/films/genres", strings.NewReader(fmt.Sprint()))
//	r = mux.SetURLVars(r, map[string]string{
//		"genre": "topic",
//	})
//	w := httptest.NewRecorder()
//
//	handler.FilmByGenre(w, r)
//	require.Equal(t, w.Code, http.StatusOK)
//	log.Print(w.Body.String())
//}
//
//func TestFilmBySelection(t *testing.T) {
//
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	usecase := films.NewMockFilmsUsecase(ctl)
//	usecase.EXPECT().GetSelection("hottest").Times(1).Return([]models.Film{}, models.Okey)
//
//	handler := NewFilmsHandler(usecase)
//
//	r := httptest.NewRequest("GET", "/films/hottest", strings.NewReader(fmt.Sprint()))
//	r = mux.SetURLVars(r, map[string]string{
//		"selection": "hottest",
//	})
//	w := httptest.NewRecorder()
//
//	handler.FilmBySelection(w, r)
//	require.Equal(t, w.Code, http.StatusOK)
//	log.Print(w.Body.String())
//}
