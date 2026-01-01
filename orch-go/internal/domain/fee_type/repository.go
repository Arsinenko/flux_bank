package fee_type

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]*FeeType, error)
	GetById(ctx context.Context, id int32) (*FeeType, error)
	Add(ctx context.Context, feeType *FeeType) (*FeeType, error)
	Update(ctx context.Context, feeType *FeeType) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, feeTypes []*FeeType) error
	UpdateBulk(ctx context.Context, feeTypes []*FeeType) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
