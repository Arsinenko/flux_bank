package card

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Card struct {
	CardID     int32
	AccountID  *int32
	CardNumber string
	CVV        string
	ExpiryDate *time.Time
	Status     *string
}

func FakeCard(cardId, accountId int32, expiryDate time.Time, status string) *Card {
	gofakeit.New(0)

	return &Card{
		CardID:     cardId,
		AccountID:  &accountId,
		CardNumber: gofakeit.CreditCardNumber(nil),
		CVV:        gofakeit.CreditCardCvv(),
		ExpiryDate: &expiryDate,
		Status:     &status,
	}

}
