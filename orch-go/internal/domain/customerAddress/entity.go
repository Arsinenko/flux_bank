package customerAdress

import "github.com/brianvoe/gofakeit/v7"

type CustomerAddress struct {
	Id         *int32
	CustomerId int32
	Country    string
	City       string
	Street     string
	ZipCode    string
	IsPrimary  bool
}

func FakeCustomerAddress(id, customerId int32) *CustomerAddress {
	gofakeit.New(0)
	return &CustomerAddress{
		Id:         &id,
		CustomerId: customerId,
		Country:    gofakeit.Country(),
		City:       gofakeit.City(),
		Street:     gofakeit.Street(),
		ZipCode:    gofakeit.Zip(),
		IsPrimary:  gofakeit.Bool(),
	}
}
