package user_credential

import "context"

type Repository interface {
	GetAll(ctx context.Context, pageN, pageSize int32) ([]*UserCredential, error)
	GetById(ctx context.Context, id int32) (*UserCredential, error)
	GetByUsername(ctx context.Context, username string) (*UserCredential, error)
	Create(ctx context.Context, cred *UserCredential) (*UserCredential, error)
	Update(ctx context.Context, cred UserCredential) error
	Delete(ctx context.Context, customerId int32) error
}
