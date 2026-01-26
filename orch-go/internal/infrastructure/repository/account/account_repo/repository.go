package account_repo

import (
	context "context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.AccountServiceClient
}

func (r Repository) GetByCustomerId(ctx context.Context, customerId int32) ([]*account.Account, error) {
	resp, err := r.client.GetByCustomerId(ctx, &pb.GetAccountByCustomerIdRequest{CustomerId: customerId})
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetByCustomerId: %w", err)
	}
	result := make([]*account.Account, 0, len(resp.Accounts))
	for _, a := range resp.Accounts {
		result = append(result, AccountToDomain(a))
	}
	return result, nil
}

func (r Repository) GetByDateRange(ctx context.Context, request account.GetByDateRange) ([]*account.Account, error) {
	resp, err := r.client.GetByDateRange(ctx, &pb.GetByDateRangeRequest{
		FromDate: timestamppb.New(request.From),
		ToDate:   timestamppb.New(request.To),
		PageN:    &request.PageN,
		PageSize: &request.PageSize,
	})
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetByDateRange: %w", err)
	}

	result := make([]*account.Account, 0, len(resp.Accounts))
	for _, a := range resp.Accounts {
		result = append(result, AccountToDomain(a))
	}
	return result, nil
}

func NewRepository(client pb.AccountServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.Account, error) {
	req := pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	}
	resp, err := r.client.GetAll(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetAll: %w", err)
	}
	result := make([]account.Account, 0, len(resp.Accounts))
	for _, c := range resp.Accounts {
		domainModel := AccountToDomain(c)
		if domainModel != nil {
			result = append(result, *domainModel)
		}
	}
	return result, nil

}

func (r Repository) GetById(ctx context.Context, id int32) (*account.Account, error) {
	resp, err := r.client.GetById(ctx, &pb.GetAccountByIdRequest{AccountId: id})
	if err != nil {
		return nil, fmt.Errorf("account_repo.GetById: %w", err)

	}
	return AccountToDomain(resp), nil
}

func (r Repository) Create(ctx context.Context, account *account.Account) (*account.Account, error) {
	req := pb.AddAccountRequest{
		CustomerId: &account.CustomerId,
		TypeId:     &account.TypeId,
		Iban:       account.Iban,
		Balance:    &account.Balance,
		IsActive:   &account.IsActive,
	}
	resp, err := r.client.Add(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("account_repo.Create: %w", err)
	}
	return AccountToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, account *account.Account) error {
	req := pb.UpdateAccountRequest{
		AccountId:  *account.Id,
		CustomerId: &account.CustomerId,
		TypeId:     &account.TypeId,
		Iban:       account.Iban,
		Balance:    &account.Balance,
		IsActive:   &account.IsActive,
	}
	_, err := r.client.Update(ctx, &req)
	if err != nil {
		return fmt.Errorf("account_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteAccountRequest{AccountId: id})
	if err != nil {
		return fmt.Errorf("account_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, accounts []account.Account) error {
	var models []*pb.UpdateAccountRequest
	for _, a := range accounts {
		models = append(models, &pb.UpdateAccountRequest{
			AccountId:  *a.Id,
			CustomerId: &a.CustomerId,
			TypeId:     &a.TypeId,
			Iban:       a.Iban,
			Balance:    &a.Balance,
			IsActive:   &a.IsActive,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateAccountBulkRequest{Accounts: models})
	if err != nil {
		return fmt.Errorf("account_repo.UpdateBulk: %w", err)
	} else {
		return nil
	}
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteAccountRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteAccountRequest{AccountId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteAccountBulkRequest{Accounts: models})
	if err != nil {
		return fmt.Errorf("account_repo.DeleteBulk: %w", err)
	}
	return nil
}
