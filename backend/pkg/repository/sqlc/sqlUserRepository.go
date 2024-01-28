package sqlc

import (
	"context"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
)

type SQLUserRepository struct {
	DB *database.Queries
}

func NewSQLUserRepository(db *database.Queries) *SQLUserRepository {
	return &SQLUserRepository{
		DB: db,
	}
}

// User queries
// CreateUser creates a new user
func (r *SQLUserRepository) CreateUser(ctx context.Context, user model.UserRegister) (model.User, error) {
	// insert user into database
	createdUser, err := r.DB.CreateUser(ctx, database.CreateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		UserRole:       user.UserRole,
	})
	if err != nil {
		return model.User{}, err
	}

	// return created user
	return model.User{
		ID:              createdUser.ID,
		Username:        createdUser.Username,
		Email:           createdUser.Email,
		HashedPassword:  createdUser.HashedPassword,
		FirstName:       createdUser.FirstName,
		LastName:        createdUser.LastName,
		PhoneNumber:     createdUser.PhoneNumber,
		DateOfBirth:     createdUser.DateOfBirth,
		Gender:          createdUser.Gender,
		ShippingAddress: createdUser.ShippingAddress,
		BillingAddress:  createdUser.BillingAddress,
		CreatedAt:       createdUser.CreatedAt,
		LastLogin:       createdUser.LastLogin,
		AccountStatus:   createdUser.AccountStatus,
		UserRole:        createdUser.UserRole,
		ProfilePicture:  createdUser.ProfilePicture,
		TwoFactorAuth:   createdUser.TwoFactorAuth,
	}, nil
}
