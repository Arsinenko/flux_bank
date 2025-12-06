package account

import "context"

type AccountTypeRepository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]AccountType, error)
	GetById(ctx context.Context, id int32) (*AccountType, error)
	Create(ctx context.Context, accountType *AccountType) (*AccountType, error)
	Update(ctx context.Context, accountType *AccountType) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, accountTypes []AccountType) error
	UpdateBulk(ctx context.Context, accountTypes []AccountType) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
