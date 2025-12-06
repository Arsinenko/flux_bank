package login_log

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]*LoginLog, error)
	GetById(ctx context.Context, id int32) (*LoginLog, error)
	Create(ctx context.Context, loginLog *LoginLog) (*LoginLog, error)
	Update(ctx context.Context, loginLog *LoginLog) error
	Delete(ctx context.Context, id int32) error
}
