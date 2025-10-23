package common

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}
type KafkaEvent struct {
	EventType string          `json:"event_type"` // Must match producer
	Timestamp time.Time       `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
}
type KafkaDeleteEvent struct {
	UserId uuid.UUID `json:"user_id"`
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
	Username     string    `json:"username"`
	Picture      *string   `json:"picture"`
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
type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Refreshes struct {
	RefreshToken string
}
type KafkaSendValues struct {
	UserId   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Picture  *string   `json:"picture"`
	Role     *string   `json:"user_role"`
	Revoked  *bool     `json:"revoked"`
}
