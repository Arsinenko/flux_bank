package branch

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]*Branch, error)
	GetById(ctx context.Context, id int32) (*Branch, error)
	Add(ctx context.Context, branch *Branch) (*Branch, error)
	Update(ctx context.Context, branch *Branch) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, branches []*Branch) error
	UpdateBulk(ctx context.Context, branches []*Branch) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
