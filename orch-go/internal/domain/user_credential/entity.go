package user_credential

import "time"

type UserCredential struct {
	CustomerId   int32
	Username     string
	PasswordHash string
	UpdatedAt    time.Time
}
