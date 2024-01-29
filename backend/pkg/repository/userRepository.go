package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.UserRegister) (model.User, error)
	CountAllUsersByUsername(ctx context.Context, username string) (int64, error)
}
