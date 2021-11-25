package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestFilmByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetCompilation("").Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmByGenre(context.Background(), &KeyWord{})
	log.Println(genre)

}

func TestFilmBySelection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetSelection(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmBySelection(context.Background(), &KeyWord{})
	log.Println(genre)

}

func TestFilmsByActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetFilmsOfActor(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmsByActor(context.Background(), &UUID{})
	log.Println(genre)

}

func TestFilmById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetFilm(gomock.Any()).Times(1).Return(models.Film{}, models.Okey)

	genre, err := cl.FilmById(context.Background(), &UUID{})
	log.Println(genre)

}

func TestFilmsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetCompilationForUser(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmsByUser(context.Background(), &UUID{})
	log.Println(genre)

}

func TestFilmStartSelection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().GetStartSelections(true, gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)
	usecase.EXPECT().GetStartSelections(false, gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmStartSelection(context.Background(), &UUID{})
	id, _ := uuid.NewUUID()
	genre, err = cl.FilmStartSelection(context.Background(), &UUID{Id: id.String()})
	log.Println(genre)

}

func TestFilmStarred(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().AddStarred(gomock.Any(), gomock.Any()).Times(1).Return( models.Okey)

	_, err = cl.AddStarred(context.Background(), &Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})

}

func TestFilmRemoveStarred(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().RemoveStarred(gomock.Any(), gomock.Any()).Times(1).Return( models.Okey)

	_, err = cl.RemoveStarred(context.Background(), &Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})

}

func TestFilmWl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().AddWatchlist(gomock.Any(), gomock.Any()).Times(1).Return( models.Okey)

	_, err = cl.AddWatchList(context.Background(), &Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})

}

func TestFilmRemoveWl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := NewFilmsServiceClient(conn)

	usecase.EXPECT().RemoveWatchlist(gomock.Any(), gomock.Any()).Times(1).Return( models.Okey)

	_, err = cl.RemoveWatchList(context.Background(), &Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})

}

func TestConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := films.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)

	//fl := Films{Data: []*Film{}}
	fln := []models.Film{}

	client.FilmsAdaptor(fln)

}
func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

func startGRPCServer(impl FilmsServiceServer) (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()
	RegisterFilmsServiceServer(srv, impl)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return srv, listener
}
