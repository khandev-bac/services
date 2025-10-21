package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/services/db-models"
	"github.com/services/internals/repository"
	"github.com/services/utils/common"
)

type AuthService struct {
	repo *repository.DBQueries
}

func NewAuthService(repo *repository.DBQueries) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) SignUp(ctx context.Context, username, email, password string) (*common.UserResponse, error) {
	_, err := as.repo.FindByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("user already found. Please login")
	}
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}

	hashedPass, err := common.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newUser, err := as.repo.CreateUser(ctx, db.SingupParams{
		Username: sql.NullString{String: username, Valid: username != ""},
		Email:    email,
		Password: hashedPass,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	tokens := common.GenerateToken(common.Payloads{
		Id:       newUser.ID,
		Email:    newUser.Email,
		Username: newUser.Username.String,
	})

	return &common.UserResponse{
		ID:           newUser.ID,
		Email:        newUser.Email,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
