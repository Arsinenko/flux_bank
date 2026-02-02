package services

import (
	"context"
	"orch-go/internal/domain/customer_address"
	"orch-go/internal/infrastructure/repository/customer_address_repo"
)

type CustomerAddressService struct {
	repo *customer_address_repo.Repository
}

func NewCustomerAddressService(repo *customer_address_repo.Repository) *CustomerAddressService {
	return &CustomerAddressService{repo: repo}
}

func (s *CustomerAddressService) CreateCustomerAddress(ctx context.Context, address *customer_address.CustomerAddress) (*customer_address.CustomerAddress, error) {
	return s.repo.Create(ctx, address)
}

func (s *CustomerAddressService) UpdateCustomerAddress(ctx context.Context, address *customer_address.CustomerAddress) error {
	return s.repo.Update(ctx, address)
}

func (s *CustomerAddressService) DeleteCustomerAddress(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *CustomerAddressService) GetCustomerAddressById(ctx context.Context, id int32) (*customer_address.CustomerAddress, error) {
	return s.repo.GetById(ctx, id)
}

func (s *CustomerAddressService) GetAllCustomerAddresses(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]customer_address.CustomerAddress, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *CustomerAddressService) GetCustomerAddressesByCustomerId(ctx context.Context, customerId int32) ([]customer_address.CustomerAddress, error) {
	return s.repo.GetByCustomerId(ctx, customerId)
}
