package sqlc

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
	"log"
	"time"
)

type SQLUserRepository struct {
	DB *database.Queries
}

func NewSQLUserRepository(db *database.Queries) *SQLUserRepository {
	return &SQLUserRepository{
		DB: db,
	}
}

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
		ID:             createdUser.ID,
		Username:       createdUser.Username,
		Email:          createdUser.Email,
		HashedPassword: createdUser.HashedPassword,
		FirstName:      createdUser.FirstName,
		LastName:       createdUser.LastName,
		PhoneNumber:    createdUser.PhoneNumber,
		DateOfBirth:    createdUser.DateOfBirth,
		Gender:         createdUser.Gender,
		CreatedAt:      createdUser.CreatedAt,
		LastLogin:      createdUser.LastLogin,
		AccountStatus:  createdUser.AccountStatus,
		UserRole:       createdUser.UserRole,
		ProfilePicture: createdUser.ProfilePicture,
		TwoFactorAuth:  createdUser.TwoFactorAuth,
	}, nil
}

// CountAllUsersByUsername returns the number of users with the given username
func (r *SQLUserRepository) CountAllUsersByUsername(ctx context.Context, username string) (int64, error) {
	return r.DB.CountAllUsersByUsername(ctx, username)
}

// GetUserByEmail returns the user with the given email
func (r *SQLUserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	// get user from database
	user, err := r.DB.FindUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	// return user
	return model.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		DateOfBirth:    user.DateOfBirth,
		Gender:         user.Gender,
		CreatedAt:      user.CreatedAt,
		LastLogin:      user.LastLogin,
		AccountStatus:  user.AccountStatus,
		UserRole:       user.UserRole,
		ProfilePicture: user.ProfilePicture,
		TwoFactorAuth:  user.TwoFactorAuth,
	}, nil
}

// StoreRefreshToken stores the refresh token in the database
func (r *SQLUserRepository) StoreRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string, expiresAt time.Time) (model.RefreshToken, error) {
	log.Printf("Storing refresh token for user with id %s", userId.String())
	// insert refresh token into database
	createdRefreshToken, err := r.DB.StoreRefreshToken(ctx, database.StoreRefreshTokenParams{
		ID:        uuid.New(),
		UserID:    userId,
		Token:     refreshToken,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	})
	if err != nil {
		log.Printf("Error storing refresh token for user with id %s: %s", userId.String(), err.Error())
		return model.RefreshToken{}, err
	}

	// return refresh token
	log.Printf("Successfully stored refresh token for user with id %s", userId.String())
	return model.RefreshToken{
		ID:        createdRefreshToken.ID,
		UserID:    createdRefreshToken.UserID,
		Token:     createdRefreshToken.Token,
		CreatedAt: createdRefreshToken.CreatedAt,
		ExpiresAt: createdRefreshToken.ExpiresAt,
		RevokedAt: createdRefreshToken.RevokedAt,
	}, nil
}

// UpdateUserLastLogin updates the last login of the user
func (r *SQLUserRepository) UpdateUserLastLogin(ctx context.Context, userId uuid.UUID) error {
	log.Printf("Updating last login of user with id %s", userId.String())

	err := r.DB.UpdateUserLastLogin(ctx, database.UpdateUserLastLoginParams{
		ID:        userId,
		LastLogin: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		log.Printf("Error updating last login of user with id %s: %s", userId.String(), err.Error())
	}
	return err
}

// GetUserById returns the user with the given id
func (r *SQLUserRepository) GetUserById(ctx context.Context, userId uuid.UUID) (model.User, error) {
	// get user from database
	user, err := r.DB.FindUserByID(ctx, userId)
	if err != nil {
		return model.User{}, err
	}

	// return user
	return model.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		DateOfBirth:    user.DateOfBirth,
		Gender:         user.Gender,
		CreatedAt:      user.CreatedAt,
		LastLogin:      user.LastLogin,
		AccountStatus:  user.AccountStatus,
		UserRole:       user.UserRole,
		ProfilePicture: user.ProfilePicture,
		TwoFactorAuth:  user.TwoFactorAuth,
	}, nil
}

// UpdateUser updates a user
func (r *SQLUserRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	log.Printf("Updating user with id %s", user.ID.String())

	// update user
	updatedUser, err := r.DB.UpdateUser(ctx, database.UpdateUserParams{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
	})
	if err != nil {
		log.Printf("Error updating user with id %s: %s", user.ID.String(), err.Error())
		return model.User{}, err
	}

	// return updated user
	return model.User{
		ID:             updatedUser.ID,
		Username:       updatedUser.Username,
		Email:          updatedUser.Email,
		HashedPassword: updatedUser.HashedPassword,
		FirstName:      updatedUser.FirstName,
		LastName:       updatedUser.LastName,
		PhoneNumber:    updatedUser.PhoneNumber,
		DateOfBirth:    updatedUser.DateOfBirth,
		Gender:         updatedUser.Gender,
		CreatedAt:      updatedUser.CreatedAt,
		LastLogin:      updatedUser.LastLogin,
		AccountStatus:  updatedUser.AccountStatus,
		UserRole:       updatedUser.UserRole,
		ProfilePicture: updatedUser.ProfilePicture,
		TwoFactorAuth:  updatedUser.TwoFactorAuth,
	}, nil
}
