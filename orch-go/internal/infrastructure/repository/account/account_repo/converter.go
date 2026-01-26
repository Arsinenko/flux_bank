package account_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"
	"time"
)

func AccountToDomain(p *pb.AccountModel) *account.Account {
	if p == nil {
		return nil
	}
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.AsTime()
	}
	return &account.Account{
		Id:         &p.AccountId,
		CustomerId: *p.CustomerId,
		TypeId:     *p.TypeId,
		Iban:       p.Iban,
		Balance:    *p.Balance,
		CreatedAt:  createdAt,
		IsActive:   *p.IsActive,
	}
}

func AccountTypeToDomain(p *pb.AccountTypeModel) *account.AccountType {
	if p == nil {
		return nil
	}
	return &account.AccountType{
		Id:          &p.TypeId,
		Name:        p.Name,
		Description: p.Description,
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
