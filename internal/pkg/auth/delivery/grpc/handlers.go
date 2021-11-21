package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
)

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
	return &Token{}, nil
}

func (h GrpcAuthHandler) SignUp(ctx context.Context, in *User) (*Token, error) {
	return &Token{}, nil
}

func (h GrpcAuthHandler) GetProfile(ctx context.Context, in *UserUUID) (*User, error) {
	return &User{}, nil
}

func (h GrpcAuthHandler) UpdateProfilePic(ctx context.Context, in *UserUpdatePic) (*Empty, error) {
	return &Empty{}, nil
}

func (h GrpcAuthHandler) UpdateProfilePass(ctx context.Context, in *UserUpdatePass) (*Empty, error) {
	return &Empty{}, nil
}

func (h GrpcAuthHandler) UpdateProfileBio(ctx context.Context, in *UserUpdateBio) (*Empty, error) {
	return &Empty{}, nil
}
