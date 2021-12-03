package main

import (
	grpc3 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc"
	generated "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	usecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/usecase"
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

	uc := usecase.NewSubsUsecase()
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
