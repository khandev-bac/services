package common

import "github.com/google/uuid"

type Token struct {
	AccessToken  string
	RefreshToken string
}
type Payloads struct {
	Id       uuid.UUID
	Email    string
	Username string
	Role     *string
}
