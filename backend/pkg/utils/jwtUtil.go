package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	jwtAccessSecret  = []byte("eK_sS2AgDstNYrh0Bx5LK3nPx-z1h2l_ZdjchgQjvyA=")
	jwtRefreshSecret = []byte("bS4RqAvfuWhiAjZJ_104wBUcDAbp4cEt2ChP1IYskI8=")
)

type UserClaims struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

// GenerateTokens generates access and refresh tokens
func GenerateTokens(userId uuid.UUID, username string, email string, role string) (string, string, error) {
	// generate access token
	accessToken, err := generateAccessToken(userId, username, email, role)
	if err != nil {
		return "", "", err
	}

	// generate refresh token
	refreshToken, err := generateRefreshToken(userId, username, email, role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(userId uuid.UUID, username string, email string, role string) (string, error) {
	// create claims
	claims := UserClaims{
		UserId:   userId,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId.String(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	return token.SignedString(jwtAccessSecret)
}

func generateRefreshToken(userId uuid.UUID, username string, email string, role string) (string, error) {
	// create claims
	claims := UserClaims{
		UserId:   userId,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 90 * time.Hour)), // expires in 90 days(3 months)
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId.String(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	return token.SignedString(jwtRefreshSecret)
}
