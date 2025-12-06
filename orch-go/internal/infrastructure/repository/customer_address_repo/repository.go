package customer_address_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	customerAdress "orch-go/internal/domain/customerAddress"
)

type Repository struct {
	client pb.CustomerAddressServiceClient
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32) ([]customerAdress.CustomerAddress, error) {
	req := pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
	}
	resp, err := r.client.GetAll(ctx, &req)
	if err != nil {
		return nil, err
	}
	result := make([]customerAdress.CustomerAddress, len(resp.CustomerAddresses))
	for _, c := range resp.CustomerAddresses {
		domainModel := ToDomain(c)
		if domainModel != nil {
			result = append(result, *domainModel)
		}
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*customerAdress.CustomerAddress, error) {
	resp, err := r.client.GetById(ctx, &pb.GetCustomerAddressByIdRequest{AddressId: id})
	if err != nil {
		return nil, fmt.Errorf("customer_address_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Create(ctx context.Context, address *customerAdress.CustomerAddress) (*customerAdress.CustomerAddress, error) {
	req := &pb.AddCustomerAddressRequest{
		CustomerId: &address.Customer_id,
		Country:    &address.Country,
		City:       &address.City,
		Street:     &address.Street,
		ZipCode:    &address.ZipCode,
		IsPrimary:  &address.IsPrimary,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("customer_address_repo.Create: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, address *customerAdress.CustomerAddress) error {
	req := &pb.UpdateCustomerAddressRequest{
		AddressId:  *address.Id,
		CustomerId: &address.Customer_id,
		Country:    &address.Country,
		City:       &address.City,
		Street:     &address.Street,
		ZipCode:    &address.ZipCode,
		IsPrimary:  &address.IsPrimary,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("customer_address_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteCustomerAddressRequest{AddressId: id})
	if err != nil {
		return fmt.Errorf("customer_address_repo.Delete: %w", err)
	}
	return nil
}

func New(client pb.CustomerAddressServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}
