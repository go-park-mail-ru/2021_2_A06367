package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"google.golang.org/grpc"
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
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmBySelection(ctx context.Context, in *KeyWord, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmsByActor(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmById(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Film, error) {
	return &Film{}, nil
}
func (g *GrpcFilmsHandler) FilmsByUser(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) FilmStartSelection(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) AddStarred(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) RemoveStarred(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) AddWatchList(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) RemoveWatchList(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}
func (g *GrpcFilmsHandler) Starred(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) WatchList(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Films, error) {
	return &Films{}, nil
}
func (g *GrpcFilmsHandler) Random(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Film, error) {
	return &Film{}, nil
}
