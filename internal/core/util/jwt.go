package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

func GenerateAccessToken(conf *config.JWT, user *domain.User) (string, error) {
	mySigningKey := []byte(conf.AccessToken)

	// convert duration to int
	duration, err := strconv.Atoi(conf.AccessTokenDuration)
	if err != nil {
		return "", err
	}

	// Create claims with multiple fields populated
	claims := domain.JWTClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func GenerateRefreshToken(conf *config.JWT, user *domain.User) (string, error) {
	mySigningKey := []byte(conf.RefreshToken)

	// convert duration to int
	duration, err := strconv.Atoi(conf.RefreshTokenDuration)
	if err != nil {
		return "", err
	}

	// Create claims with multiple fields populated
	claims := domain.JWTClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseToken(tokenStr string, secret []byte) (*domain.JWTClaims, error) {
	// parse token
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// extract claims
	claims, ok := token.Claims.(*domain.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, err
}
