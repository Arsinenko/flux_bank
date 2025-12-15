package services

import (
	"context"
	"orch-go/internal/domain/customer"
	"orch-go/internal/infrastructure/repository/customer_repo"
)

type CustomerService struct {
	repo *customer_repo.Repository
}

func NewCustomerService(repo *customer_repo.Repository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetCustomerById(ctx context.Context, id int32) (*customer.Customer, error) {
	return s.repo.GetById(ctx, id)
}

func (s *CustomerService) GetCustomersBySubstring(ctx context.Context, req customer.GetBySubStrRequest) ([]customer.Customer, error) {
	return s.repo.GetBySubstring(ctx, req)
}

func (s *CustomerService) GetCustomersByDateRange(ctx context.Context, req customer.GetByDateRangeRequest) ([]customer.Customer, error) {
	return s.repo.GetByDateRange(ctx, req)
}

func (s *CustomerService) GetAllCustomers(ctx context.Context, pageN, pageSize int32) ([]customer.Customer, error) {
	return s.repo.GetAll(ctx, pageN, pageSize)
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	return s.repo.Create(ctx, customer)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer *customer.Customer) error {
	return s.repo.Update(ctx, customer)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *CustomerService) UpdateCustomerBulk(ctx context.Context, customers []*customer.Customer) error {
	return s.repo.UpdateBulk(ctx, customers)
}

func (s *CustomerService) DeleteCustomerBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
