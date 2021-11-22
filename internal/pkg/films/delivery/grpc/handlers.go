package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

type GrpcFilmsHandler struct {
	uc films.FilmsUsecase
}

func NewGrpcFilmsHandler(uc films.FilmsUsecase) *GrpcFilmsHandler {
	return &GrpcFilmsHandler{
		uc: uc,
	}
}

func (g *GrpcFilmsHandler) FilmByGenre(ctx context.Context, in *KeyWord, opts ...grpc.CallOption) (*Films, error) {
	filmSet, status := g.uc.GetCompilation(in.Word)
	if status == models.Okey {
		return g.FilmsAdaptor(filmSet), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmBySelection(ctx context.Context, in *KeyWord, opts ...grpc.CallOption) (*Films, error) {
	filmSet, status := g.uc.GetSelection(in.Word)
	if status == models.Okey {
		return g.FilmsAdaptor(filmSet), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmsByActor(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	id, _ := uuid.Parse(in.Id)

	actor := models.Actors{Id: id}
	filmSet, status := g.uc.GetFilmsOfActor(actor)
	if status == models.Okey {
		return g.FilmsAdaptor(filmSet), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmById(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Film, error) {
	id, _ := uuid.Parse(in.Id)

	filmReq := models.Film{Id: id}
	film, status := g.uc.GetFilm(filmReq)
	if status == models.Okey {
		return g.FilmAdaptor(film), nil
	}
	return &Film{}, nil
}
func (g *GrpcFilmsHandler) FilmsByUser(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	id, _ := uuid.Parse(in.Id)
	user := models.User{Id: id}
	filmSet, status := g.uc.GetCompilationForUser(user)
	if status == models.Okey {
		return g.FilmsAdaptor(filmSet), nil
	}

	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmStartSelection(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		user := models.User{}
		filmSet, status := g.uc.GetStartSelections(false, user)
		if status == models.Okey {
			return g.FilmsAdaptor(filmSet), nil
		}

		return &Films{}, nil
	}
	user := models.User{Id: id}
	filmSet, status := g.uc.GetStartSelections(true, user)
	if status == models.Okey {
		return g.FilmsAdaptor(filmSet), nil
	}

	return &Films{}, nil
}
func (g *GrpcFilmsHandler) AddStarred(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.AddStarred(film, user)
	log.Print(status)
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) RemoveStarred(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.RemoveStarred(film, user)
	log.Print(status)
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) AddWatchList(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.AddWatchlist(film, user)
	log.Print(status)
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) RemoveWatchList(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	filmId, _ := uuid.Parse(in.FilmUUID)
	film := models.Film{Id: filmId}
	userId, _ := uuid.Parse(in.UserUUID)
	user := models.User{Id: userId}
	status := g.uc.RemoveWatchlist(film, user)
	log.Print(status)
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) Starred(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {

	userId, _ := uuid.Parse(in.Id)
	user := models.User{Id: userId}
	films, status := g.uc.GetStarred(user)
	if status == models.Okey {
		return g.FilmsAdaptor(films), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) WatchList(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	userId, _ := uuid.Parse(in.Id)
	user := models.User{Id: userId}
	films, status := g.uc.GetWatchlist(user)
	if status == models.Okey {
		return g.FilmsAdaptor(films), nil
	}
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) Random(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Film, error) {
	film, status := g.uc.Randomize()
	if status == models.Okey {
		return g.FilmAdaptor(film), nil
	}
	return &Film{}, nil
}

func (g *GrpcFilmsHandler) FilmAdaptor(film models.Film) *Film {

	var actors []string
	for i := 0; i < len(film.Actors); i++ {
		actors = append(actors, film.Actors[i].String())
	}
	return &Film{
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
}

func (g *GrpcFilmsHandler) FilmsAdaptor(films []models.Film) *Films {

	var result Films
	for i := 0; i < len(films); i++ {
		film := g.FilmAdaptor(films[i])
		result.Data = append(result.Data, film)
	}
	return &result
}
