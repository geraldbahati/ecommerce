package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.UserRegister) (model.User, error)
	CountAllUsersByUsername(ctx context.Context, username string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	StoreRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string, expiresAt time.Time) (model.RefreshToken, error)
	UpdateUserLastLogin(ctx context.Context, userId uuid.UUID) error
	GetUserById(ctx context.Context, userId uuid.UUID) (model.User, error)
}
