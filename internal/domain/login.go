package domain

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/romanzimoglyad/inquiry-backend/internal/auth"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) Login(ctx context.Context, request *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := i.database.GetUser(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("error in GetUser: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return nil, fmt.Errorf("error in CreateToken: %w", err)
	}

	return &domain.LoginResponse{
		Token: token,
	}, nil
}
