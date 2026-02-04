package account

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Account struct {
	Id         *int32    `json:"id"`
	CustomerId int32     `json:"customer_id"`
	TypeId     int32     `json:"type_id"`
	Iban       string    `json:"iban"`
	Balance    string    `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	IsActive   bool      `json:"is_active"`
}

type AccountType struct {
	Id          *int32  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
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
		Iban:       gofakeit.AchAccount(),
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
