package delivery

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	mocks2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/mocks"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	mocks3 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewSearchHandler(t *testing.T) {
	type args struct {
		fu films.FilmsUsecase
		pu auth.AuthUsecase
		au actors.ActorsUsecase
	}
	tests := []struct {
		name string
		args args
		want *SearchHandler
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSearchHandler(tt.args.fu, tt.args.pu, tt.args.au)
			if got == nil {
				t.Error("nil constructor")
			}
		})
	}
}

func TestSearchHandler_Search(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	f := mocks.NewMockFilmsUsecase(ctl)
	f.EXPECT().GetByKeyword(gomock.Any()).Return([]models.Film{}, models.Okey)
	a := mocks2.NewMockActorsUsecase(ctl)
	a.EXPECT().GetByKeyword(gomock.Any()).Return([]models.Actors{}, models.Okey)
	at := mocks3.NewMockAuthUsecase(ctl)
	at.EXPECT().GetByKeyword(gomock.Any()).Return([]models.Profile{}, models.Okey)

	type fields struct {
		fu films.FilmsUsecase
		pu auth.AuthUsecase
		au actors.ActorsUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{{
		fields: fields{
			fu: f,
			pu: at,
			au: a,
		},
	},
	}

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"keyword": uuid.New().String(),
	})
	w := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &SearchHandler{
				fu: tt.fields.fu,
				pu: tt.fields.pu,
				au: tt.fields.au,
			}

			sh.Search(w, r)
		})
	}
}

func TestSearchHandler_SearchFail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	f := mocks.NewMockFilmsUsecase(ctl)
	a := mocks2.NewMockActorsUsecase(ctl)
	at := mocks3.NewMockAuthUsecase(ctl)
	type fields struct {
		fu films.FilmsUsecase
		pu auth.AuthUsecase
		au actors.ActorsUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{{
		fields: fields{
			fu: f,
			pu: at,
			au: a,
		},
	},
	}

	r := httptest.NewRequest("GET", "/actor/id", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &SearchHandler{
				fu: tt.fields.fu,
				pu: tt.fields.pu,
				au: tt.fields.au,
			}

			sh.Search(w, r)
		})
	}
}
