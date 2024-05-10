package auth

import (
	"fmt"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/config"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	"github.com/golang-jwt/jwt/v5"
)

type Role int32

const (
	RoleAdmin Role = 1
	RoleUser  Role = 2
)

type CustomClaims struct {
	Username string `json:"user_name"`
	UserId   int32  `json:"user_id"`
	Role     int32  `json:"role"`
	jwt.RegisteredClaims
}

// CreateToken Function to create JWT tokens with claims
func CreateToken(user *domain.User) (string, error) {
	// Create a new JWT token with token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserId:   user.Id,
		Username: user.Name,
		Role:     user.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "inquiry-app",
		},
	})

	// Print information about the created token
	logger.Info().Msgf("Token token added: %+v\n", token)
	tokenString, err := token.SignedString([]byte(config.Config.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken function to validate JWT token
func ValidateToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Auth.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in ParseWithClaims: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token.Claims.(*CustomClaims), nil
}
