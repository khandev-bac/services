package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (as *AuthService) DeleteUserAccount(ctx context.Context, id uuid.UUID) (string, error) {
	userFound, err := as.repo.FindById(ctx, id)
	if err == nil && userFound.Email == "" {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", errors.New("something went wrong")
	}
	err = as.repo.DeleteUser(ctx, id)
	if err != nil {
		return "", errors.New("failed to delete user")
	}
	return "Successfully deleted", nil
}
