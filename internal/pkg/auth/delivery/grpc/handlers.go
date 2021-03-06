package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	grpc "github.com/go-park-mail-ru/2021_2_A06367/internal/models/grpc"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	"github.com/google/uuid"
	"log"
	"time"
)

//go:generate mockgen -source=auth_grpc.pb.go -destination=auth_grpc.go -package=grpc

type GrpcAuthHandler struct {
	uc auth.AuthUsecase
	generated.AuthServiceServer
}

func NewGrpcAuthHandler(uc auth.AuthUsecase) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc: uc,
	}
}

func (h GrpcAuthHandler) Login(ctx context.Context, in *generated.LoginUser) (*generated.Token, error) {

	user := models.LoginUser{Login: in.Login, EncryptedPassword: in.EncryptedPassword}

	token, status := h.uc.SignIn(user)

	return &generated.Token{Cookie: token, Status: grpc.StatusCode(status)}, nil
}

func (h GrpcAuthHandler) SignUp(ctx context.Context, in *generated.User) (*generated.Token, error) {
	user := models.User{
		Login:             in.Login,
		EncryptedPassword: in.EncryptedPassword,
		CreatedAt:         time.Now(),
	}
	token, status := h.uc.SignUp(user)
	return &generated.Token{Cookie: token, Status: grpc.StatusCode(status)}, nil
}

func (h GrpcAuthHandler) GetProfile(ctx context.Context, in *generated.UserUUID) (*generated.Profile, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.Profile{
			Status: grpc.StatusCode(models.InternalError),
		}, nil
	}
	profile := models.Profile{
		Id: id,
	}
	user, status := h.uc.GetProfile(profile)
	return &generated.Profile{
		UUID:          user.Id.String(),
		Login:         user.Login,
		Subscribers:   int64(user.Subscribers),
		Subscriptions: int64(user.Subscriptions),
		About:         user.About,
		Avatar:        user.Avatar,
		Status:        grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfilePic(ctx context.Context, in *generated.UserUpdatePic) (*generated.Empty, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.Empty{
			Status: grpc.StatusCode(models.InternalError),
		}, nil
	}
	user := models.Profile{
		Id:     id,
		Login:  in.Login,
		Avatar: in.Avatar,
	}

	status := h.uc.SetAvatar(user)
	return &generated.Empty{
		Status: grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfilePass(ctx context.Context, in *generated.UserUpdatePass) (*generated.Empty, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.Empty{
			Status: grpc.StatusCode(models.InternalError),
		}, nil
	}
	user := models.User{
		Id:                id,
		Login:             in.Login,
		EncryptedPassword: in.Password,
	}

	status := h.uc.SetPass(user)
	return &generated.Empty{
		Status: grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfileBio(ctx context.Context, in *generated.UserUpdateBio) (*generated.Empty, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.Empty{
			Status: grpc.StatusCode(models.InternalError),
		}, nil
	}
	user := models.Profile{
		Id:    id,
		Login: in.Login,
		About: in.About,
	}

	status := h.uc.SetBio(user)
	return &generated.Empty{
		Status: grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) CheckByLogin(ctx context.Context, in *generated.LoginUser) (*generated.UserUUID, error) {

	log.Println(in)
	user := models.User{Login: in.Login}

	token, err := h.uc.CheckUserLogin(user)
	log.Println(token, err)
	return &generated.UserUUID{ID: token.Id.String()}, nil
}
