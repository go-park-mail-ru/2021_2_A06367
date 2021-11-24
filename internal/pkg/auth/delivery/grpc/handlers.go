package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	grpc "github.com/go-park-mail-ru/2021_2_A06367/internal/models/grpc"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/google/uuid"
	"time"
)

//go:generate mockgen -source=auth_grpc.pb.go -destination=auth_grpc.go -package=grpc

type GrpcAuthHandler struct {
	uc auth.AuthUsecase
	AuthServiceServer
}

func NewGrpcAuthHandler(uc auth.AuthUsecase) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc: uc,
	}
}

func (h GrpcAuthHandler) Login(ctx context.Context, in *LoginUser) (*Token, error) {

	user := models.LoginUser{Login: in.Login, EncryptedPassword: in.EncryptedPassword}

	token, status := h.uc.SignIn(user)

	return &Token{Cookie: token, Status: grpc.StatusCode(status)}, nil
}

func (h GrpcAuthHandler) SignUp(ctx context.Context, in *User) (*Token, error) {
	user := models.User{
		Login:             in.Login,
		EncryptedPassword: in.EncryptedPassword,
		CreatedAt:         time.Now(),
	}
	token, status := h.uc.SignUp(user)
	return &Token{Cookie: token, Status: grpc.StatusCode(status)}, nil
}

func (h GrpcAuthHandler) GetProfile(ctx context.Context, in *UserUUID) (*Profile, error) {
	id, _ := uuid.Parse(in.ID)
	profile := models.Profile{
		Id: id,
	}
	user, status := h.uc.GetProfile(profile)
	return &Profile{
		UUID:          user.Id.String(),
		Login:         user.Login,
		Subscribers:   int64(user.Subscribers),
		Subscriptions: int64(user.Subscriptions),
		About:         user.About,
		Avatar:        user.Avatar,
		Status:        grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfilePic(ctx context.Context, in *UserUpdatePic) (*Empty, error) {
	id, _ := uuid.Parse(in.ID)
	user := models.Profile{
		Id:     id,
		Login:  in.Login,
		Avatar: in.Avatar,
	}

	status := h.uc.SetAvatar(user)
	return &Empty{
		Status: grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfilePass(ctx context.Context, in *UserUpdatePass) (*Empty, error) {
	id, _ := uuid.Parse(in.ID)
	user := models.User{
		Id:                id,
		Login:             in.Login,
		EncryptedPassword: in.Password,
	}

	status := h.uc.SetPass(user)
	return &Empty{
		Status: grpc.StatusCode(status),
	}, nil
}

func (h GrpcAuthHandler) UpdateProfileBio(ctx context.Context, in *UserUpdateBio) (*Empty, error) {
	id, _ := uuid.Parse(in.ID)
	user := models.Profile{
		Id:    id,
		Login: in.Login,
		About: in.About,
	}

	status := h.uc.SetBio(user)
	return &Empty{
		Status: grpc.StatusCode(status),
	}, nil
}
