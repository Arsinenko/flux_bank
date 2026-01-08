package customer

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Customer struct {
	Id        int32
	FirstName string
	LastName  string
	Email     string
	Phone     *string
	BirthDate *time.Time
	CreatedAt *time.Time
}

type GetBySubStrRequest struct {
	SubStr          string
	PageN, PageSize int32
	Order           string
	Desk            bool
}

type GetByDateRangeRequest struct {
	From, To        time.Time
	PageN, PageSize int32
}

func FakeCustomer(id int32, createdAt time.Time) *Customer {
	gofakeit.New(0)
	phone := gofakeit.Phone()
	bDate := gofakeit.Date()
	return &Customer{
		Id:        id,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     &phone,
		BirthDate: &bDate,
		CreatedAt: &createdAt,
	}
}
