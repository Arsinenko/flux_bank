package customer

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]Customer, error)
	GetById(ctx context.Context, id int32) (*Customer, error)
	GetBySubstring(ctx context.Context, request GetBySubStrRequest) ([]Customer, error)
	GetByDateRange(ctx context.Context, request GetByDateRangeRequest) ([]Customer, error)
	Create(ctx context.Context, customer *Customer) (*Customer, error)
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id int32) error
	UpdateBulk(ctx context.Context, customers []*Customer) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
