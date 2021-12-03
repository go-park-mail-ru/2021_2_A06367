package http

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	subs "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type FilmsHandler struct {
	logger *zap.SugaredLogger
	client generated.FilmsServiceClient
	subsClient subs.SubsServiceClient
}

func NewFilmsHandler(logger *zap.SugaredLogger, cl generated.FilmsServiceClient) *FilmsHandler {
	return &FilmsHandler{
		logger: logger,
		client: cl,
	}
}

// FilmByGenre godoc
// @Summary Get films of genre
// @Description Get films of genre
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} []models.Film
// @Param genre path string true "Боевик"
// @Failure 400,404 {string} 1
// @Router /films/genre/{genre} [get]
func (h FilmsHandler) FilmByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genres, found := vars["genre"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	films, err := h.client.FilmByGenre(context.Background(), &generated.KeyWord{Word: genres})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(films)

	util.Response(w, models.StatusCode(films.Status), filmSet)
}

// FilmBySelection godoc
// @Summary Get details of films
// @Description Get details of films
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} []models.Film
// @Param selection path string true "КАКАЯ-ТО СТРОКА"
// @Failure 400,404 {string} 1
// @Router /films/selection/{selection} [get]
// @Router /films/selection/{selection} [options]
func (h FilmsHandler) FilmBySelection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	selection, found := vars["selection"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	films, err := h.client.FilmBySelection(context.Background(), &generated.KeyWord{Word: selection})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(films)
	util.Response(w, models.StatusCode(films.Status), filmSet)
}

// FilmByActor godoc
// @Summary Get details of actor from selection
// @Description Get details of actor from selection
// @Tags Film
// @Accept json
// @Produce json
// @Param actor_id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {object} []models.Film
// @Failure 400,404 {string} 1
// @Router /films/selection/actor/{actor_id} [get]
func (h FilmsHandler) FilmByActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, found := vars["actor_id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	idActor, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	films, err := h.client.FilmsByActor(context.Background(), &generated.UUID{Id: idActor.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(films)
	util.Response(w, models.StatusCode(films.Status), filmSet)
}

// FilmById godoc
// @Summary Get details of film
// @Description Get details of film
// @Tags Film
// @Accept json
// @Produce json
// @Param film_id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {object} models.Film
// @Failure 400,404 {string} 1
// @Router /films/film/{film_id} [get]
func (h FilmsHandler) FilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, found := vars["film_id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film, err := h.client.FilmById(context.Background(), &generated.UUID{Id: id.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}
	filmSet := h.FilmToModel(film)

	if filmSet.NeedsPayment {
		//если токена нет и пользователь неавторизован
		access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
		if err != nil || access == nil {
			filmSet.IsAvailable = false
		} else {
			//если токен есть
			license, err := h.subsClient.GetLicense(context.Background(), &subs.LicenseUUID{ID: access.Id.String()})
			//если микросервис отвалился
			if err != nil {
				filmSet.IsAvailable = false
			} else {
				//если микросервис ок и надо просто проверить лицензию
				parsed, _ := time.Parse(time.RFC3339, license.ExpiresDate)
				filmSet.IsAvailable = time.Now().Before(parsed)
			}
		}
	}

	util.Response(w, models.StatusCode(film.Status), filmSet)

}

// FilmsByUser godoc
// @Summary Get details of personal film
// @Description Get details of personal film
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} models.Film
// @Failure 400,404 {string} 1
// @Router /films/selection/user/personal [get]
func (h FilmsHandler) FilmsByUser(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	films, err := h.client.FilmsByUser(context.Background(), &generated.UUID{Id: access.Id.String()})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	filmSet := h.FilmsToModels(films)
	util.Response(w, models.StatusCode(films.Status), filmSet)
}

func (h FilmsHandler) FilmStartSelection(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	if access == nil {
		access = &models.AccessDetails{}
	}
	selection, err := h.client.FilmStartSelection(context.Background(), &generated.UUID{Id: access.Id.String()})
	if err != nil {
		return
	}
	film := h.FilmsToModels(selection)
	util.Response(w, models.StatusCode(selection.Status), film)
}

func (h FilmsHandler) AddStarred(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.AddStarred(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) IfStarred(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.IfStarred(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) IfWl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.IfWatchList(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) RemoveStarred(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.RemoveStarred(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) AddWatchlist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.AddWatchList(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) RemoveWatchlist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	none, err := h.client.RemoveWatchList(context.Background(), &generated.Pair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}

	util.Response(w, models.StatusCode(none.Status), nil)
}

func (h FilmsHandler) GetStarred(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	films, err := h.client.Starred(context.Background(), &generated.UUID{
		Id: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}
	filmSet := h.FilmsToModels(films)
	util.Response(w, models.StatusCode(films.Status), filmSet)
}

func (h FilmsHandler) GetWatchlist(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}
	films, err := h.client.WatchList(context.Background(), &generated.UUID{
		Id: user.Id.String(),
	})
	if err != nil {
		util.Response(w, models.NotFound, nil)
		return
	}
	filmSet := h.FilmsToModels(films)
	util.Response(w, models.StatusCode(films.Status), filmSet)
}

func (h FilmsHandler) RandomFilm(w http.ResponseWriter, r *http.Request) {

	film, err := h.client.Random(context.Background(), &generated.Nothing{})
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	util.Response(w, models.StatusCode(film.Status), film)

}

func (h FilmsHandler) SetRating(w http.ResponseWriter, r *http.Request) {

	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	user := models.User{Id: access.Id}

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	rating := r.URL.Query().Get("rating")
	mark, err := strconv.ParseFloat(rating, 32)

	res, err := h.client.SetRating(context.Background(), &generated.RatingPair{
		FilmUUID: film.Id.String(),
		UserUUID: user.Id.String(),
		Rating:   float32(mark),
	})
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	util.Response(w, models.StatusCode(res.Status), film)

}

func (h FilmsHandler) GetRating(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, found := vars["id"]
	if !found {
		util.Response(w, models.NotFound, nil)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	film := models.Film{Id: id}

	res, err := h.client.GetRating(context.Background(), &generated.UUID{
		Id: film.Id.String(),
	})

	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}
	film.Rating = float64(res.Rating)
	util.Response(w, models.StatusCode(res.Status), film)

}

func (h FilmsHandler) FilmToModel(film *generated.Film) models.Film {
	layout := "2006-01-02"
	id, err := uuid.Parse(film.Id)
	if err != nil {
		return models.Film{}
	}
	dateRus := film.ReleaseRus[0:len(layout)]
	releaseru, err := time.Parse(layout, dateRus)
	if err != nil {
		return models.Film{}
	}
	dateRelease := film.Release[0:len(layout)]
	release, err := time.Parse(layout, dateRelease)
	if err != nil {
		return models.Film{}
	}

	var actors []uuid.UUID
	for i := 0; i < len(film.Actors); i++ {
		id2, err2 := uuid.Parse(film.Actors[i])
		if err2 != nil {
			return models.Film{}
		}
		actors = append(actors, id2)
	}

	SeasonsOut := []models.Season{}
	for _, season := range film.Seasons {
		temp := models.Season{
			Num:  int(season.Num),
			Src:  season.Src,
			Pics: season.Pics,
		}
		SeasonsOut = append(SeasonsOut, temp)
	}
	if len(SeasonsOut) == 0 {
		temp := models.Season{}
		SeasonsOut = append(SeasonsOut, temp)
	}

	return models.Film{
		Id:           id,
		Title:        film.Title,
		Genres:       film.Genres,
		Country:      film.Country,
		ReleaseRus:   releaseru,
		Year:         int(film.Year),
		Director:     film.Director,
		Authors:      film.Authors,
		Actors:       actors,
		Release:      release,
		Duration:     int(film.Duration),
		Budget:       film.Budget,
		Age:          int(film.Age),
		Pic:          film.Pic,
		Src:          film.Src,
		Description:  film.Description,
		IsSeries:     film.IsSeries,
		Seasons:      &SeasonsOut,
		Rating:       float64(film.Rating),
		NeedsPayment: film.NeedsPayment,
		Slug:         film.Slug,
	}
}

func (h FilmsHandler) FilmsToModels(films *generated.Films) []models.Film {
	var result []models.Film
	for i := 0; i < len(films.Data); i++ {
		film := h.FilmToModel(films.Data[i])
		result = append(result, film)
	}
	return result
}
