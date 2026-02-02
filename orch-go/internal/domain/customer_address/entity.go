package customer_address

import "github.com/brianvoe/gofakeit/v7"

type CustomerAddress struct {
	Id         *int32 `json:"id"`
	CustomerId int32  `json:"customer_id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Street     string `json:"street"`
	ZipCode    string `json:"zip_code"`
	IsPrimary  bool   `json:"is_primary"`
}

func FakeCustomerAddress(customerId int32) *CustomerAddress {
	gofakeit.New(0)
	return &CustomerAddress{
		CustomerId: customerId,
		Country:    gofakeit.Country(),
		City:       gofakeit.City(),
		Street:     gofakeit.Street(),
		ZipCode:    gofakeit.Zip(),
		IsPrimary:  gofakeit.Bool(),
	}
}
