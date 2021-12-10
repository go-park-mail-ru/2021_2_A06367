package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().SignIn(gomock.Any()).Return("", models.Okey)

	_, err = cl.Login(context.Background(), &generated.LoginUser{})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().SignUp(gomock.Any()).Return("", models.Okey)

	_, err = cl.SignUp(context.Background(), &generated.User{})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().GetProfile(gomock.Any()).Return(models.Profile{}, models.Okey)

	_, err = cl.GetProfile(context.Background(), &generated.UserUUID{ID: uuid.New().String()})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestUpdateProfilePic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().SetAvatar(gomock.Any()).Return(models.Okey)

	_, err = cl.UpdateProfilePic(context.Background(), &generated.UserUpdatePic{ID: uuid.New().String()})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestUpdateProfileBio(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().SetBio(gomock.Any()).Return(models.Okey)

	_, err = cl.UpdateProfileBio(context.Background(), &generated.UserUpdateBio{ID: uuid.New().String()})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func TestUpdateProfilePass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockAuthUsecase(ctrl)
	client := NewGrpcAuthHandler(usecase)
	srv, listener := startGRPCServer(client)
	defer srv.Stop()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cl := generated.NewAuthServiceClient(conn)

	usecase.EXPECT().SetPass(gomock.Any()).Return(models.Okey)

	_, err = cl.UpdateProfilePass(context.Background(), &generated.UserUpdatePass{ID: uuid.New().String()})
	if err != nil {
		t.Fatalf("failed due to err: %v", err)
	}

}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

func startGRPCServer(impl generated.AuthServiceServer) (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()
	generated.RegisterAuthServiceServer(srv, impl)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return srv, listener
}
