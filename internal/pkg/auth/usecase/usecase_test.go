package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/mocks"
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
	},
	models.User{
		Id:                uuid.New(),
		Login:             "Donald",
		EncryptedPassword: "maga",
	},
	models.User{
		Id:                uuid.New(),
		Login:             "Anonym",
		EncryptedPassword: "",
	},
}

func TestNewAuthUsecase(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockTokenGen := mocks.NewMockTokenGenerator(ctl)
	mockEncrypter := mocks.NewMockEncrypter(ctl)
	testUC := NewAuthUsecase(mockAuthRepo, mockTokenGen, mockEncrypter, nil)
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
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockTokenGen := mocks.NewMockTokenGenerator(ctl)
	mockEncrypter := mocks.NewMockEncrypter(ctl)

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
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockTokenGenereator := mocks.NewMockTokenGenerator(ctl)
	mockEncrypter := mocks.NewMockEncrypter(ctl)

	tests := []struct {
		Login       string
		fields      fields
		args        args
		returnToken string
	}{
		{
			Login:       testUsers[0].Login,
			fields:      fields{mockAuthRepo, mockTokenGenereator, mockEncrypter},
			args:        args{statusReturn: models.Okey, OnlineStatus: models.InternalError},
			returnToken: "TEST TOKEN",
		},
		{
			Login:       testUsers[1].Login,
			fields:      fields{mockAuthRepo, mockTokenGenereator, mockEncrypter},
			args:        args{statusReturn: models.Conflict, OnlineStatus: models.Unauthed},
			returnToken: "",
		},
	}
	mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[0].Login, EncryptedPassword: testUsers[0].EncryptedPassword}).
		Return(testUsers[0], tests[0].args.OnlineStatus).Times(1)
	mockEncrypter.EXPECT().EncryptPswd(testUsers[0].EncryptedPassword).Return(testUsers[0].EncryptedPassword).Times(1)

	mockAuthRepo.EXPECT().CheckUser(models.User{Login: testUsers[1].Login, EncryptedPassword: testUsers[1].EncryptedPassword}).
		Return(testUsers[1], tests[1].args.OnlineStatus).Times(1)

	mockEncrypter.EXPECT().EncryptPswd(testUsers[1].EncryptedPassword).Return(testUsers[1].EncryptedPassword).Times(1)
	mockAuthRepo.EXPECT().CreateUser(gomock.Any()).Return(testUsers[0], tests[0].args.statusReturn).Times(1)
	mockTokenGenereator.EXPECT().GetToken(gomock.Any()).Return(tests[0].returnToken).Times(1)

	for i, tt := range tests {
		t.Run(tt.Login, func(t *testing.T) {
			h := &AuthUsecase{
				repo:      mockAuthRepo,
				tokenator: mockTokenGenereator,
				encrypter: mockEncrypter,
			}
			CreateUser := models.User{
				Login:             testUsers[i].Login,
				EncryptedPassword: testUsers[i].EncryptedPassword,
			}

			_, code := h.SignUp(CreateUser)
			if tt.args.statusReturn != code {
				t.Error(tt.Login)
			}
		})
	}
}

func TestAuthUsecase_Follow(t *testing.T) {
	who := uuid.New()
	whom := uuid.New()
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "TESTS")
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().AddFollowing(who, whom).Return(models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	st := usecase.Follow(who, whom)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestAuthUsecase_Unfollow(t *testing.T) {
	who := uuid.New()
	whom := uuid.New()
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "TESTS")
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().RemoveFollowing(who, whom).Return(models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	st := usecase.Unfollow(who, whom)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}


func TestAuthUsecase_GetProfile(t *testing.T) {
	who := uuid.New()
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	os.Setenv("SECRET", "TESTS")
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().GetProfile(gomock.Any()).Return(models.Profile{}, models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	_, st := usecase.GetProfile(models.Profile{Id: who})
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestAuthUsecase_GetSubscribers(t *testing.T) {
	usecase := NewAuthUsecase(nil, nil, nil, nil)
	_, st := usecase.GetSubscribers()
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestAuthUsecase_GetSubscriptions(t *testing.T) {
	usecase := NewAuthUsecase(nil, nil, nil, nil)
	_, st := usecase.GetSubscriptions()
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}


func TestAuthUsecase_GetByKeyword(t *testing.T) {
	keyword := "test"
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().GetProfileByKeyword(keyword).Return(nil, models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	_, st := usecase.GetByKeyword(keyword)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}


func TestAuthUsecase_UpdateBio(t *testing.T) {
	data := models.Profile{}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().UpdateBio(gomock.Any()).Return( models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	st := usecase.SetBio(data)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestAuthUsecase_UpdateAvatar(t *testing.T) {
	data := models.Profile{}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockAuthRepo.EXPECT().UpdateAvatar(gomock.Any()).Return( models.Okey)
	usecase := NewAuthUsecase(mockAuthRepo, nil, nil, nil)
	st := usecase.SetAvatar(data)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestAuthUsecase_UpdatePass(t *testing.T) {
	data := models.User{}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockAuthRepo := mocks.NewMockAuthRepo(ctl)
	mockEncrypter := mocks.NewMockEncrypter(ctl)
	mockAuthRepo.EXPECT().UpdatePass(gomock.Any()).Return( models.Okey)
	mockEncrypter.EXPECT().EncryptPswd(gomock.Any()).Return("password")
	usecase := NewAuthUsecase(mockAuthRepo, nil, mockEncrypter, nil)
	st := usecase.SetPass(data)
	if st != models.Okey {
		t.Error("wrong status code returned")
	}
}

func TestNewEncrypter(t *testing.T) {
	t.Setenv("SECRET", "salt")
	enc := NewEncrypter()
	if enc == nil {
		t.Error("nothing here")
	}
}


func TestNewEncrypterPass(t *testing.T) {
	t.Setenv("SECRET", "salt")
	enc := NewEncrypter()

	str := enc.EncryptPswd("")
	if str == "" {
		t.Error("nothing here")
	}
}


func TestNewTokenator(t *testing.T) {
	t.Setenv("SECRET", "salt")
	enc := NewTokenator()
	if enc == nil {
		t.Error("nothing here")
	}
}

func TestNewTokenatorGet(t *testing.T) {
	t.Setenv("SECRET", "salt")
	enc := NewTokenator()

	str := enc.GetToken(models.User{})
	if str == "" {
		t.Error("nothing here")
	}
}
