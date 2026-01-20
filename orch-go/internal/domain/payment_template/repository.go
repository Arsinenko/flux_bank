package payment_template

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*PaymentTemplate, error)
	GetById(ctx context.Context, id int32) (*PaymentTemplate, error)
	Add(ctx context.Context, paymentTemplate *PaymentTemplate) (*PaymentTemplate, error)
	Update(ctx context.Context, paymentTemplate *PaymentTemplate) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, paymentTemplates []*PaymentTemplate) error
	UpdateBulk(ctx context.Context, paymentTemplates []*PaymentTemplate) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
