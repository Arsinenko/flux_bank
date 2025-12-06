package card

import "time"

type Card struct {
	CardID     int32
	AccountID  *int32
	CardNumber string
	CVV        string
	ExpiryDate *time.Time
	Status     *string
}
