package card

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*Card, error)
	GetById(ctx context.Context, id int32) (*Card, error)
	Add(ctx context.Context, card *Card) (*Card, error)
	Update(ctx context.Context, card *Card) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, cards []*Card) error
	UpdateBulk(ctx context.Context, cards []*Card) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
