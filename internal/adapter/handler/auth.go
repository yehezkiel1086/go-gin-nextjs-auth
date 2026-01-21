package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/util"
)

type AuthHandler struct {
	conf *config.JWT
	svc port.AuthService
}

func NewAuthHandler(conf *config.JWT, svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		conf,
		svc,
	}
}

type LoginUserReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("email and password are required")})
		return
	}

	refreshToken, accessToken, err := ah.svc.Login(c, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": domain.ErrUnauthorized})
		return
	}

	// convert duration to int
	accessTokenDuration, err := strconv.Atoi(ah.conf.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": domain.ErrInternal,
		})
		return
	}

	refreshTokenDuration, err := strconv.Atoi(ah.conf.RefreshTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": domain.ErrInternal,
		})
		return
	}

	// set jwt token in cookie
	// c.SetCookie("access_token", accessToken, accessTokenDuration * 60, "/", "", false, true)
	c.SetCookie("access_token", accessToken, accessTokenDuration, "/", "", false, true)

	c.SetCookie("refresh_token", refreshToken, refreshTokenDuration * 24 * 60 * 60, "/api/v1/refresh", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (ah *AuthHandler) Refresh(c *gin.Context) {
	// get refresh token
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrUnauthorized)
		return
	}

	// get claims from refresh token
	claims, err := util.ParseToken(refreshToken, []byte(ah.conf.RefreshToken))
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrUnauthorized)
		return
	}

	// get user from email claims
	email := claims.Email

	user, err := ah.svc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrUnauthorized)
		return
	}

	// generate access token
	newAccessToken, err := util.GenerateAccessToken(ah.conf, user)

	accessTokenDuration, err := strconv.Atoi(ah.conf.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrUnauthorized)
		return
	}

	// c.SetCookie("access_token", newAccessToken, accessTokenDuration * 60, "/", "", false, true)
	c.SetCookie("access_token", newAccessToken, accessTokenDuration, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}