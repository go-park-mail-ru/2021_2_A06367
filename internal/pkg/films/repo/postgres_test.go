package repo

import (
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func prepare(t *testing.T) (*FilmsRepo, uuid.UUID) { // mock on repo and uid used as uuid for everything that has any uid filed
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "title", "year", "director", "authors", "actors", "release", "duration", "language", "pic", "src"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "Policeman from Rublevka",
			2017, []string{"Alex"}, []string{"Alex"},
			[]uuid.UUID{uid}, time.Now(), 120, "russian", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	return filmRepo, uid
}

func TestFilmsRepo_GetFilmById(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)
	filmRequest := models.Film{
		Id: uid,
	}
	film, status := filmRepo.GetFilmById(filmRequest)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film.Id)
}

func TestFilmsRepo_GetFilmsByActor(t *testing.T) {
	filmRepo, uid := prepare(t)
	filmRequest := models.Actors{
		Id: uid,
	}
	_, status := filmRepo.GetFilmsByActor(filmRequest)
	if status == models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, filmRequest.Id)
	//assert.Equal(t, uid, film[0].Id)
	//assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetFilmsByKeyword(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetFilmsByKeyword("filmRequest")
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsByTopicOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetFilmsByTopic("filmRequest")
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsHottestOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetHottestFilms()
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsNewestOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetNewestFilms()
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsByKeywordOk(t *testing.T) {
	keyword := "SATANA"
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", false, false, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", false, false, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow("4").AddRow("4").ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	_, status := filmRepo.GetFilmsByKeyword(keyword)
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsByTopic(t *testing.T) {
	topic := "SATANA"
	filmRepo, _ := prepare(t)

	_, status := filmRepo.GetFilmsByTopic(topic)
	if status == models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, filmRepo)
	//assert.Equal(t, uid, film[0].Id)
	//assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetFilmsByActors(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetFilmsByActor(models.Actors{Id: uid})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsByUserOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetFilmsByUser(models.User{Id: uid})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetFilmsByUser(t *testing.T) {
	filmRepo, uid := prepare(t)
	user := models.User{Id: uid}
	_, status := filmRepo.GetFilmsByUser(user)
	if status == models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, filmRepo)
	//assert.Equal(t, uid, film[0].Id)
	//assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetHottestFilms(t *testing.T) {
	filmRepo, _ := prepare(t)
	_, status := filmRepo.GetHottestFilms()
	if status == models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, filmRepo)
	//assert.Equal(t, uid, film[0].Id)
	//assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetNewestFilms(t *testing.T) {
	filmRepo, _ := prepare(t)
	_, status := filmRepo.GetNewestFilms()
	if status == models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, filmRepo)
	//assert.Equal(t, uid, film[0].Id)
	//assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_InsertStarred(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	filmRepo := NewFilmsRepo(mockPool, nil)

	status := filmRepo.InsertStarred(models.Film{}, models.User{})
	if status != models.Conflict {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_DeleteStarred(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	filmRepo := NewFilmsRepo(mockPool, nil)

	status := filmRepo.DeleteStarred(models.Film{}, models.User{})
	if status != models.Conflict {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_InsertWL(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	filmRepo := NewFilmsRepo(mockPool, nil)

	status := filmRepo.InsertWatchlist(models.Film{}, models.User{})
	if status != models.Conflict {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_DeleteWL(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	filmRepo := NewFilmsRepo(mockPool, nil)

	status := filmRepo.DeleteWatchlist(models.Film{}, models.User{})
	if status != models.Conflict {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetStarredOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	_, status := filmRepo.GetStarredFilms(models.User{})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_GetWLOk(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	_, status := filmRepo.GetWatchlistFilms(models.User{})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_SetRating(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().Exec(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

	filmRepo := NewFilmsRepo(mockPool, nil)

	status := filmRepo.SetRating(models.Film{}, models.User{}, 0)
	if status != models.Conflict {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_IfStarred(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	status := filmRepo.IfStarred(models.Film{}, models.User{})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}
func TestFilmsRepo_IfWl(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	status := filmRepo.IfWatchList(models.Film{}, models.User{})
	if status != models.Okey {
		t.Error("Wrong result")
	}
}

func TestFilmsRepo_Random(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		t.Error(err)
	}
	filmRepo := NewFilmsRepo(mockPool, nil)
	columns := []string{"id", "genres", "country",
		"releaseRus", "title", "year",
		"director", "authors", "actors",
		"release", "duration", "language",
		"budget", "age", "pic",
		"src", "description", "isSeries",
		"needsPayment", "slug"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		AddRow(uid, []string{"comedy"}, "abc",
			time.Now(), "Policeman from Rublevka", 2017,
			[]string{"Alex"}, []string{"Alex"}, []uuid.UUID{uid},
			time.Now(), int(time.Hour), "200",
			"120", 120, []string{"img.png"}, []string{"img.png"}, "avc", true, true, "avc").
		ToPgxRows()
	pgxRows.Next()

	pgxRowsr2 := pgxpoolmock.NewRows([]string{"id", "pic", "src"}).
		AddRow("", []string{"img.png"}, []string{"img.png"}).ToPgxRows()
	pgxRowsr2.Next()

	pgxRowsr := pgxpoolmock.NewRows([]string{"rating"}).
		AddRow(4).ToPgxRows()
	pgxRowsr.Next()

	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr2, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsr)

	_, status := filmRepo.GetRandom()
	if status != models.Okey {
		t.Error("Wrong result")
	}
}
