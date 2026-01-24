package service

import (
	"context"
	"errors"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
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

const (
	RefreshType = "refresh"
	AccessType = "access"
)

func (as *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	// check email and password
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if !user.IsVerified {
		return "", "", errors.New("user is not verified")
	}

	if err := util.CompareHashedPwd(user.Password, password); err != nil {
		return "", "", err
	}

	// generate jwt tokens
	refreshToken, err := util.GenerateToken(as.conf, user, RefreshType)
	if err != nil {
		return "", "", err
	}

	accessToken, err := util.GenerateToken(as.conf, user, AccessType)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (as *AuthService) Refresh(ctx context.Context, refreshToken string) (string, error) {
	// parse refresh token
	claims, err := util.ParseToken(refreshToken, []byte(as.conf.RefreshToken))
	if err != nil {
		return "", err
	}

	user, err := as.userRepo.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}

	// generate access token
	return util.GenerateToken(as.conf, user, AccessType)
}
