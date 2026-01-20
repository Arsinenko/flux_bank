package account_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AccountTypeRepository struct {
	client pb.AccountTypeServiceClient
}

func NewAccountTypeRepository(client pb.AccountTypeServiceClient) AccountTypeRepository {
	return AccountTypeRepository{
		client: client,
	}
}

func (a AccountTypeRepository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.AccountType, error) {
	req := pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	}
	resp, err := a.client.GetAll(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetAll: %w", err)
	}
	result := make([]account.AccountType, 0, len(resp.AccountTypes))
	for _, c := range resp.AccountTypes {
		domainModel := AccountTypeToDomain(c)
		if domainModel != nil {
			result = append(result, *domainModel)
		}
	}
	return result, nil

}

func (a AccountTypeRepository) GetById(ctx context.Context, id int32) (*account.AccountType, error) {
	resp, err := a.client.GetById(ctx, &pb.GetAccountTypeByIdRequest{TypeId: id})
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetById: %w", err)
	}
	return AccountTypeToDomain(resp), nil
}

func (a AccountTypeRepository) Create(ctx context.Context, accountType *account.AccountType) (*account.AccountType, error) {
	req := pb.AddAccountTypeRequest{
		Name:        accountType.Name,
		Description: accountType.Description,
	}
	resp, err := a.client.Add(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("account_repo.Create: %w", err)
	}
	return AccountTypeToDomain(resp), nil
}

func (a AccountTypeRepository) Update(ctx context.Context, accountType *account.AccountType) error {
	req := pb.UpdateAccountTypeRequest{
		TypeId:      accountType.Id,
		Name:        accountType.Name,
		Description: accountType.Description,
	}
	_, err := a.client.Update(ctx, &req)
	if err != nil {
		return fmt.Errorf("account_repo.Update: %w", err)
	}
	return nil

}

func (a AccountTypeRepository) Delete(ctx context.Context, id int32) error {
	_, err := a.client.Delete(ctx, &pb.DeleteAccountTypeRequest{TypeId: id})
	if err != nil {
		return fmt.Errorf("account_repo.Delete: %w", err)
	}
	return nil

}

func (a AccountTypeRepository) AddBulk(ctx context.Context, accountTypes []account.AccountType) error {
	var models []*pb.AddAccountTypeRequest
	for _, at := range accountTypes {
		models = append(models, &pb.AddAccountTypeRequest{
			Name:        at.Name,
			Description: at.Description,
		})
	}
	_, err := a.client.AddBulk(ctx, &pb.AddAccountTypeBulkRequest{AccountTypes: models})
	if err != nil {
		return fmt.Errorf("account_repo.AddBulk: %w", err)
	}
	return nil
}

func (a AccountTypeRepository) UpdateBulk(ctx context.Context, accountTypes []account.AccountType) error {
	var models []*pb.UpdateAccountTypeRequest
	for _, at := range accountTypes {
		models = append(models, &pb.UpdateAccountTypeRequest{
			TypeId:      at.Id,
			Name:        at.Name,
			Description: at.Description,
		})
	}
	_, err := a.client.UpdateBulk(ctx, &pb.UpdateAccountTypeBulkRequest{AccountTypes: models})
	if err != nil {
		return fmt.Errorf("account_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (a AccountTypeRepository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteAccountTypeRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteAccountTypeRequest{TypeId: id})
	}
	_, err := a.client.DeleteBulk(ctx, &pb.DeleteAccountTypeBulkRequest{AccountTypes: models})
	if err != nil {
		return fmt.Errorf("account_repo.DeleteBulk: %w", err)
	} else {
		return nil
	}
}
