package services

type ServiceContainer struct {
	Account         *AccountService
	UserCredential  *UserCredentialService
	Customer        *CustomerService
	AccountType     *AccountTypeService
	Atm             *AtmService
	Branch          *BranchService
	Card            *CardService
	Deposit         *DepositService
	Loan            *LoanService
	LoanPayment     *LoanService
	ExchangeRate    *ExchangeRateService
	FeeType         *FeeTypeService
	Transaction     *TransactionService
	PaymentTemplate *PaymentTemplateService
	Notification    *NotificationService
	LoginLog        *LoginLogService
}
