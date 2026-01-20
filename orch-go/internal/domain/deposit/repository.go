package deposit

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*Deposit, error)
	GetById(ctx context.Context, id int32) (*Deposit, error)
	GetByCustomer(ctx context.Context, customerId int32) ([]*Deposit, error)
	Add(ctx context.Context, deposit *Deposit) (*Deposit, error)
	Update(ctx context.Context, deposit *Deposit) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, deposits []*Deposit) error
	UpdateBulk(ctx context.Context, deposits []*Deposit) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
