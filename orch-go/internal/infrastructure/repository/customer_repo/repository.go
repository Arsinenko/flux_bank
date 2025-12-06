package customer_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/customer"
)

type Repository struct {
	client pb.CustomerServiceClient
}

func NewRepository(client pb.CustomerServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32) ([]customer.Customer, error) {
	req := &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
	}
	resp, err := r.client.GetAll(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("customer_repo.GetAll: %w", err)
	}
	result := make([]customer.Customer, len(resp.Customers))
	for _, c := range resp.Customers {
		domainModel := ToDomain(c)
		if domainModel != nil {
			result = append(result, *domainModel)
		}
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*customer.Customer, error) {
	resp, err := r.client.GetById(ctx, &pb.GetCustomerByIdRequest{CustomerId: id})
	if err != nil {
		return nil, fmt.Errorf("customer_repo.GetById: %w", err)
	}
	domainModel := ToDomain(resp)
	if domainModel != nil {
		return domainModel, nil
	} else {
		return nil, fmt.Errorf("customer_repo.GetById: customer not found")
	}
}

func (r Repository) Create(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	req := &pb.AddCustomerRequest{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
		BirthDate: ToDateOnly(customer.BirthDate),
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("customer_repo.Create: %w", err)
	}
	domainModel := ToDomain(resp)
	if domainModel != nil {
		return domainModel, nil
	} else {
		return nil, fmt.Errorf("customer_repo.Create: customer not created")
	}
}

func (r Repository) Update(ctx context.Context, customer *customer.Customer) error {
	req := &pb.UpdateCustomerRequest{
		CustomerId: customer.Id,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Phone:      customer.Phone,
		BirthDate:  ToDateOnly(customer.BirthDate),
	}

	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("customer_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteCustomerRequest{CustomerId: id})
	if err != nil {
		return fmt.Errorf("customer_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, customers []*customer.Customer) error {
	var updates []*pb.UpdateCustomerRequest
	for _, c := range customers {
		updates = append(updates, &pb.UpdateCustomerRequest{
			CustomerId: c.Id,
			FirstName:  c.FirstName,
			LastName:   c.LastName,
			Email:      c.Email,
			Phone:      c.Phone,
			BirthDate:  ToDateOnly(c.BirthDate),
		},
		)
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateCustomerBulkRequest{Customers: updates})
	if err != nil {
		return fmt.Errorf("customer_repo.UpdateBulk: %w", err)
	} else {
		return nil
	}

}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var deletes []*pb.DeleteCustomerRequest
	for _, id := range ids {
		deletes = append(deletes, &pb.DeleteCustomerRequest{CustomerId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteCustomerBulkRequest{Customers: deletes})
	if err != nil {
		return fmt.Errorf("customer_repo.DeleteBulk: %w", err)
	}
	return nil
}
