package services

import (
	"context"
	"orch-go/internal/domain/user_credential"
	"orch-go/internal/infrastructure/repository/user_credential_repo"

	"golang.org/x/crypto/bcrypt"
)

type UserCredentialService struct {
	repo user_credential_repo.Repository
}

func NewUserCredentialService(repo user_credential_repo.Repository) *UserCredentialService {
	return &UserCredentialService{repo: repo}
}

func (s *UserCredentialService) GetUserCredentialById(ctx context.Context, id int32) (*user_credential.UserCredential, error) {
	return s.repo.GetById(ctx, id)
}

func (s *UserCredentialService) GetUserCredentialByUsername(ctx context.Context, username string) (*user_credential.UserCredential, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *UserCredentialService) GetAllUserCredentials(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*user_credential.UserCredential, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *UserCredentialService) CreateUserCredential(ctx context.Context, id int32, login, password string) (*user_credential.UserCredential, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	uc := &user_credential.UserCredential{
		CustomerId:   &id,
		Username:     login,
		PasswordHash: string(passwordHash),
	}

	return s.repo.Create(ctx, uc)
}

func (s *UserCredentialService) UpdateUserCredential(ctx context.Context, uc *user_credential.UserCredential) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(uc.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	uc.PasswordHash = string(passwordHash)
	return s.repo.Update(ctx, *uc)
}

func (s *UserCredentialService) DeleteUserCredential(ctx context.Context, customerId int32) error {
	return s.repo.Delete(ctx, customerId)
}
