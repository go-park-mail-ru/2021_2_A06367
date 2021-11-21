package main

import (
	"context"
	grpc3 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc"
	authRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/config"
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
	defer logger.Sync()
	if err != nil {
		return err
	}
	defer logger.Sync()
	zapSugar := logger.Sugar()

	encrypter := authUsecase.NewEncrypter()
	tokenGenerator := authUsecase.NewTokenator()
	authRepo := authRepository.NewAuthRepo(pool, zapSugar)
	authUse := authUsecase.NewAuthUsecase(authRepo, tokenGenerator, encrypter, zapSugar)
	service := grpc3.NewGrpcAuthHandler(authUse)

	srv, ok := net.Listen("tcp", ":8020")
	if ok != nil {
		log.Fatalln("can't listen port", err)
	}

	server := grpc.NewServer()

	grpc3.RegisterAuthServiceServer(server, service)

	log.Print("main running on: ", srv.Addr())
	return server.Serve(srv)
}
