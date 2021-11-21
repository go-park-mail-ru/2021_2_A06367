package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"google.golang.org/grpc"
)

type GrpcAuthHandler struct {
	uc auth.AuthUsecase
}

func NewGrpcAuthHandler(uc auth.AuthUsecase) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc: uc,
	}
}

func (h GrpcAuthHandler) Login(ctx context.Context, in *LoginUser, opts ...grpc.CallOption) (*Token, error) {
	return &Token{}, nil
}

func (h GrpcAuthHandler) SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	return &Token{}, nil
}

func (h GrpcAuthHandler) GetProfile(ctx context.Context, in *UserUUID, opts ...grpc.CallOption) (*User, error) {
	return &User{}, nil
}

func (h GrpcAuthHandler) UpdateProfilePic(ctx context.Context, in *UserUpdatePic, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}

func (h GrpcAuthHandler) UpdateProfilePass(ctx context.Context, in *UserUpdatePass, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}

func (h GrpcAuthHandler) UpdateProfileBio(ctx context.Context, in *UserUpdateBio, opts ...grpc.CallOption) (*Nothing, error) {
	return &Nothing{}, nil
}
