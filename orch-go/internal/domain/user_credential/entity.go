package user_credential

import "time"

type UserCredential struct {
	CustomerId   int32
	Username     string
	PasswordHash string
	UpdatedAt    time.Time
}

type GetBySubStrRequest struct {
	subStr          string
	pageN, pageSize int32
	order           string
	desk            bool
}
