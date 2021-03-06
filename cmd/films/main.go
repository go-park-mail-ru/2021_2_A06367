package main

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/config"
	grpc2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	filmsRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/repo"
	filmsUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {

	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)

	zapSugar := logger.Sugar()

	filmsRepo := filmsRepository.NewFilmsRepo(pool, zapSugar)
	filmsUse := filmsUsecase.NewFilmsUsecase(filmsRepo, zapSugar)
	service := grpc2.NewGrpcFilmsHandler(filmsUse)

	srv, ok := net.Listen("tcp", ":8010")
	if ok != nil {
		log.Fatalln("can't listen port", err)
	}

	server := grpc.NewServer()

	generated.RegisterFilmsServiceServer(server, service)

	log.Print("films running on: ", srv.Addr())
	return server.Serve(srv)
}
