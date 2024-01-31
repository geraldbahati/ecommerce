package utils

import (
	"errors"
	"fmt"
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
func GenerateTokens(userId uuid.UUID, username string, email string, role string) (string, string, time.Time, error) {
	// generate access token
	accessToken, err := generateAccessToken(userId, username, email, role)
	if err != nil {
		return "", "", time.Time{}, err
	}

	// generate refresh token
	refreshToken, expireTime, err := generateRefreshToken(userId, username, email, role)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, expireTime, nil
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

func generateRefreshToken(userId uuid.UUID, username string, email string, role string) (string, time.Time, error) {
	expireTime := time.Now().Add(24 * 90 * time.Hour) // expires in 90 days(3 months)

	// create claims
	claims := UserClaims{
		UserId:   userId,
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId.String(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	refreshToken, err := token.SignedString(jwtRefreshSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return refreshToken, expireTime, nil
}

// ParseToken parses the token and returns the claims
func ParseToken(tokenString string, isAccessToken bool) (*UserClaims, error) {
	var claims UserClaims
	var jwtSecret []byte

	if isAccessToken {
		jwtSecret = jwtAccessSecret
	} else {
		jwtSecret = jwtRefreshSecret
	}

	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// check if token is valid
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return &claims, nil
}

// RefreshToken generates a new access token
func RefreshToken(refreshToken string) (string, error) {
	// parse refresh token
	claims, err := ParseToken(refreshToken, false)
	if err != nil {
		return "", err
	}

	// validate
	if claims.UserId.String() == "" || claims.Username == "" || claims.Email == "" || claims.Role == "" {
		return "", errors.New("invalid token claims")
	}

	// generate new access token
	newAccessToken, err := generateAccessToken(claims.UserId, claims.Username, claims.Email, claims.Role)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

// ValidateToken validates the token
func ValidateToken(tokenString string, isAccessToken bool) error {
	var jwtSecret []byte

	if isAccessToken {
		jwtSecret = jwtAccessSecret
	} else {
		jwtSecret = jwtRefreshSecret
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	}

	// check if token is valid
	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}

// IsTokenExpired checks if the token is expired
func IsTokenExpired(tokenString string, isAccessToken bool) bool {
	var jwtSecret []byte

	if isAccessToken {
		jwtSecret = jwtAccessSecret
	} else {
		jwtSecret = jwtRefreshSecret
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return true
	}

	// check if token is valid
	if !token.Valid {
		return true
	}

	// check if token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		if time.Now().Unix() > int64(exp) {
			return true
		}
	}

	return false
}
