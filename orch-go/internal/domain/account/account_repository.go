package account

import "context"

type AccountRepository interface {
	GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]Account, error)
	GetById(ctx context.Context, id int32) (*Account, error)
	GetByCustomerId(ctx context.Context, customerId int32) ([]*Account, error)
	GetByDateRange(ctx context.Context, request GetByDateRange) ([]*Account, error)
	Create(ctx context.Context, account *Account) (*Account, error)
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id int32) error
	UpdateBulk(ctx context.Context, accounts []Account) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
