package payment_template_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/payment_template"
)

func ToDomain(model *pb.PaymentTemplateModel) *payment_template.PaymentTemplate {
	if model == nil {
		return nil
	}
	return &payment_template.PaymentTemplate{
		TemplateID:    model.TemplateId,
		CustomerID:    *model.CustomerId,
		Name:          *model.Name,
		TargetIBAN:    *model.TargetIban,
		DefaultAmount: *model.DefaultAmount,
	}
}
