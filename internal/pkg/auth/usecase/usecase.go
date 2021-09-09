package usecase

import "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}
