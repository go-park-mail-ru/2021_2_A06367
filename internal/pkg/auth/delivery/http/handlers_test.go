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

func bodyPrepare(user models.User) []byte {
	userjson, err := json.Marshal(&user)
	if err != nil {
		return nil
	}
	return userjson
}

var testUsers []models.User = []models.User{
	models.User{
		Login:             "Phil",
		EncryptedPassword: "mancity",
		Email:             "phil@yandex.ru",
	},
	models.User{
		Login:             "Donald",
		EncryptedPassword: "maga",
		Email:             "usa@gmail.com",
	},
	models.User{
		Login:             "",
		EncryptedPassword: "",
		Email:             "",
	},
}

func TestAuthHandler_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := auth.NewMockAuthUsecase(ctl)
	mockOnlineRepo := auth.NewMockOnlineRepo(ctl)
	type fields struct {
		Usecase    auth.AuthUsecase
		OnlineRepo auth.OnlineRepo
	}

	type args struct {
		r            *http.Request
		result       http.Response
		statusReturn models.StatusCode
		OnlineStatus bool
		SetOnline    models.StatusCode
		SetOffline   models.StatusCode
	}

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[0]))),
				statusReturn: models.Okey,
				result:       http.Response{StatusCode: http.StatusOK},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[1]))),
				statusReturn: models.Unauthed,
				result:       http.Response{StatusCode: http.StatusConflict},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  "Anonymous",
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[2]))),
				statusReturn: models.Unauthed,
				result:       http.Response{StatusCode: http.StatusNotAcceptable},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		mockUsecase.EXPECT().
			SignIn(models.LoginUser{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}).
			Return("token", tests[i].args.statusReturn)
		mockOnlineRepo.EXPECT().IsOnline(testUsers[i]).Return(tests[i].args.OnlineStatus)
		mockOnlineRepo.EXPECT().UserOn(testUsers[i]).Return(models.Okey)
		mockOnlineRepo.EXPECT().UserOff(testUsers[i]).Return(models.Okey)
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				uc: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()

			h.Login(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthHandler_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := auth.NewMockAuthUsecase(ctl)
	mockOnlineRepo := auth.NewMockOnlineRepo(ctl)
	type fields struct {
		Usecase    auth.AuthUsecase
		OnlineRepo auth.OnlineRepo
	}

	type args struct {
		r            *http.Request
		result       http.Response
		statusReturn models.StatusCode
		OnlineStatus bool
		SetOnline    models.StatusCode
		SetOffline   models.StatusCode
	}

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[0]))),
				statusReturn: models.Okey,
				result:       http.Response{StatusCode: http.StatusOK},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[1]))),
				statusReturn: models.Conflict,
				result:       http.Response{StatusCode: http.StatusConflict},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  "Anonymous",
			fields: fields{Usecase: mockUsecase, OnlineRepo: mockOnlineRepo},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[2]))),
				statusReturn: models.NotFound,
				result:       http.Response{StatusCode: http.StatusNotAcceptable},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		mockUsecase.EXPECT().SignUp(testUsers[i]).Return("token", tests[i].args.statusReturn)
		mockOnlineRepo.EXPECT().IsOnline(testUsers[i]).Return(tests[i].args.OnlineStatus)
		mockOnlineRepo.EXPECT().UserOn(testUsers[i]).Return(models.Okey)
		mockOnlineRepo.EXPECT().UserOff(testUsers[i]).Return(models.Okey)
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				uc: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()

			h.Login(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {

}
