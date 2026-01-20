package user_credential_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/user_credential"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.UserCredentialServiceClient
}

func (r Repository) GetByUsername(ctx context.Context, username string) (*user_credential.UserCredential, error) {
	resp, err := r.client.GetByUsername(ctx, &pb.GetUserCredentialByUsernameRequest{Username: username})
	if err != nil {
		return nil, fmt.Errorf("user_credential_repo.GetByUsername: %w", err)
	}
	return ToDomain(resp), nil
}

func NewRepository(client pb.UserCredentialServiceClient) Repository {
	return Repository{client: client}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*user_credential.UserCredential, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, fmt.Errorf("user_credential_repo.GetAll: %w", err)
	}
	var userCredentials []*user_credential.UserCredential
	for _, uc := range resp.UserCredentials {
		userCredentials = append(userCredentials, ToDomain(uc))
	}
	return userCredentials, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*user_credential.UserCredential, error) {
	resp, err := r.client.GetById(ctx, &pb.GetUserCredentialByIdRequest{CustomerId: id})
	if err != nil {
		return nil, fmt.Errorf("user_credential_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Create(ctx context.Context, cred *user_credential.UserCredential) (*user_credential.UserCredential, error) {
	req := &pb.AddUserCredentialRequest{
		CustomerId:   cred.CustomerId,
		Username:     cred.Username,
		PasswordHash: cred.PasswordHash,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("user_credential_repo.Create: %w", err)
	}
	return ToDomain(resp), nil

}

func (r Repository) Update(ctx context.Context, cred user_credential.UserCredential) error {
	req := &pb.UpdateUserCredentialRequest{
		CustomerId:   cred.CustomerId,
		Username:     cred.Username,
		PasswordHash: cred.PasswordHash,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("user_credential_repo.Update: %w", err)
	}
	return nil

}

func (r Repository) Delete(ctx context.Context, customerId int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteUserCredentialRequest{CustomerId: customerId})
	if err != nil {
		return fmt.Errorf("user_credential_repo.Delete: %w", err)
	}
	return nil

}
