package usecases

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
	firstName string,
	lastName string,
	userRole string,
) (model.User, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	// create user
	newUser := model.UserRegister{
		ID:             uuid.New(),
		Username:       username,
		Email:          email,
		HashedPassword: string(hashedPassword),
		FirstName:      firstName,
		LastName:       lastName,
		UserRole:       userRole,
	}

	return s.userRepo.CreateUser(ctx, newUser)
}
