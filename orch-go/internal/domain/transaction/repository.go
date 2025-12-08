package transaction

import "context"

type TransactionRepository interface {
	GetAll(ctx context.Context) ([]*Transaction, error)
	GetById(ctx context.Context, id int32) (*Transaction, error)
	GetByDateRange(ctx context.Context, request GetByDateRange) ([]*Transaction, error)
	Add(ctx context.Context, transaction *Transaction) (*Transaction, error)
	Update(ctx context.Context, transaction *Transaction) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, transactions []*Transaction) error
	UpdateBulk(ctx context.Context, transactions []*Transaction) error
	DeleteBulk(ctx context.Context, ids []int32) error
}

type TransactionCategoryRepository interface {
	GetAll(ctx context.Context) ([]*TransactionCategory, error)
	GetById(ctx context.Context, id int32) (*TransactionCategory, error)
	Add(ctx context.Context, transactionCategory *TransactionCategory) (*TransactionCategory, error)
	Update(ctx context.Context, transactionCategory *TransactionCategory) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, transactionCategories []*TransactionCategory) error
	UpdateBulk(ctx context.Context, transactionCategories []*TransactionCategory) error
	DeleteBulk(ctx context.Context, ids []int32) error
}

type TransactionFeeRepository interface {
	GetAll(ctx context.Context) ([]*TransactionFee, error)
	GetById(ctx context.Context, id int32) (*TransactionFee, error)
	Add(ctx context.Context, transactionFee *TransactionFee) (*TransactionFee, error)
	Update(ctx context.Context, transactionFee *TransactionFee) error
	Delete(ctx context.Context, id int32) error
	AddBulk(ctx context.Context, transactionFees []*TransactionFee) error
	UpdateBulk(ctx context.Context, transactionFees []*TransactionFee) error
	DeleteBulk(ctx context.Context, ids []int32) error
}
