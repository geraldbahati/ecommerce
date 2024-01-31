package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
	"time"
)

type UserRepository interface {
	// create
	CreateUser(ctx context.Context, user model.UserRegister) (model.User, error)

	// update
	StoreRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string, expiresAt time.Time) (model.RefreshToken, error)
	UpdateUserLastLogin(ctx context.Context, userId uuid.UUID) error
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUserProfilePicture(ctx context.Context, user model.User) (model.User, error)

	// delete

	// get
	CountAllUsersByUsername(ctx context.Context, username string) (int64, error)
	GetUserById(ctx context.Context, userId uuid.UUID) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}
