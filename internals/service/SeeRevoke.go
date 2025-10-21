package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (as *AuthService) SeeRevoke(ctx context.Context, id uuid.UUID) (bool, error) {
	revoked, err := as.repo.SeeRevoke(ctx, id)
	if err != nil {
		return false, errors.New("error while revoking")
	}
	return revoked.Bool, nil
}
