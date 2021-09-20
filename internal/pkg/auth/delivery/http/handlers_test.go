package http

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userOk := models.User{
		Login:             "Rocky",
		Email:             "d@mail.ru",
		EncryptedPassword: "ddd",
	}
	userjson, err := json.Marshal(&userOk)
	if err != nil {
		t.Error(err.Error())
	}

	userNOk := models.User{
		Login:             "Rocky",
		Email:             "d@mail.ru",
		EncryptedPassword: "ddd",
	}
	baduserjson, err := json.Marshal(&userNOk)
	if err != nil {
		t.Error(err.Error())
	}

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
				r:            httptest.NewRequest("POST", "/persons", bytes.NewReader(userjson)),
				statusReturn: models.Okey,
			},
		},
		{
			name:   "no user with such login etc",
			fields: fields{Usecase: mockUsecase},
			args: args{
				r:            httptest.NewRequest("POST", "/persons", bytes.NewReader(baduserjson)),
				statusReturn: models.Conflict,
			},
		},
	}
	mockUsecase.EXPECT().SignIn(userOk).Return(tests[0].args.statusReturn)
	mockUsecase.EXPECT().SignIn(userOk).Return(tests[1].args.statusReturn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := &AuthHandler{
				uc: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()

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
