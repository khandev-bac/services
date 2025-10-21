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
