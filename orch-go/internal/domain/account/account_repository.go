package account

import "context"

type AccountRepository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]Account, error)
	GetById(ctx context.Context, id int32) (*Account, error)
	Create(ctx context.Context, account *Account) (*Account, error)
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id int32) error
	UpdateBulk(ctx context.Context, accounts []Account) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
