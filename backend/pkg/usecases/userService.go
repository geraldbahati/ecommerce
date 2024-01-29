package usecases

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"

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

	// generate username
	username, err := generateUsername(ctx, s.userRepo, firstName, lastName)
	if err != nil {
		return model.User{}, err
	}

	// create user
	newUser := model.UserRegister{
		ID:             uuid.New(),
		Username:       username,
		Email:          email,
		HashedPassword: string(hashedPassword),
		FirstName:      strings.ToLower(firstName),
		LastName:       strings.ToLower(lastName),
		UserRole:       userRole,
	}

	return s.userRepo.CreateUser(ctx, newUser)
}

// sanitizeUsername sanitizes the given username
func sanitizeUsername(username string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9_.-]+")
	return reg.ReplaceAllString(username, "")
}

// generateUsername generates a username from the given first and last name
func generateUsername(ctx context.Context, userRepo repository.UserRepository, firstName string, lastName string) (string, error) {
	// generate username from first and last name
	baseUsername := strings.ToLower(sanitizeUsername(firstName + lastName))
	username := baseUsername
	const maxUsernameLength = 20

	// trim username to max length
	if len(baseUsername) > maxUsernameLength {
		baseUsername = baseUsername[:maxUsernameLength]
		username = baseUsername
	}

	for suffix := 1; ; suffix++ {
		// check if username is available
		count, err := userRepo.CountAllUsersByUsername(ctx, username)
		if err != nil {
			return "", err
		}
		if count == 0 {
			return username, nil
		}

		// append suffix to username
		suffixStr := strconv.Itoa(suffix)
		cutOffLength := maxUsernameLength - len(suffixStr)
		if cutOffLength < len(baseUsername) {
			username = baseUsername[:cutOffLength]
		}

		username += suffixStr
	}
}
