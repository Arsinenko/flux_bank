package user_credential

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
)

type UserCredential struct {
	CustomerId   *int32
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

func FakeUserCreds() UserCredential {
	gofakeit.New(0)

	password := gofakeit.Password(true, true, true, true, false, 10)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return UserCredential{
		CustomerId:   nil,
		Username:     gofakeit.Username(),
		PasswordHash: string(passwordHash),
		UpdatedAt:    time.Time{},
	}
}
