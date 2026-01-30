package services

type ServiceContainer struct {
	AccountService         *AccountService
	UserCredentialService  *UserCredentialService
	CustomerService        *CustomerService
	CustomerAddressService *CustomerAddressService
	AccountTypeService     *AccountTypeService
	AtmService             *AtmService
	BranchService          *BranchService
	CardService            *CardService
	DepositService         *DepositService
	LoanService            *LoanService
	ExchangeRateService    *ExchangeRateService
	FeeTypeService         *FeeTypeService
	TransactionService     *TransactionService
	PaymentTemplateService *PaymentTemplateService
	NotificationService    *NotificationService
	LoginLogService        *LoginLogService
}
