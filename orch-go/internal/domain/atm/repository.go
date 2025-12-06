package atm

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*Atm, error)
	GetById(ctx context.Context, id int32) (*Atm, error)
	Add(ctx context.Context, atm *Atm) (*Atm, error)
	Update(ctx context.Context, atm *Atm) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, atms []*Atm) error
	UpdateBulk(ctx context.Context, atms []*Atm) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
