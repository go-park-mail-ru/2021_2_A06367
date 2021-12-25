package main

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/config"
	grpc3 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc"
	generated "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/repo"
	usecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
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

	r := repo.NewSubsRepo(pool)
	uc := usecase.NewSubsUsecase(r)
	service := grpc3.NewGrpcSubsHandler(uc)

	srv, ok := net.Listen("tcp", ":8030")
	if ok != nil {
		log.Fatalln("can't listen port")
	}

	server := grpc.NewServer()

	generated.RegisterSubsServiceServer(server, service)

	log.Print("subs running on: ", srv.Addr())
	return server.Serve(srv)
}
