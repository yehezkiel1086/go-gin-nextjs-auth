package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

func GenerateToken(conf *config.JWT, user *domain.User, tokenType string) (string, error) {
	var mySigningKey []byte
	var expiry *jwt.NumericDate

	switch tokenType {
		case "access":
			mySigningKey = []byte(conf.AccessToken)
			duration, err := strconv.Atoi(conf.AccessTokenDuration)
			if err != nil {
				return "", err
			}
			expiry = jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Second))
		case "refresh":
			mySigningKey = []byte(conf.RefreshToken)
			duration, err := strconv.Atoi(conf.RefreshTokenDuration)
			if err != nil {
				return "", err
			}
			expiry = jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Hour * 24))
	}

	claims := domain.JWTClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiry,
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
