package common

import "github.com/google/uuid"

type Token struct {
	AccessToken  string
	RefreshToken string
}
type Payloads struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     *string   `json:"role"`
}
type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}
type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}
type SuccessResponse struct {
	Message string       `json:"message"`
	Code    int          `json:"code" `
	Data    UserResponse `json:"data"`
}
type UserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
