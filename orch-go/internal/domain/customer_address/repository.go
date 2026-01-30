package customer_address

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]CustomerAddress, error)
	GetById(ctx context.Context, id int32) (*CustomerAddress, error)
	Create(ctx context.Context, address *CustomerAddress) (*CustomerAddress, error)
	Update(ctx context.Context, address *CustomerAddress) error
	Delete(ctx context.Context, id int32) error
}
