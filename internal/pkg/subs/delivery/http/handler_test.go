package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	generated2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	usecase2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewSubsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		subsClient generated.SubsServiceClient
		cl         generated2.AuthServiceClient
	}
	tests := []struct {
		name string
		args args
		want *SubsHandler
	}{
		{
			name: "simple check",
			args: args{
				subsClient: generated.NewMockSubsServiceClient(ctrl),
				cl:         generated2.NewMockAuthServiceClient(ctrl),
			},
			want: &SubsHandler{
				subsClient: generated.NewMockSubsServiceClient(ctrl),
				cl:         generated2.NewMockAuthServiceClient(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubsHandler(tt.args.subsClient, tt.args.cl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsHandler_GetLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := generated.NewMockSubsServiceClient(ctrl)
	a.EXPECT().GetLicense(gomock.Any(), gomock.Any()).Return(&generated.License{IsValid: true,
		ExpiresDate: time.Now().Format(time.RFC3339)}, nil)
	b := generated2.NewMockAuthServiceClient(ctrl)

	r := httptest.NewRequest("GET", "/abcd/", strings.NewReader(fmt.Sprint()))
	os.Setenv("SECRET", "salt")
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

	type fields struct {
		subsClient generated.SubsServiceClient
		cl         generated2.AuthServiceClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "simple check",
			fields: fields{
				subsClient: a,
				cl:         b,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := SubsHandler{
				subsClient: tt.fields.subsClient,
				cl:         tt.fields.cl,
			}
			h.GetLicense(w, r)
		})
	}
}

func TestSubsHandler_SetLicense(t *testing.T) {
	/*ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := generated.NewMockSubsServiceClient(ctrl)
	a.EXPECT().SetLicense(gomock.Any(), gomock.Any()).Return(&generated.License{IsValid: true,
		ExpiresDate: time.Now().Format(time.RFC3339)}, nil)
	b := generated2.NewMockAuthServiceClient(ctrl)
	b.EXPECT().CheckByLogin(gomock.Any(), gomock.Any()).Return(&generated2.UserUUID{ID: uuid.New().String()}, nil)
	pr, pw := io.Pipe()
	writer := part.NewWriter(pw)
	go func() {
		defer writer.Close()
		part, err := writer.CreateFormField("label")
		if err != nil {
			return
		}
		write, err := part.Write([]byte("test"))
		if err != nil || write == 0 {
			return
		}
	}()

	r := httptest.NewRequest("GET", "/abcd/", pr)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	os.Setenv("SECRET", "salt")
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
	defer ctrl.Finish()

	type fields struct {
		subsClient generated.SubsServiceClient
		cl         generated2.AuthServiceClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "simple check",
			fields: fields{
				subsClient: a,
				cl:         b,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := SubsHandler{
				subsClient: tt.fields.subsClient,
				cl:         tt.fields.cl,
			}
			h.SetLicense(w, r)
		})
	}
	 */
}
