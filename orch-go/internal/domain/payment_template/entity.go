package payment_template

type PaymentTemplate struct {
	TemplateID    int32
	CustomerID    int32
	Name          string
	TargetIBAN    string
	DefaultAmount string
}
