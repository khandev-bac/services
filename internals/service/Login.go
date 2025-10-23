package service

import (
	"context"
	"errors"

	"github.com/services/utils/common"
)

func (as *AuthService) Login(ctx context.Context, email, password string) (*common.UserResponse, error) {

	userFound, err := as.repo.FindFullWithEmail(ctx, email)
	if err == nil && userFound.Email == "" {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := common.CheckPasswordHash(password, userFound.Password); err != nil {
		return nil, errors.New("invalid password or email")
	}
	tokens, _ := common.GenerateToken(common.Payloads{
		Id:       userFound.ID,
		Email:    userFound.Email,
		Username: userFound.Username.String,
	})
	return &common.UserResponse{
		ID:           userFound.ID,
		Email:        userFound.Email,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
