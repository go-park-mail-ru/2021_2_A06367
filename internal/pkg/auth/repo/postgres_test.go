package repo

import (
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const pswd = "test"

func prepare(t *testing.T, pgxRows pgx.Rows) *AuthRepo { // mock on repo and uid used as uuid for everything that has any uid filed
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	return repo
}

func TestAuthRepo_CheckUser(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
	ecnryptedPswd := usecase.NewEncrypter().EncryptPswd(pswd)
	columns := []string{"encrypted_password"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(ecnryptedPswd).ToPgxRows()
	pgxRows.Next()
	repo := prepare(t, pgxRows)
	user := models.User{Id: uid, Login: "tester2", EncryptedPassword: ecnryptedPswd}
	userCheck, st := repo.CheckUser(user)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, user.EncryptedPassword, userCheck.EncryptedPassword)
}

func TestAuthRepo_CreateUser(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
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

func TestAuthRepo_AddFollowing(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
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

func TestAuthRepo_GetProfile(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
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

	repo := NewAuthRepo(mockPool)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	userCheck, st := repo.GetProfileByKeyword(keyword)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	assert.Equal(t, "tester2", userCheck[0].Login)
}

func TestAuthRepo_RemoveFollowing(t *testing.T) {
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
	columns := []string{"id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(1).ToPgxRows()
	pgxRows.Next()
	os.Setenv("SECRET", "test")
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewAuthRepo(mockPool)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	st := repo.RemoveFollowing(uid, uuid.New())

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}
