package atm

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
)

type Atm struct {
	AtmID    *int32
	BranchID int32
	Location *string
	Status   *string
}

func FakeAtm(branchId int32) Atm {
	gofakeit.New(0)
	location := gofakeit.Address().Address
	status := strconv.FormatBool(gofakeit.Bool())
	return Atm{
		AtmID:    nil,
		BranchID: branchId,
		Location: &location,
		Status:   &status,
	}
}
