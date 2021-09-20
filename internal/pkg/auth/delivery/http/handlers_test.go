package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthHandler_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := auth.NewMockAuthUsecase(ctl)

	type fields struct {
		Usecase auth.AuthUsecase
	}

	type args struct {
		r            *http.Request
		result       http.Response
		status       int
		statusReturn models.StatusCode
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple create",
			fields: fields{Usecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" }`, "name"))),
				status:       http.StatusCreated,
				statusReturn: models.Okey,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := &AuthHandler{
				uc: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()

			mockUsecase.EXPECT().SignIn().Return(uint(0), tt.args.statusReturn)

			h.Login(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error(tt.name)
			}
		})
	}
}

func TestAuthHandler_SignUp(t *testing.T) {

}

func TestAuthHandler_Logout(t *testing.T) {

}
