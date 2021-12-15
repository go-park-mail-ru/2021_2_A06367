package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	grpc "github.com/go-park-mail-ru/2021_2_A06367/internal/models/grpc"
	generated2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	Usecase generated2.AuthServiceClient
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
	},
	models.User{
		Login:             "Donald",
		EncryptedPassword: "maga",
	},
	models.User{
		Login:             "Anonym",
		EncryptedPassword: "",
	},
	models.User{
		Login:             "Bad",
		EncryptedPassword: "User",
	},
	models.User{
		Login:             "Pom",
		EncryptedPassword: "Pom",
	},
}

func TestNewAuthHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		t.Error(err.Error())
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	zapSugar := logger.Sugar()
	testHandler := NewAuthHandler(mockUsecase, zapSugar)
	if testHandler.client != mockUsecase {
		t.Error("bad constructor")
	}
}

func TestAuthHandler_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase},
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
			fields: fields{Usecase: mockUsecase},
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
			fields: fields{Usecase: mockUsecase},
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
		if tests[i].args.statusReturn != models.Forbidden {
			log.Print(LoginUserCopy)
		}
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				client: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()
			if tt.args.statusReturn != models.Forbidden {
				mockUsecase.EXPECT().Login(gomock.Any(), gomock.Any()).Return(&generated2.Token{
					Status: grpc.StatusCode(tt.args.statusReturn),
				}, nil).Times(1)
			}
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

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{Usecase: mockUsecase},
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
			fields: fields{Usecase: mockUsecase},
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
			fields: fields{Usecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/persons",
					bytes.NewReader([]byte("d"))),
				statusReturn: models.BadRequest,
				result:       http.Response{StatusCode: http.StatusBadRequest},
				OnlineStatus: false,
			},
		},
	}

	mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(&generated2.Token{
		Status: grpc.StatusCode(tests[0].args.statusReturn), Cookie: "hello"}, nil).Times(1)
	mockUsecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(&generated2.Token{
		Status: grpc.StatusCode(tests[1].args.statusReturn), Cookie: "hello"}, nil).Times(1)

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{
				client: tt.fields.Usecase,
			}
			w := httptest.NewRecorder()

			enc := usecase.NewTokenator()

			str := enc.GetToken(models.User{Id: uuid.New()})
			SSCookie := &http.Cookie{
				Name:     "SSID",
				Value:    str,
				Path:     "/",
				Domain:   "a06367.ru",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour * 24),
			}

			tt.args.r.AddCookie(SSCookie)
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

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{},
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
		{
			Login:  testUsers[1].Login,
			fields: fields{},
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
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{}
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
	tkn := &usecase.Tokenator{}
	bdy := tkn.GetToken(models.User{Login: "Phi", Id: uuid.New()})
	badBody, err := json.Marshal(models.TokenView{Token: bdy})
	if err != nil {
		t.Error(err)
	}
	//bdyOK := tkn.GetToken(models.User{Login: "Phil", Id: uuid.New()})
	//goodBody, _ := json.Marshal(models.TokenView{Token: bdyOK})

	tests := []struct {
		Login  string
		body   []byte
		fields fields
		args   args
	}{
		//{
		//	Login:  testUsers[0].Login,
		//	fields: fields{},
		//	args: args{
		//		r: httptest.NewRequest("GET", "/auth?user=",
		//			bytes.NewReader([]byte(""))),
		//		statusReturn: models.BadRequest,
		//		result:       http.Response{StatusCode: http.StatusBadRequest},
		//		OnlineStatus: false,
		//		SetOnline:    models.Okey,
		//		SetOffline:   models.Okey,
		//	},
		//},
		{
			Login:  testUsers[1].Login,
			fields: fields{},
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
		//{
		//	Login:  testUsers[2].Login,
		//	fields: fields{},
		//	args: args{
		//		r: httptest.NewRequest("GET", "/user/auth?user=Phil",
		//			bytes.NewReader(goodBody)),
		//		statusReturn: models.Okey,
		//		result:       http.Response{StatusCode: http.StatusOK},
		//		OnlineStatus: true,
		//		SetOnline:    models.Okey,
		//		SetOffline:   models.Okey,
		//	},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthHandler{}

			w := httptest.NewRecorder()
			h.AuthStatus(w, tt.args.r)
			if tt.args.result.StatusCode != w.Code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthHandler_GetProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	handler := NewAuthHandler(nil, nil)

	handler.GetProfile(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("wrong result")
	}
}

func TestAuthHandler_GetProfile2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "salt")

	data := models.Profile{Id: uuid.New()}
	js, err := json.Marshal(data)
	if err != nil {
		t.Skip()
	}
	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(string(js)))
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	enc := usecase.NewTokenator()

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
	mockUsecase := generated2.NewMockAuthServiceClient(ctl)
	mockUsecase.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(&generated2.Profile{}, nil)
	handler := NewAuthHandler(mockUsecase, nil)

	handler.GetProfile(w, r)
	if w.Code != http.StatusOK {
		t.Error("wrong result")
	}
}

func TestAuthHandler_Follow(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(nil, nil)

	handler.Follow(w, r)
	if w.Code != http.StatusOK {
		t.Error("wrong result")
	}
}

func TestAuthHandler_Unfollow(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(nil, nil)

	handler.Unfollow(w, r)
	if w.Code != http.StatusOK {
		t.Error("wrong result")
	}
}

func TestAuthHandler_UpdateProfilePic(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	data := models.PassUpdate{
		Password: "data",
	}
	js, err := json.Marshal(data)
	if err != nil {
		t.Skip()
	}

	r := httptest.NewRequest("GET", "/", strings.NewReader(string(js)))
	os.Setenv("SECRET", "salt")
	enc := usecase.NewTokenator()

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
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(mockUsecase, nil)

	//mockUsecase.EXPECT().UpdateProfilePic(gomock.Any(), gomock.Any()).Return(&generated2.Empty{Status: grpc.StatusCode_Okey}, nil).Times(1)
	handler.UpdateProfilePic(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("wrong result")
	}
}

func TestAuthHandler_UpdateProfilePass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	data := models.PassUpdate{
		Password: "data",
	}
	js, err := json.Marshal(data)
	if err != nil {
		t.Skip()
	}

	r := httptest.NewRequest("GET", "/", strings.NewReader(string(js)))
	os.Setenv("SECRET", "salt")
	enc := usecase.NewTokenator()

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
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(mockUsecase, nil)

	mockUsecase.EXPECT().UpdateProfilePass(gomock.Any(), gomock.Any()).Return(&generated2.Empty{Status: grpc.StatusCode_Okey}, nil).Times(1)
	handler.UpdateProfilePass(w, r)
	if w.Code != http.StatusOK {
		t.Error("wrong result")
	}
}

func TestAuthHandler_UpdateProfilePassErr(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	data := models.PassUpdate{
		Password: "data",
	}
	js, err := json.Marshal(data)
	if err != nil {
		t.Skip()
	}

	r := httptest.NewRequest("GET", "/", strings.NewReader(string(js)))
	os.Setenv("SECRET", "salt")
	enc := usecase.NewTokenator()

	str := enc.GetToken(models.User{Id: uuid.New()})
	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    str,
		Path:     "/",
		Domain:   "a06367.ru",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	r.AddCookie(SSCookie)
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(mockUsecase, nil)

	//mockUsecase.EXPECT().UpdateProfilePass(gomock.Any(), gomock.Any()).Return(&generated2.Empty{Status: grpc.StatusCode_Okey}, nil).Times(1)
	handler.UpdateProfilePass(w, r)
	if w.Code != http.StatusUnauthorized {
		t.Error("wrong result")
	}
}
func TestAuthHandler_UpdateProfileBio(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	data := models.PassUpdate{
		Password: "data",
	}
	js, err := json.Marshal(data)
	if err != nil {
		t.Skip()
	}

	r := httptest.NewRequest("GET", "/", strings.NewReader(string(js)))
	os.Setenv("SECRET", "salt")
	enc := usecase.NewTokenator()

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
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(mockUsecase, nil)

	mockUsecase.EXPECT().UpdateProfileBio(gomock.Any(), gomock.Any()).Return(&generated2.Empty{Status: grpc.StatusCode_Okey}, nil).Times(1)
	handler.UpdateProfileBio(w, r)
	if w.Code != http.StatusOK {
		t.Error("wrong result")
	}
}

func TestToken(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := generated2.NewMockAuthServiceClient(ctl)

	r := httptest.NewRequest("GET", "/selection/user/personal", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, map[string]string{
		"id": "uid.String()",
	})
	handler := NewAuthHandler(mockUsecase, nil)

	handler.Token(w, r)
}
