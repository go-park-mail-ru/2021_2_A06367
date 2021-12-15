package repo

import (
	"errors"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const pswd = "test"

func prepare(t *testing.T, pgxRows pgx.Rows) *AuthRepo { // mock on repo and uid used as uuid for everything that has any uid filed
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	return repo
}

func TestAuthRepo_CheckUserLogin(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id", "encrypted_password"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, ecnryptedPswd).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, Login: "tester2", EncryptedPassword: ecnryptedPswd}
	_, st := repo.CheckUserLogin(user)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
}


func TestAuthRepo_CheckUser(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id", "encrypted_password"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, ecnryptedPswd).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, Login: "tester2", EncryptedPassword: ecnryptedPswd}
	userCheck, st := repo.CheckUser(user)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, user.EncryptedPassword, userCheck.EncryptedPassword)
}

func TestAuthRepo_CheckUserPassError(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id", "encrypted_password"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, ecnryptedPswd).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, Login: "tester2", EncryptedPassword: "abcd"}
	_, st := repo.CheckUser(user)

	if st != models.Unauthed {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_CheckUserInternalError(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id", "encrypted_password"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, ecnryptedPswd).RowError(0, errors.New("")).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, Login: "tester2", EncryptedPassword: "abcd"}
	_, st := repo.CheckUser(user)

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_CreateUser(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, EncryptedPassword: ecnryptedPswd}
	userCheck, st := repo.CreateUser(user)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, user.Id, userCheck.Id)
}

func TestAuthRepo_CreateUserErr(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid).RowError(0, errors.New("")).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, EncryptedPassword: ecnryptedPswd}
	_, st := repo.CreateUser(user)

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_AddFollowing(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	st := repo.AddFollowing(uid, uuid.New())

	if st != models.Okey {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_AddFollowingErr(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).RowError(0, errors.New("")).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	st := repo.AddFollowing(uid, uuid.New())

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_GetProfile(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"login", "about", "avatar", "subscriptions", "subscribers"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow("tester2",
			"new user", "img.png", uint(12),
			uint(10)).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	profile := models.Profile{Id: uid, Login: "tester2"}
	userCheck, st := repo.GetProfile(profile)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, profile.Login, userCheck.Login)
}

func TestAuthRepo_GetProfileErr(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"login", "about", "avatar", "subscriptions", "subscribers"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow("tester2",
			"new user", "img.png", uint(12),
			uint(10)).RowError(0, errors.New("")).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	profile := models.Profile{Id: uid, Login: "tester2"}
	_, st := repo.GetProfile(profile)

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_GetProfileByKeyword(t *testing.T) {
	keyword := "key"
	columns := []string{"login", "about", "avatar", "subscriptions", "subscribers"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow("tester2",
			"new user", "img.png", uint(12),
			uint(10)).ToPgxRows()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	userCheck, st := repo.GetProfileByKeyword(keyword)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, "tester2", userCheck[0].Login)
}

func TestAuthRepo_RemoveFollowingOk(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.RemoveFollowing(uid, uuid.New())

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_RemoveFollowing(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.RemoveFollowing(uid, uuid.New())

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_RemoveFollowingErr(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).RowError(0, errors.New("")).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errors.New(""))

	st := repo.RemoveFollowing(uid, uuid.New())

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdateBio(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	log.Print(uid)
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.UpdateBio(models.Profile{})

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdateBioErr(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	log.Print(uid)
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errors.New(""))

	st := repo.UpdateBio(models.Profile{Id: uid})

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdatePic(t *testing.T) {

	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.UpdateAvatar(models.Profile{})

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdatePicErr(t *testing.T) {

	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errors.New(""))

	st := repo.UpdateAvatar(models.Profile{})

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdatePass(t *testing.T) {
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.UpdatePass(models.User{})

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestAuthRepo_UpdatePassErr(t *testing.T) {

	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool, nil)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errors.New(""))

	st := repo.UpdatePass(models.User{})

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}
