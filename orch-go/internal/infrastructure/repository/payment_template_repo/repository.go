package payment_template_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/payment_template"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.PaymentTemplateServiceClient
}

func NewRepository(client pb.PaymentTemplateServiceClient) Repository {
	return Repository{client: client}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*payment_template.PaymentTemplate, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, err
	}

	var paymentTemplates []*payment_template.PaymentTemplate
	for _, pt := range resp.PaymentTemplates {
		paymentTemplates = append(paymentTemplates, ToDomain(pt))
	}

	return paymentTemplates, nil

}

func (r Repository) GetById(ctx context.Context, id int32) (*payment_template.PaymentTemplate, error) {
	resp, err := r.client.GetById(ctx, &pb.GetPaymentTemplateByIdRequest{TemplateId: id})
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}
}

func (r Repository) Add(ctx context.Context, paymentTemplate *payment_template.PaymentTemplate) (*payment_template.PaymentTemplate, error) {
	req := pb.AddPaymentTemplateRequest{
		CustomerId:    &paymentTemplate.CustomerID,
		Name:          &paymentTemplate.Name,
		TargetIban:    &paymentTemplate.TargetIBAN,
		DefaultAmount: &paymentTemplate.DefaultAmount,
	}
	resp, err := r.client.Add(ctx, &req)
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}

}

func (r Repository) Update(ctx context.Context, paymentTemplate *payment_template.PaymentTemplate) error {
	req := pb.UpdatePaymentTemplateRequest{
		TemplateId:    paymentTemplate.TemplateID,
		CustomerId:    &paymentTemplate.CustomerID,
		Name:          &paymentTemplate.Name,
		TargetIban:    &paymentTemplate.TargetIBAN,
		DefaultAmount: &paymentTemplate.DefaultAmount,
	}
	_, err := r.client.Update(ctx, &req)
	if err != nil {
		return fmt.Errorf("payment_template_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeletePaymentTemplateRequest{TemplateId: id})
	if err != nil {
		return fmt.Errorf("payment_template_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, paymentTemplates []*payment_template.PaymentTemplate) error {
	var models []*pb.AddPaymentTemplateRequest
	for _, pt := range paymentTemplates {
		models = append(models, &pb.AddPaymentTemplateRequest{
			CustomerId:    &pt.CustomerID,
			Name:          &pt.Name,
			TargetIban:    &pt.TargetIBAN,
			DefaultAmount: &pt.DefaultAmount,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddPaymentTemplateBulkRequest{Templates: models})
	if err != nil {
		return fmt.Errorf("payment_template_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, paymentTemplates []*payment_template.PaymentTemplate) error {
	var models []*pb.UpdatePaymentTemplateRequest
	for _, pt := range paymentTemplates {
		models = append(models, &pb.UpdatePaymentTemplateRequest{
			TemplateId:    pt.TemplateID,
			CustomerId:    &pt.CustomerID,
			Name:          &pt.Name,
			TargetIban:    &pt.TargetIBAN,
			DefaultAmount: &pt.DefaultAmount,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdatePaymentTemplateBulkRequest{Templates: models})
	if err != nil {
		return fmt.Errorf("payment_template_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeletePaymentTemplateRequest
	for _, id := range ids {
		models = append(models, &pb.DeletePaymentTemplateRequest{TemplateId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeletePaymentTemplateBulkRequest{Templates: models})
	if err != nil {
		return fmt.Errorf("payment_template_repo.DeleteBulk: %w", err)
	}
	return nil
}
