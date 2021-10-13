package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"os"
	"testing"
)

type fields struct {
	AuthRepo auth.AuthRepo
	TokenGen auth.TokenGenerator
	Encrypt  auth.Encrypter
}

type args struct {
	statusReturn models.StatusCode
	OnlineStatus models.StatusCode
	SetOnline    models.StatusCode
	SetOffline   models.StatusCode
}

var testUsers []models.User = []models.User{
	models.User{
		Id:                uuid.New(),
		Login:             "Phil",
		EncryptedPassword: "mancity",
		Email:             "phil@yandex.ru",
	},
	models.User{
		Id:                uuid.New(),
		Login:             "Donald",
		EncryptedPassword: "maga",
		Email:             "usa@gmail.com",
	},
	models.User{
		Id:                uuid.New(),
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
	mockEncrypter := auth.NewMockEncrypter(ctl)
	testUC := NewAuthUsecase(mockAuthRepo, mockTokenGen, mockEncrypter)
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
	mockEncrypter := auth.NewMockEncrypter(ctl)

	tests := []struct {
		Login  string
		fields fields
		args   args
	}{
		{
			Login:  testUsers[0].Login,
			fields: fields{mockAuthRepo, mockTokenGen, mockEncrypter},
			args:   args{statusReturn: models.Okey, OnlineStatus: models.Okey},
		},
		{
			Login:  testUsers[1].Login,
			fields: fields{mockAuthRepo, mockTokenGen, mockEncrypter},
			args:   args{statusReturn: models.Unauthed, OnlineStatus: models.Unauthed},
		},
		{
			Login:  testUsers[2].Login,
			fields: fields{mockAuthRepo, mockTokenGen, mockEncrypter},
			args:   args{statusReturn: models.BadRequest},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.BadRequest {
			continue
		}

		mockEncrypter.EXPECT().EncryptPswd(testUsers[i].EncryptedPassword).Return(testUsers[i].EncryptedPassword)
		mockTokenGen.EXPECT().GetToken(testUsers[i]).Return("TEST TOKEN")
		mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[i].Login, EncryptedPassword: testUsers[i].EncryptedPassword}).Return(testUsers[i], tests[i].args.statusReturn)
	}

	for i, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthUsecase{
				repo:      mockAuthRepo,
				tokenator: mockTokenGen,
				encrypter: mockEncrypter,
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
	mockEncrypter := auth.NewMockEncrypter(ctl)

	tests := []struct {
		Login       string
		fields      fields
		args        args
		returnToken string
	}{
		{
			Login:       testUsers[0].Login,
			fields:      fields{mockAuthRepo, mockTokenGenereator, mockEncrypter},
			args:        args{statusReturn: models.Conflict, OnlineStatus: models.Okey},
			returnToken: "TEST TOKEN",
		},
		{
			Login:       testUsers[1].Login,
			fields:      fields{mockAuthRepo, mockTokenGenereator, mockEncrypter},
			args:        args{statusReturn: models.Okey, OnlineStatus: models.Unauthed},
			returnToken: "TEST TOKEN",
		},
		{
			Login:       testUsers[2].Login,
			fields:      fields{mockAuthRepo, mockTokenGenereator, mockEncrypter},
			args:        args{statusReturn: models.Conflict, OnlineStatus: models.Unauthed},
			returnToken: "",
		},
	}

	for i := 0; i < len(tests); i++ {
		CreateUser := models.User{
			Login:             testUsers[i].Login,
			EncryptedPassword: testUsers[i].EncryptedPassword,
			Email:             testUsers[i].Email,
		}
		if tests[i].args.statusReturn == models.BadRequest {
			continue
		}

		if tests[i].args.statusReturn == models.Conflict && tests[i].args.OnlineStatus == models.Okey {
			mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[i].Login, Email: testUsers[i].Email,
				EncryptedPassword: testUsers[i].EncryptedPassword}).Return(testUsers[i], tests[i].args.OnlineStatus)
			continue
		}
		mockAuthRepo.EXPECT().CreateUser(CreateUser).Return(testUsers[i], tests[i].args.statusReturn)
		mockTokenGenereator.EXPECT().GetToken(testUsers[i]).Return(tests[i].returnToken)
		mockEncrypter.EXPECT().EncryptPswd(testUsers[i].EncryptedPassword).Return(testUsers[i].EncryptedPassword)
		mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[i].Login, Email: testUsers[i].Email, EncryptedPassword: testUsers[i].EncryptedPassword}).Return(testUsers[i], tests[i].args.OnlineStatus)
	}

	for i, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthUsecase{
				repo:      mockAuthRepo,
				tokenator: mockTokenGenereator,
				encrypter: mockEncrypter,
			}
			if tt.Login == "Anonym" {
				tt.Login = tt.Login
			}
			CreateUser := models.User{
				Login:             testUsers[i].Login,
				EncryptedPassword: testUsers[i].EncryptedPassword,
				Email:             testUsers[i].Email,
			}

			_, code := h.SignUp(CreateUser)
			if tt.args.statusReturn != code {
				t.Error(tt.Login)
			}
		})
	}
}
