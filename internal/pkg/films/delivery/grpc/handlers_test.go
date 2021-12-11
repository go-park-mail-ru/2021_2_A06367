package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestFilmByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetCompilation("").Times(1).Return([]models.Film{
		models.Film{
			Id:           uuid.UUID{},
			Title:        "",
			Genres:       nil,
			Country:      "",
			ReleaseRus:   time.Time{},
			Year:         0,
			Director:     nil,
			Authors:      nil,
			Actors:       nil,
			Release:      time.Time{},
			Duration:     0,
			Language:     "",
			Budget:       "",
			Age:          0,
			Pic:          nil,
			Src:          nil,
			Description:  "",
			IsSeries:     true,
			Seasons:      nil,
			Rating:       0,
			NeedsPayment: true,
			IsAvailable:  false,
			Slug:         "",
		},
	}, models.Okey)

	genre, err := cl.FilmByGenre(context.Background(), &generated.KeyWord{})
	log.Println(genre)
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmBySelection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetSelection(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmBySelection(context.Background(), &generated.KeyWord{})
	log.Println(genre)
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmsByActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetFilmsOfActor(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmsByActor(context.Background(), &generated.UUID{Id: uuid.New().String()})
	log.Println(genre)
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetFilm(gomock.Any()).Times(1).Return(models.Film{}, models.Okey)

	genre, err := cl.FilmById(context.Background(), &generated.UUID{Id: uuid.New().String()})
	log.Println(genre)
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetCompilationForUser(gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	genre, err := cl.FilmsByUser(context.Background(), &generated.UUID{Id: uuid.NewString()})
	log.Println(genre)
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmStartSelection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().GetStartSelections(true, gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)
	usecase.EXPECT().GetStartSelections(false, gomock.Any()).Times(1).Return([]models.Film{}, models.Okey)

	_, err = cl.FilmStartSelection(context.Background(), &generated.UUID{})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}
	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}
	_, err = cl.FilmStartSelection(context.Background(), &generated.UUID{Id: id.String()})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmStarred(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().AddStarred(gomock.Any(), gomock.Any()).Times(1).Return(models.Okey)

	_, err = cl.AddStarred(context.Background(), &generated.Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmRemoveStarred(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().RemoveStarred(gomock.Any(), gomock.Any()).Times(1).Return(models.Okey)

	_, err = cl.RemoveStarred(context.Background(), &generated.Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmWl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().AddWatchlist(gomock.Any(), gomock.Any()).Times(1).Return(models.Okey)

	_, err = cl.AddWatchList(context.Background(), &generated.Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestFilmRemoveWl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockFilmsUsecase(ctrl)
	client := NewGrpcFilmsHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewFilmsServiceClient(conn)

	usecase.EXPECT().RemoveWatchlist(gomock.Any(), gomock.Any()).Times(1).Return(models.Okey)

	_, err = cl.RemoveWatchList(context.Background(), &generated.Pair{
		FilmUUID: uuid.UUID{}.String(),
		UserUUID: uuid.UUID{}.String(),
	})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mocks.NewMockFilmsUsecase(ctrl)
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

func startGRPCServer(impl generated.FilmsServiceServer) (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()
	generated.RegisterFilmsServiceServer(srv, impl)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return srv, listener
}
