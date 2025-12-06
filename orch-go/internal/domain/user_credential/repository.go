package user_credential

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*UserCredential, error)
	GetById(ctx context.Context, id int32) (*UserCredential, error)
	Create(ctx context.Context, cred *UserCredential) (*UserCredential, error)
	Update(ctx context.Context, cred UserCredential) error
	Delete(ctx context.Context, customerId int32) error
}
