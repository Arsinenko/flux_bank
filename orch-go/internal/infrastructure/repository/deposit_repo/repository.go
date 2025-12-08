package deposit_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/deposit"
)

type Repository struct {
	client pb.DepositServiceClient
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32) ([]*deposit.Deposit, error) {
	req := pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
	}
	resp, err := r.client.GetAll(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("deposit_repo.GetAll: %w", err)
	}
	result := ToDepositsDomain(resp)
	return result, nil

}

func (r Repository) GetById(ctx context.Context, id int32) (*deposit.Deposit, error) {
	resp, err := r.client.GetById(ctx, &pb.GetDepositByIdRequest{DepositId: id})
	if err != nil {
		return nil, fmt.Errorf("deposit_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil

}

func (r Repository) GetByCustomer(ctx context.Context, customerId int32) ([]*deposit.Deposit, error) {
	resp, err := r.client.GetByCustomer(ctx, &pb.GetDepositsByCustomerRequest{CustomerId: customerId})
	if err != nil {
		return nil, fmt.Errorf("deposit_repo.GetByCustomer: %w", err)
	}
	result := ToDepositsDomain(resp)
	return result, nil

}

func (r Repository) Add(ctx context.Context, deposit *deposit.Deposit) (*deposit.Deposit, error) {
	req := &pb.AddDepositRequest{
		CustomerId:   &deposit.CustomerID,
		Amount:       &deposit.Amount,
		InterestRate: &deposit.InterestRate,
		StartDate:    ToDateOnly(&deposit.StartDate),
		EndDate:      ToDateOnly(&deposit.EndDate),
		Status:       &deposit.Status,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("deposit_repo.Add: %w", err)
	}
	return ToDomain(resp), nil

}

func (r Repository) Update(ctx context.Context, deposit *deposit.Deposit) error {
	req := &pb.UpdateDepositRequest{
		DepositId:    deposit.DepositID,
		CustomerId:   &deposit.CustomerID,
		Amount:       &deposit.Amount,
		InterestRate: &deposit.InterestRate,
		StartDate:    ToDateOnly(&deposit.StartDate),
		EndDate:      ToDateOnly(&deposit.EndDate),
		Status:       &deposit.Status,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("deposit_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteDepositRequest{DepositId: id})
	if err != nil {
		return fmt.Errorf("deposit_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, deposits []*deposit.Deposit) error {
	var models []*pb.AddDepositRequest
	for _, d := range deposits {
		models = append(models, &pb.AddDepositRequest{
			CustomerId:   &d.CustomerID,
			Amount:       &d.Amount,
			InterestRate: &d.InterestRate,
			StartDate:    ToDateOnly(&d.StartDate),
			EndDate:      ToDateOnly(&d.EndDate),
			Status:       &d.Status,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddDepositBulkRequest{Deposits: models})
	if err != nil {
		return fmt.Errorf("deposit_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, deposits []*deposit.Deposit) error {
	var models []*pb.UpdateDepositRequest
	for _, d := range deposits {
		models = append(models, &pb.UpdateDepositRequest{
			DepositId:    d.DepositID,
			CustomerId:   &d.CustomerID,
			Amount:       &d.Amount,
			InterestRate: &d.InterestRate,
			StartDate:    ToDateOnly(&d.StartDate),
			EndDate:      ToDateOnly(&d.EndDate),
			Status:       &d.Status,
		})

	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateDepositBulkRequest{Deposits: models})
	if err != nil {
		return fmt.Errorf("deposit_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteDepositRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteDepositRequest{DepositId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteDepositBulkRequest{Deposits: models})
	if err != nil {
		return fmt.Errorf("deposit_repo.DeleteBulk: %w", err)
	}
	return nil

}

func NewRepository(client pb.DepositServiceClient) Repository {
	return Repository{client: client}
}
