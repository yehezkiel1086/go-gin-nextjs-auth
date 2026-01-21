package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/util"
)

type AuthService struct {
	conf *config.JWT
	userRepo port.UserRepository
}

func NewAuthService(conf *config.JWT, userRepo port.UserRepository) *AuthService {
	return &AuthService{
		conf,
		userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if err := util.CompareHashedPwd(user.Password, password); err != nil {
		return "", "", err
	}

	// generate jwt tokens
	refreshToken, err := util.GenerateRefreshToken(as.conf, user)
	if err != nil {
		return "", "", err
	}

	accessToken, err := util.GenerateAccessToken(as.conf, user)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (as *AuthService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return as.userRepo.GetUserByEmail(ctx, email)
}
