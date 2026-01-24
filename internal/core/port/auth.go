package port

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, error)
}
