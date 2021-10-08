package http

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func bodyPrepare(user models.User) []byte {
	userjson, err := json.Marshal(&user)
	if err != nil {
		return nil
	}
	return userjson
}

type fields struct {
	Usecase       auth.AuthUsecase
	OnlineUsecase auth.OnlineUsecase
}

type args struct {
	r            *http.Request
	result       http.Response
	statusReturn models.StatusCode
	OnlineStatus bool
	SetOnline    models.StatusCode
	SetOffline   models.StatusCode
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
		Login:             "Anonym",
		EncryptedPassword: "",
		Email:             "",
	},
	models.User{
		Login:             "Bad",
		EncryptedPassword: "User",
		Email:             "KKK",
	},
	models.User{
		Login:             "Pom",
		EncryptedPassword: "Pom",
		Email:             "Pom",
	},
}

func TestNewAuthHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := auth.NewMockAuthUsecase(ctl)
	mockOnlineUsecase := auth.NewMockOnlineUsecase(ctl)

	testHandler := NewAuthHandler(mockUsecase, mockOnlineUsecase)
	if testHandler.uc != mockUsecase {
		t.Error("bad constructor")
	}

	if testHandler.online != mockOnlineUsecase {
		t.Error("bad constructor")
	}
}

func TestAuthHandler_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := auth.NewMockAuthUsecase(ctl)
	mockOnlineUsecase := auth.NewMockOnlineUsecase(ctl)

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
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
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[1]))),
				statusReturn: models.Unauthed,
				result:       http.Response{StatusCode: http.StatusUnauthorized},
				OnlineStatus: false,
				SetOnline:    models.Unauthed,
			},
		},
		{
			Login:  testUsers[2].Login,
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader([]byte("Hi there"))),
				statusReturn: models.Forbidden,
				result:       http.Response{StatusCode: http.StatusForbidden},
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		LoginUserCopy := models.LoginUser{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}
		if tests[i].args.statusReturn == models.Okey {
			mockOnlineUsecase.EXPECT().Activate(LoginUserCopy).Return(tests[i].args.statusReturn)
		}
		if tests[i].args.statusReturn != models.Forbidden {
			mockUsecase.EXPECT().
				SignIn(LoginUserCopy).
				Return("", tests[i].args.statusReturn)
		}
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				uc:     tt.fields.Usecase,
				online: tt.fields.OnlineUsecase,
			}
			if tt.Login == testUsers[1].Login {
				tt.Login = tt.Login
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
	mockOnlineUsecase := auth.NewMockOnlineUsecase(ctl)

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
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
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader(bodyPrepare(testUsers[1]))),
				statusReturn: models.Conflict,
				result:       http.Response{StatusCode: http.StatusConflict},
				OnlineStatus: false,
			},
		},
		{
			Login:  testUsers[2].Login,
			fields: fields{Usecase: mockUsecase, OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader([]byte("d"))),
				statusReturn: models.BadRequest,
				result:       http.Response{StatusCode: http.StatusBadRequest},
				OnlineStatus: false,
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.BadRequest {
			continue
		}
		if tests[i].args.statusReturn == models.Okey {
			mockUsecase.EXPECT().SignUp(testUsers[i]).Return("token", tests[i].args.statusReturn)
			mockOnlineUsecase.EXPECT().
				Activate(models.LoginUser{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}).
				Return(tests[i].args.statusReturn)
			continue
		}
		mockUsecase.EXPECT().
			SignUp(models.User{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword, Email: testUsers[i].Email}).
			Return("", tests[i].args.statusReturn)
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				uc:     tt.fields.Usecase,
				online: tt.fields.OnlineUsecase,
			}
			w := httptest.NewRecorder()

			h.SignUp(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	os.Setenv("SECRET", "TEST")
	tkn := &usecase.Tokenator{}
	bdy := tkn.GetToken(models.User{Login: testUsers[1].Login, Id: uuid.New()})

	mockOnlineUsecase := auth.NewMockOnlineUsecase(ctl)

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					nil),
				statusReturn: models.BadRequest,
				result:       http.Response{StatusCode: http.StatusBadRequest},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					nil),
				statusReturn: models.Okey,
				result:       http.Response{StatusCode: http.StatusOK},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		LoginUserCopy := models.LoginUser{Login: testUsers[i].Login}
		if tests[i].args.statusReturn == models.BadRequest {
			tests[i].args.r.AddCookie(&http.Cookie{
				Name:     "SSID",
				Value:    "bdy",
				Expires:  time.Time{},
				HttpOnly: true,
			})
			continue
		}
		tests[i].args.r.AddCookie(&http.Cookie{
			Name:     "SSID",
			Value:    bdy,
			Expires:  time.Time{},
			HttpOnly: true,
		})
		mockOnlineUsecase.EXPECT().Deactivate(LoginUserCopy).Return(tests[i].args.statusReturn)
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				online: tt.fields.OnlineUsecase,
			}

			if tt.Login == testUsers[1].Login {
				tt.Login = tt.Login
			}
			w := httptest.NewRecorder()
			h.Logout(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthHandler_AuthStatus(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	os.Setenv("SECRET", "TEST")
	mockOnlineUsecase := auth.NewMockOnlineUsecase(ctl)
	tkn := &usecase.Tokenator{}
	bdy := tkn.GetToken(models.User{Login: "Phi", Id: uuid.New()})
	badBody, _ := json.Marshal(models.TokenView{Token: bdy})
	bdyOK := tkn.GetToken(models.User{Login: "Phil", Id: uuid.New()})
	goodBody, _ := json.Marshal(models.TokenView{Token: bdyOK})

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("GET", "/auth?user=",
					bytes.NewReader([]byte(""))),
				statusReturn: models.BadRequest,
				result:       http.Response{StatusCode: http.StatusBadRequest},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("GET", "/user/auth?user=Phil",
					bytes.NewReader(badBody)),
				statusReturn: models.Unauthed,
				result:       http.Response{StatusCode: http.StatusUnauthorized},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[2].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("GET", "/user/auth?user=Phil",
					bytes.NewReader(goodBody)),
				statusReturn: models.Okey,
				result:       http.Response{StatusCode: http.StatusOK},
				OnlineStatus: true,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
		{
			Login:  testUsers[3].Login,
			fields: fields{OnlineUsecase: mockOnlineUsecase},
			args: args{
				r: httptest.NewRequest("GET", "/user/auth?user=Phil",
					bytes.NewReader(goodBody)),
				statusReturn: models.Okey,
				result:       http.Response{StatusCode: http.StatusUnauthorized},
				OnlineStatus: false,
				SetOnline:    models.Okey,
				SetOffline:   models.Okey,
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn != models.BadRequest && tests[i].args.statusReturn != models.Unauthed {
			LoginUserCopy := models.LoginUser{Login: "Phil"}
			mockOnlineUsecase.EXPECT().IsAuthed(LoginUserCopy).Return(tests[i].args.OnlineStatus)
		}
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				online: tt.fields.OnlineUsecase,
			}

			w := httptest.NewRecorder()
			h.AuthStatus(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}
