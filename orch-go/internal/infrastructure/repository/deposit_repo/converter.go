package deposit_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/deposit"
	"time"
)

func ToDomain(model *pb.DepositModel) *deposit.Deposit {
	if model == nil {
		return nil
	}
	var startDate time.Time
	var endDate time.Time
	if model.StartDate != nil {
		t := time.Date(int(model.StartDate.Year), time.Month(model.StartDate.Month), int(model.StartDate.Day), 0, 0, 0, 0, time.UTC)
		startDate = t
	}
	if model.EndDate != nil {
		t := time.Date(int(model.EndDate.Year), time.Month(model.EndDate.Month), int(model.EndDate.Day), 0, 0, 0, 0, time.UTC)
		endDate = t
	}

	return &deposit.Deposit{
		DepositID:    model.DepositId,
		CustomerID:   *model.CustomerId,
		Amount:       *model.Amount,
		InterestRate: *model.InterestRate,
		StartDate:    startDate,
		EndDate:      endDate,
		Status:       *model.Status,
	}

}

func ToDateOnly(t *time.Time) *pb.DateOnly {
	if t == nil {
		return nil
	}
	return &pb.DateOnly{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func ToDepositsDomain(resp *pb.GetAllDepositsResponse) []*deposit.Deposit {
	result := make([]*deposit.Deposit, 0, len(resp.Deposits))
	for _, d := range resp.Deposits {
		result = append(result, ToDomain(d))
	}
	return result
}
