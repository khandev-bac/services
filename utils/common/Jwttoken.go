package common

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ACCESSTOKEN_KEY=
// REFRESHTOKEN_KEY=
var AccessToken = []byte(os.Getenv("ACCESSTOKEN_KEY"))
var RefreshToken = []byte(os.Getenv("REFRESHTOKEN_KEY"))

func GenerateToken(payload Payloads) *Token {
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  payload.Id,
		"username": payload.Username,
		"email":    payload.Email,
		"role":     payload.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	access_token_string, err := access_token.SignedString(AccessToken)
	if err != nil {
		// logger.Error("Failed to generate accesstoken", zap.Error(err))
		return nil
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  payload.Id,
		"username": payload.Username,
		"email":    payload.Email,
		"role":     payload.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	refresh_token_string, err := refresh_token.SignedString(RefreshToken)
	if err != nil {
		// logger.Error("Failed to generate refreshtoken", zap.Error(err))
		return nil
	}
	return &Token{
		AccessToken:  access_token_string,
		RefreshToken: refresh_token_string,
	}
}
func VerifyAccessToken(accessToken string) (*Payloads, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		return AccessToken, nil
	})
	if err != nil {
		// logger.Error("Failed to verify accessToken", zap.Error(err))
		return nil, fmt.Errorf("Error while verifying accessToken")
	}
	if !token.Valid {
		// logger.Error("accessToken is invalid")
		return nil, fmt.Errorf("Invalid accessToken")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token")
	}
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims: %w", err)
	}
	var payload Payloads
	if err := json.Unmarshal(claimsJson, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims into struct: %w", err)
	}
	idStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'user_id' claim")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format in 'user_id': %w", err)
	}
	payload.Id = id
	return &payload, nil
}
func VerifyRefreshToken(refreshToken string) (*Payloads, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (any, error) {
		return RefreshToken, nil
	})
	if err != nil {
		// logger.Error("Failed to verify accessToken", zap.Error(err))
		return nil, fmt.Errorf("Error while verifying refreshToken")
	}
	if !token.Valid {
		// logger.Error("accessToken is invalid")
		return nil, fmt.Errorf("Invalid refreshToken")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token")
	}
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims: %w", err)
	}
	var payload Payloads
	if err := json.Unmarshal(claimsJson, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims into struct: %w", err)
	}
	idStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'user_id' claim")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format in 'user_id': %w", err)
	}
	payload.Id = id
	return &payload, nil
}
