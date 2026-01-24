package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.UserResponse, error)
}

type UserService interface {
	RegisterUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error)
	GetUsers(ctx context.Context) ([]domain.UserResponse, error)
	ConfirmEmail(ctx context.Context, token string) error
}
