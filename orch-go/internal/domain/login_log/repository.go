package login_log

import (
	"context"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*LoginLog, error)
	GetById(ctx context.Context, id int32) (*LoginLog, error)
	GetByCustomer(ctx context.Context, customerId int32) ([]*LoginLog, error)
	GetInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*LoginLog, error)
	Create(ctx context.Context, loginLog *LoginLog) (*LoginLog, error)
	Update(ctx context.Context, loginLog *LoginLog) error
	Delete(ctx context.Context, id int32) error
}
