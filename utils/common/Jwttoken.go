package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var Publickey *rsa.PublicKey

func InitPublicKey() {
	key, err := os.ReadFile("public.pem")
	if err != nil {
		log.Println("Failed to read publicket ")
		return
	}
	extrat, _ := pem.Decode(key)
	if extrat == nil || extrat.Type != "PUBLIC KEY" {
		log.Fatalf("Failed to decode PEM block containing public key")
	}
	keyParsed, err := x509.ParsePKIXPublicKey(extrat.Bytes)
	if err != nil {
		log.Println("Failed to parse publickey ")
		return
	}
	Publickey = keyParsed.(*rsa.PublicKey)
}

var privatekey *rsa.PrivateKey

func InitKey() {
	key, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatalf("Could not read private key: %v", err)
	}
	extract, _ := pem.Decode(key)
	if extract == nil || extract.Type != "RSA PRIVATE KEY" {
		log.Fatalf("Failed to decode PEM block containing private key")
	}
	prsedkey, err := x509.ParsePKCS1PrivateKey(extract.Bytes)
	if err != nil {
		log.Fatalf("Could not parse RSA private key: %v", err)
	}
	privatekey = prsedkey
	// fmt.Println("prvate key : ", privatekey)
}

func GenerateToken(payload Payloads) (*Token, error) {
	access_token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":  payload.Id,
		"username": payload.Username,
		"email":    payload.Email,
		"role":     payload.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	access_token_string, err := access_token.SignedString(privatekey)
	if err != nil {
		return nil, err
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":  payload.Id,
		"username": payload.Username,
		"email":    payload.Email,
		"role":     payload.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	refresh_token_string, err := refresh_token.SignedString(privatekey)
	if err != nil {
		// logger.Error("Failed to generate refreshtoken", zap.Error(err))
		return nil, err
	}
	return &Token{
		AccessToken:  access_token_string,
		RefreshToken: refresh_token_string,
	}, nil
}
func VerifyAccessToken(accessToken string) (*Payloads, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		return Publickey, nil
	})
	if err != nil {
		// logger.Error("Failed to verify accessToken", zap.Error(err))
		return nil, fmt.Errorf("error while verifying accessToken")
	}
	if !token.Valid {
		// logger.Error("accessToken is invalid")
		return nil, fmt.Errorf("invalid accessToken")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
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
		return Publickey, nil
	})
	if err != nil {
		// logger.Error("Failed to verify accessToken", zap.Error(err))
		return nil, fmt.Errorf("error while verifying refreshToken")
	}
	if !token.Valid {
		// logger.Error("accessToken is invalid")
		return nil, fmt.Errorf("invalid refreshToken")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
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
