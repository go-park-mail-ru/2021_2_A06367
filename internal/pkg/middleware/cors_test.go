package middleware

import (
	"bytes"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type HandlerTest struct {
}

func prepare()  *zap.SugaredLogger{
	logger, err := zap.NewProduction()
	if err != nil {
		return nil
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)

	zapSugar := logger.Sugar()
	return zapSugar
}


func (t HandlerTest) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func TestCORSMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		{
			args: args{next: HandlerTest{}},
			want: HandlerTest{},
		},
	}

	r := httptest.NewRequest("POST", "/persons",
		bytes.NewReader(nil))
	w := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CORSMiddleware(tt.args.next); got == nil {
				t.Errorf("CORSMiddleware() = %v, want %v", got, tt.want)
			}
			got := CORSMiddleware(tt.args.next)
			got.ServeHTTP(w, r)
		})
	}
}

func TestLoggerMiddleware_LogRequest(t *testing.T) {
	type fields struct {
		logger *zap.SugaredLogger
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			args: args{next: HandlerTest{}},
			fields: fields{logger:prepare()},
		},
	}
	r := httptest.NewRequest("POST", "/persons",
		bytes.NewReader(nil))
	w := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMiddleware{
				logger: tt.fields.logger,
			}
			if got := m.LogRequest(tt.args.next); got == nil {
				t.Errorf("LogRequest() = %v", got)
			}
			got := m.LogRequest(tt.args.next)
			got.ServeHTTP(w, r)
		})
	}
}

func TestLoginUserIsValid(t *testing.T) {
	type args struct {
		user models.LoginUser
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{user: models.LoginUser{}},
			want: false,
		},
		{
			args: args{user: models.LoginUser{Login: "abc",EncryptedPassword: "abc"}},
			want: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoginUserIsValid(tt.args.user); got != tt.want {
				t.Errorf("LoginUserIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMetricsMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want *MetricsMiddleware
	}{
		{ },
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMetricsMiddleware();got == nil {
				t.Errorf("NewMetricsMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMiddleware(t *testing.T) {
	type args struct {
		logger *zap.SugaredLogger
	}
	tests := []struct {
		name string
		args args
		want LoggerMiddleware
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMiddleware(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWriter(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
		want *writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriter(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserIsValid(t *testing.T) {
	type args struct {
		user models.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{user: models.User{}},
			want: false,
		},
		{
			args: args{user: models.User{Login: "abc",EncryptedPassword: "abc"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserIsValid(tt.args.user); got != tt.want {
				t.Errorf("UserIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writer_WriteHeader(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		statusCode     int
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &writer{
				ResponseWriter: tt.fields.ResponseWriter,
				statusCode:     tt.fields.statusCode,
			}

			w.Header()
		})
	}
}
