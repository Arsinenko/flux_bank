package services

import (
	"context"
	"orch-go/internal/domain/payment_template"
	"orch-go/internal/infrastructure/repository/payment_template_repo"
)

type PaymentTemplateService struct {
	repo payment_template_repo.Repository
}

func NewPaymentTemplateService(repo payment_template_repo.Repository) *PaymentTemplateService {
	return &PaymentTemplateService{repo: repo}
}

func (s *PaymentTemplateService) GetPaymentTemplateById(ctx context.Context, id int32) (*payment_template.PaymentTemplate, error) {
	return s.repo.GetById(ctx, id)
}

func (s *PaymentTemplateService) GetAllPaymentTemplates(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*payment_template.PaymentTemplate, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *PaymentTemplateService) CreatePaymentTemplate(ctx context.Context, pt *payment_template.PaymentTemplate) (*payment_template.PaymentTemplate, error) {
	return s.repo.Add(ctx, pt)
}

func (s *PaymentTemplateService) UpdatePaymentTemplate(ctx context.Context, pt *payment_template.PaymentTemplate) error {
	return s.repo.Update(ctx, pt)
}

func (s *PaymentTemplateService) DeletePaymentTemplate(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *PaymentTemplateService) CreatePaymentTemplateBulk(ctx context.Context, pts []*payment_template.PaymentTemplate) error {
	return s.repo.AddBulk(ctx, pts)
}

func (s *PaymentTemplateService) UpdatePaymentTemplateBulk(ctx context.Context, pts []*payment_template.PaymentTemplate) error {
	return s.repo.UpdateBulk(ctx, pts)
}

func (s *PaymentTemplateService) DeletePaymentTemplateBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
