package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
)

type UserRepository interface {
	// User queries
	CreateUser(ctx context.Context, user model.UserRegister) (model.User, error)
}
