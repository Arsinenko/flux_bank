package account

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Account struct {
	Id         *int32
	CustomerId int32
	TypeId     int32
	Iban       string
	Balance    string
	CreatedAt  time.Time
	IsActive   bool
}

type AccountType struct {
	Id          *int32
	Name        string
	Description *string
}

type GetByDateRange struct {
	From, To        time.Time
	PageN, PageSize int32
}

func FakeAccount(customerId int32, typeId int32) *Account {
	gofakeit.New(0)
	return &Account{
		Id:         nil,
		CustomerId: customerId,
		TypeId:     typeId,
		Iban:       "Account Iban",
		Balance:    string(rune(gofakeit.Number(1000, 10000))),
		CreatedAt:  time.Now(),
		IsActive:   true,
	}
}

func FakeAccountType(name string) AccountType {
	return AccountType{
		Id:          nil,
		Name:        name,
		Description: nil,
	}
}
