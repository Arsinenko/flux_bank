package branch

import "github.com/brianvoe/gofakeit/v7"

type Branch struct {
	BranchID *int32
	Name     *string
	City     *string
	Address  *string
	Phone    *string
}

func FakeBranch() Branch {
	gofakeit.New(0)
	name := gofakeit.Name()
	city := gofakeit.City()
	address := gofakeit.Address().Address
	phone := gofakeit.Phone()

	return Branch{
		BranchID: nil,
		Name:     &name,
		City:     &city,
		Address:  &address,
		Phone:    &phone,
	}
}
