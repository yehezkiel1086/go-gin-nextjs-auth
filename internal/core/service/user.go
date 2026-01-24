package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/util"
)

type UserService struct {
	httpConf *config.HTTP
	jwtConf *config.JWT
	repo port.UserRepository
}

func NewUserService(httpConf *config.HTTP, jwtConf *config.JWT, repo port.UserRepository) (*UserService) {
	return &UserService{
		httpConf,
		jwtConf,
		repo,
	}
}

func (us *UserService) RegisterUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
	// hash password
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPwd

	// generate email token
	token, err := util.GenerateToken(us.jwtConf, user, "email")
	if err != nil {
		return nil, err
	}

	user.VerificationToken = token

	// send verification email
	util.SendConfirmationEmail(us.httpConf, user.Email, token, us.jwtConf.EmailTokenDuration)

	return us.repo.CreateUser(ctx, user)
}

func (us *UserService) ConfirmEmail(ctx context.Context, token string) error {
	// parse token
	claims, err := util.ParseToken(token, []byte(us.jwtConf.EmailToken))
	if err != nil {
		return err
	}

	// verify email
	user, err := us.repo.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return err
	}

	user.VerificationToken = ""
	user.IsVerified = true

	if _, err := us.repo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUsers(ctx context.Context) ([]domain.UserResponse, error) {
	return us.repo.GetUsers(ctx)
}
