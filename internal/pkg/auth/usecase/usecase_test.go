package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

type fields struct {
	AuthRepo auth.AuthRepo
	TokenGen auth.TokenGenerator
}

type args struct {
	statusReturn models.StatusCode
	OnlineStatus models.StatusCode
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
}

func TestNewAuthUsecase(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := auth.NewMockAuthRepo(ctl)
	mockTokenGen := auth.NewMockTokenGenerator(ctl)
	testUC := NewAuthUsecase(mockAuthRepo, mockTokenGen)
	if testUC.repo != mockAuthRepo {
		t.Error("bad constructor")
	}

	if testUC.tokenator != mockTokenGen {
		t.Error("bad constructor")
	}
}

func TestAuthUsecase_SignIn(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "TESTS")
	mockAuthRepo := auth.NewMockAuthRepo(ctl)
	mockTokenGen := auth.NewMockTokenGenerator(ctl)

	tests := []struct {
		Login  string
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{mockAuthRepo, mockTokenGen},
			args:   args{statusReturn: models.Okey, OnlineStatus: models.Okey},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{mockAuthRepo, mockTokenGen},
			args:   args{statusReturn: models.Unauthed, OnlineStatus: models.Unauthed},
		},
		{
			Login:  testUsers[2].Login,
			fields: fields{mockAuthRepo, mockTokenGen},
			args:   args{statusReturn: models.BadRequest},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.BadRequest {
			continue
		}
		if tests[i].args.statusReturn == models.Okey {
			mockTokenGen.EXPECT().GetToken(models.User{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}).Return("TEST TOKEN")
		}
		mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}).Return(tests[i].args.statusReturn)
	}

	for i, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthUsecase{
				repo:      mockAuthRepo,
				tokenator: mockTokenGen,
			}

			_, code := h.SignIn(models.LoginUser{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword})
			if tt.args.statusReturn != code {
				t.Error(tt.Login)
			}
		})
	}

}

func TestAuthUsecase_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "TESTS")
	mockAuthRepo := auth.NewMockAuthRepo(ctl)
	mockTokenGenereator := auth.NewMockTokenGenerator(ctl)

	tests := []struct {
		Login  string
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{mockAuthRepo, mockTokenGenereator},
			args:   args{statusReturn: models.Conflict, OnlineStatus: models.Okey},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{mockAuthRepo, mockTokenGenereator},
			args:   args{statusReturn: models.Okey, OnlineStatus: models.Unauthed},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.BadRequest {
			continue
		}
		if tests[i].args.statusReturn == models.Okey {
			mockAuthRepo.EXPECT().CreateUser(testUsers[i]).Return(tests[i].args.statusReturn)
			mockTokenGenereator.EXPECT().GetToken(testUsers[i]).Return("TEST TOKEN")
		}
		mockAuthRepo.EXPECT().CheckUser(testUsers[i]).Return(tests[i].args.OnlineStatus)
	}

	for i, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthUsecase{
				repo:      mockAuthRepo,
				tokenator: mockTokenGenereator,
			}

			_, code := h.SignUp(testUsers[i])
			if tt.args.statusReturn != code {
				t.Error(tt.Login)
			}
		})
	}

}
