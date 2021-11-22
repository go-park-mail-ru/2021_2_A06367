package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models/grpc"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/google/uuid"
	"log"
)

type GrpcFilmsHandler struct {
	uc films.FilmsUsecase
	FilmsServiceServer
}

func NewGrpcFilmsHandler(uc films.FilmsUsecase) *GrpcFilmsHandler {
	return &GrpcFilmsHandler{
		uc: uc,
	}
}

func (g *GrpcFilmsHandler) FilmByGenre(ctx context.Context, in *KeyWord) (*Films, error) {
	filmSet, status := g.uc.GetCompilation(in.Word)
	data := g.FilmsAdaptor(filmSet)

	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) FilmBySelection(ctx context.Context, in *KeyWord) (*Films, error) {
	filmSet, status := g.uc.GetSelection(in.Word)
	data := g.FilmsAdaptor(filmSet)

	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) FilmsByActor(ctx context.Context, in *UUID) (*Films, error) {
	id, _ := uuid.Parse(in.Id)

	actor := models.Actors{Id: id}
	filmSet, status := g.uc.GetFilmsOfActor(actor)
	data := g.FilmsAdaptor(filmSet)

	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) FilmById(ctx context.Context, in *UUID) (*Film, error) {
	id, _ := uuid.Parse(in.Id)

	filmReq := models.Film{Id: id}
	film, status := g.uc.GetFilm(filmReq)
	data := g.FilmAdaptor(film)

	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) FilmsByUser(ctx context.Context, in *UUID) (*Films, error) {
	id, _ := uuid.Parse(in.Id)
	user := models.User{Id: id}
	filmSet, status := g.uc.GetCompilationForUser(user)
	data := g.FilmsAdaptor(filmSet)

	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) FilmStartSelection(ctx context.Context, in *UUID) (*Films, error) {
	var (
		films  []models.Film
		status models.StatusCode
	)

	id, err := uuid.Parse(in.Id)
	if err != nil {
		user := models.User{}
		films, status = g.uc.GetStartSelections(false, user)
	} else {
		user := models.User{Id: id}
		films, status = g.uc.GetStartSelections(true, user)
	}
	data := g.FilmsAdaptor(films)
	data.Status = grpc.StatusCode(status)

	return data, nil
}
func (g *GrpcFilmsHandler) AddStarred(ctx context.Context, in *Pair) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.AddStarred(film, user)
	return &Nothing{
		Status: grpc.StatusCode(status),
	}, nil
}
func (g *GrpcFilmsHandler) RemoveStarred(ctx context.Context, in *Pair) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.RemoveStarred(film, user)
	log.Print(status)
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) AddWatchList(ctx context.Context, in *Pair) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.AddWatchlist(film, user)
	return &Nothing{
		Status: grpc.StatusCode(status),
	}, nil
}
func (g *GrpcFilmsHandler) RemoveWatchList(ctx context.Context, in *Pair) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.RemoveWatchlist(film, user)
	return &Nothing{
		Status: grpc.StatusCode(status),
	}, nil
}
func (g *GrpcFilmsHandler) Starred(ctx context.Context, in *UUID) (*Films, error) {

	userId, _ := uuid.Parse(in.Id)
	user := models.User{Id: userId}
	films, status := g.uc.GetStarred(user)
	if status == models.Okey {
		return g.FilmsAdaptor(films), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) WatchList(ctx context.Context, in *UUID) (*Films, error) {
	userId, _ := uuid.Parse(in.Id)
	user := models.User{Id: userId}
	films, status := g.uc.GetWatchlist(user)
	if status == models.Okey {
		return g.FilmsAdaptor(films), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) Random(ctx context.Context, in *Nothing) (*Film, error) {
	film, status := g.uc.Randomize()
	data := g.FilmAdaptor(film)
	data.Status = grpc.StatusCode(status)
	return data, nil
}

func (g *GrpcFilmsHandler) FilmAdaptor(film models.Film) *Film {

	var actors []string
	for i := 0; i < len(film.Actors); i++ {
		actors = append(actors, film.Actors[i].String())
	}
	var gfilm *Film
	gfilm = &Film{
		Id:                 film.Id.String(),
		Title:              film.Title,
		Genres:             film.Genres,
		Country:            film.Country,
		Year:               int64(film.Year),
		ReleaseRus:         film.ReleaseRus.String(),
		Director:           film.Director,
		Authors:            film.Authors,
		Actors:             actors,
		Release:            film.Release.String(),
		Duration:           int64(film.Duration),
		ReleaseRusLanguage: film.ReleaseRus.String(),
		Budget:             film.Budget,
		Age:                int64(film.Age),
		Pic:                film.Pic,
		Src:                film.Src,
		Description:        film.Description,
		IsSeries:           film.IsSeries,
		Seasons:            nil,
	}

	if film.Seasons != nil {
		var gs []*Season
		for i := 0; i < len(*film.Seasons); i++ {
			gs = append(gs, &Season{
				Num:  int64((*film.Seasons)[i].Num),
				Src:  (*film.Seasons)[i].Src,
				Pics: (*film.Seasons)[i].Pics,
			})
		}
		gfilm.Seasons = gs
	}
	return gfilm
}

func (g *GrpcFilmsHandler) FilmsAdaptor(films []models.Film) *Films {

	var result Films
	for i := 0; i < len(films); i++ {
		film := g.FilmAdaptor(films[i])
		result.Data = append(result.Data, film)
	}
	return &result
}
