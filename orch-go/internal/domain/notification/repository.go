package notification

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]*Notification, error)
	GetById(ctx context.Context, id int32) (*Notification, error)
	Add(ctx context.Context, notification *Notification) (*Notification, error)
	Update(ctx context.Context, notification *Notification) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, notifications []*Notification) error
	UpdateBulk(ctx context.Context, notifications []*Notification) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
