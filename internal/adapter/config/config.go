package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		DB   *DB
		JWT  *JWT
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host           string
		Port           string
		AllowedOrigins string
		Email string
		EmailPassword string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	JWT struct {
		AccessToken   string
		RefreshToken string
		EmailToken string
		AccessTokenDuration string
		RefreshTokenDuration string
		EmailTokenDuration string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			errMsg := fmt.Errorf("unable to load .env: %v", err.Error())
			return nil, errMsg
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host:           os.Getenv("HTTP_HOST"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
		Email: os.Getenv("HTTP_EMAIL"),
		EmailPassword: os.Getenv("HTTP_EMAIL_PASSWORD"),
	}

	DB := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	JWT := &JWT{
		AccessToken:   os.Getenv("ACCESS_TOKEN"),
		RefreshToken:   os.Getenv("REFRESH_TOKEN"),
		EmailToken: os.Getenv("EMAIL_TOKEN"),
		AccessTokenDuration: os.Getenv("ACCESS_TOKEN_DURATION"),
		RefreshTokenDuration: os.Getenv("REFRESH_TOKEN_DURATION"),
		EmailTokenDuration: os.Getenv("EMAIL_TOKEN_DURATION"),
	}

	return &Container{
		App:  App,
		HTTP: HTTP,
		DB:   DB,
		JWT:  JWT,
	}, nil
}
