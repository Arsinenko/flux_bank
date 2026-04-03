package services

import (
	"context"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/atm"
	"orch-go/internal/domain/branch"
	"orch-go/internal/domain/card"
	"orch-go/internal/domain/customer"
	"orch-go/internal/domain/customer_address"
	"orch-go/internal/domain/deposit"
	"orch-go/internal/domain/exchange_rate"
	"orch-go/internal/domain/fee_type"
	"orch-go/internal/domain/loan"
	"orch-go/internal/domain/login_log"
	"orch-go/internal/domain/notification"
	"orch-go/internal/domain/payment_template"
	"orch-go/internal/domain/transaction"
	"orch-go/internal/domain/user_credential"
	"time"
)

type AccountServiceInt interface {
	GetAccountById(ctx context.Context, id int32) (*account.Account, error)
	GetAccountsByCustomer(ctx context.Context, customerId int32) ([]*account.Account, error)
	GetAccountsByDateRange(ctx context.Context, req account.GetByDateRange) ([]*account.Account, error)
	CreateAccount(ctx context.Context, acc *account.Account) (*account.Account, error)
	UpdateAccount(ctx context.Context, acc *account.Account) error
	DeleteAccount(ctx context.Context, id int32) error
	GetAllAccounts(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.Account, error)
	DeleteAccountBulk(ctx context.Context, ids []int32) error
	CreateAccountBulk(ctx context.Context, accs []*account.Account) error
	UpdateAccountBulk(ctx context.Context, accs []account.Account) error
}
type AccountTypeServiceInt interface {
	GetAccountTypeById(ctx context.Context, id int32) (*account.AccountType, error)
	GetAllAccountTypes(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]account.AccountType, error)
	CreateAccountType(ctx context.Context, at *account.AccountType) (*account.AccountType, error)
	UpdateAccountType(ctx context.Context, at *account.AccountType) error
	DeleteAccountType(ctx context.Context, id int32) error
	CreateAccountTypeBulk(ctx context.Context, ats []account.AccountType) error
	UpdateAccountTypeBulk(ctx context.Context, ats []account.AccountType) error
	DeleteAccountTypeBulk(ctx context.Context, ids []int32) error
}
type AtmServiceInt interface {
	GetAtmById(ctx context.Context, id int32) (*atm.Atm, error)
	GetAtmsByStatus(ctx context.Context, status string) ([]*atm.Atm, error)
	GetAtmsByLocationSubStr(ctx context.Context, subStr string) ([]*atm.Atm, error)
	GetAtmsByBranch(ctx context.Context, branchId int32) ([]*atm.Atm, error)
	GetAllAtms(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*atm.Atm, error)
	CreateAtm(ctx context.Context, atm *atm.Atm) (*atm.Atm, error)
	UpdateAtm(ctx context.Context, atm *atm.Atm) error
	DeleteAtm(ctx context.Context, id int32) error
	CreateAtmBulk(ctx context.Context, atms []*atm.Atm) error
	UpdateAtmBulk(ctx context.Context, atms []*atm.Atm) error
	DeleteAtmBulk(ctx context.Context, ids []int32) error
}
type BranchServiceInt interface {
	GetBranchById(ctx context.Context, id int32) (*branch.Branch, error)
	GetAllBranches(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*branch.Branch, error)
	CreateBranch(ctx context.Context, branch *branch.Branch) (*branch.Branch, error)
	UpdateBranch(ctx context.Context, branch *branch.Branch) error
	DeleteBranch(ctx context.Context, id int32) error
	CreateBranchBulk(ctx context.Context, branches []*branch.Branch) error
	UpdateBranchBulk(ctx context.Context, branches []*branch.Branch) error
	DeleteBranchBulk(ctx context.Context, ids []int32) error
}
type CardServiceInt interface {
	GetCardById(ctx context.Context, id int32) (*card.Card, error)
	GetCardsByAccountId(ctx context.Context, accountId int32) ([]*card.Card, error)
	GetAllCards(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*card.Card, error)
	CreateCard(ctx context.Context, card *card.Card) (*card.Card, error)
	UpdateCard(ctx context.Context, card *card.Card) error
	DeleteCard(ctx context.Context, id int32) error
	CreateCardBulk(ctx context.Context, cards []*card.Card) error
	UpdateCardBulk(ctx context.Context, cards []*card.Card) error
	DeleteCardBulk(ctx context.Context, ids []int32) error
}
type CustomerAddressServiceInt interface {
	CreateCustomerAddress(ctx context.Context, address *customer_address.CustomerAddress) (*customer_address.CustomerAddress, error)
	UpdateCustomerAddress(ctx context.Context, address *customer_address.CustomerAddress) error
	DeleteCustomerAddress(ctx context.Context, id int32) error
	GetCustomerAddressById(ctx context.Context, id int32) (*customer_address.CustomerAddress, error)
	GetAllCustomerAddresses(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]customer_address.CustomerAddress, error)
	GetCustomerAddressesByCustomerId(ctx context.Context, customerId int32) ([]customer_address.CustomerAddress, error)
}
type CustomerServiceInt interface {
	GetCustomerById(ctx context.Context, id int32) (*customer.Customer, error)
	GetCustomersBySubstring(ctx context.Context, req customer.GetBySubStrRequest) ([]customer.Customer, error)
	GetCustomersByDateRange(ctx context.Context, req customer.GetByDateRangeRequest) ([]customer.Customer, error)
	GetAllCustomers(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]customer.Customer, error)
	CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error)
	UpdateCustomer(ctx context.Context, customer *customer.Customer) error
	DeleteCustomer(ctx context.Context, id int32) error
	CreateCustomerBulk(ctx context.Context, customers []*customer.Customer) error
	UpdateCustomerBulk(ctx context.Context, customers []*customer.Customer) error
	DeleteCustomerBulk(ctx context.Context, ids []int32) error
}
type DepositServiceInt interface {
	GetDepositById(ctx context.Context, id int32) (*deposit.Deposit, error)
	GetDepositsByCustomer(ctx context.Context, customerId int32) ([]*deposit.Deposit, error)
	GetAllDeposits(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*deposit.Deposit, error)
	CreateDeposit(ctx context.Context, deposit *deposit.Deposit) (*deposit.Deposit, error)
	UpdateDeposit(ctx context.Context, deposit *deposit.Deposit) error
	DeleteDeposit(ctx context.Context, id int32) error
	CreateDepositBulk(ctx context.Context, deposits []*deposit.Deposit) error
	UpdateDepositBulk(ctx context.Context, deposits []*deposit.Deposit) error
	DeleteDepositBulk(ctx context.Context, ids []int32) error
}
type ExchangeRateServiceInt interface {
	GetExchangeRateById(ctx context.Context, id int32) (*exchange_rate.ExchangeRate, error)
	GetExchangeRatesByBaseCurrency(ctx context.Context, baseCurrency string) ([]*exchange_rate.ExchangeRate, error)
	GetAllExchangeRates(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*exchange_rate.ExchangeRate, error)
	CreateExchangeRate(ctx context.Context, er *exchange_rate.ExchangeRate) (*exchange_rate.ExchangeRate, error)
	UpdateExchangeRate(ctx context.Context, er *exchange_rate.ExchangeRate) error
	DeleteExchangeRate(ctx context.Context, id int32) error
	CreateExchangeRateBulk(ctx context.Context, ers []*exchange_rate.ExchangeRate) error
	UpdateExchangeRateBulk(ctx context.Context, ers []*exchange_rate.ExchangeRate) error
	DeleteExchangeRateBulk(ctx context.Context, ids []int32) error
}
type FeeTypeServiceInt interface {
	GetFeeTypeById(ctx context.Context, id int32) (*fee_type.FeeType, error)
	GetAllFeeTypes(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*fee_type.FeeType, error)
	CreateFeeType(ctx context.Context, ft *fee_type.FeeType) (*fee_type.FeeType, error)
	UpdateFeeType(ctx context.Context, ft *fee_type.FeeType) error
	DeleteFeeType(ctx context.Context, id int32) error
	CreateFeeTypeBulk(ctx context.Context, fts []*fee_type.FeeType) error
	UpdateFeeTypeBulk(ctx context.Context, fts []*fee_type.FeeType) error
	DeleteFeeTypeBulk(ctx context.Context, ids []int32) error
}
type LoanServiceInt interface {
	GetLoanById(ctx context.Context, id int32) (*loan.Loan, error)
	GetLoansByCustomer(ctx context.Context, customerId int32) ([]*loan.Loan, error)
	GetAllLoans(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*loan.Loan, error)
	CreateLoan(ctx context.Context, l *loan.Loan) (*loan.Loan, error)
	UpdateLoan(ctx context.Context, l *loan.Loan) error
	DeleteLoan(ctx context.Context, id int32) error
	DeleteLoanBulk(ctx context.Context, ids []int32) error
	GetLoanPaymentById(ctx context.Context, id int32) (*loan.LoanPayment, error)
	GetLoanPaymentsByLoan(ctx context.Context, loanId int32) ([]*loan.LoanPayment, error)
	GetAllLoanPayments(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*loan.LoanPayment, error)
	CreateLoanPayment(ctx context.Context, lp *loan.LoanPayment) (*loan.LoanPayment, error)
	UpdateLoanPayment(ctx context.Context, lp *loan.LoanPayment) error
	DeleteLoanPayment(ctx context.Context, id int32) error
	DeleteLoanPaymentBulk(ctx context.Context, ids []int32) error
}
type LoginLogServiceInt interface {
	GetLoginLogById(ctx context.Context, id int32) (*login_log.LoginLog, error)
	GetLoginLogsByCustomer(ctx context.Context, customerId int32) ([]*login_log.LoginLog, error)
	GetLoginLogsInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*login_log.LoginLog, error)
	GetAllLoginLogs(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*login_log.LoginLog, error)
	CreateLoginLog(ctx context.Context, log *login_log.LoginLog) (*login_log.LoginLog, error)
	UpdateLoginLog(ctx context.Context, log *login_log.LoginLog) error
	DeleteLoginLog(ctx context.Context, id int32) error
}
type NotificationServiceInt interface {
	GetNotificationById(ctx context.Context, id int32) (*notification.Notification, error)
	GetNotificationsByCustomer(ctx context.Context, customerId int32, isRead bool) ([]*notification.Notification, error)
	GetNotificationsByDateRange(ctx context.Context, req notification.GetByDateRangeRequest) ([]*notification.Notification, error)
	GetAllNotifications(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*notification.Notification, error)
	CreateNotification(ctx context.Context, notif *notification.Notification) (*notification.Notification, error)
	UpdateNotification(ctx context.Context, notif *notification.Notification) error
	DeleteNotification(ctx context.Context, id int32) error
	CreateNotificationBulk(ctx context.Context, notifs []*notification.Notification) error
	UpdateNotificationBulk(ctx context.Context, notifs []*notification.Notification) error
	DeleteNotificationBulk(ctx context.Context, ids []int32) error
}
type PaymentTemplateServiceInt interface {
	GetPaymentTemplateById(ctx context.Context, id int32) (*payment_template.PaymentTemplate, error)
	GetAllPaymentTemplates(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*payment_template.PaymentTemplate, error)
	CreatePaymentTemplate(ctx context.Context, pt *payment_template.PaymentTemplate) (*payment_template.PaymentTemplate, error)
	UpdatePaymentTemplate(ctx context.Context, pt *payment_template.PaymentTemplate) error
	DeletePaymentTemplate(ctx context.Context, id int32) error
	CreatePaymentTemplateBulk(ctx context.Context, pts []*payment_template.PaymentTemplate) error
	UpdatePaymentTemplateBulk(ctx context.Context, pts []*payment_template.PaymentTemplate) error
	DeletePaymentTemplateBulk(ctx context.Context, ids []int32) error
}
type TransactionServiceInt interface {
	GetTransactionRevenue(ctx context.Context, accountId int32, req transaction.GetByDateRange) ([]*transaction.Transaction, error)
	GetAllTransactions(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.Transaction, error)
	GetTransactionById(ctx context.Context, id int32) (*transaction.Transaction, error)
	CreateTransaction(ctx context.Context, t *transaction.Transaction) (*transaction.Transaction, error)
	UpdateTransaction(ctx context.Context, t *transaction.Transaction) error
	DeleteTransaction(ctx context.Context, id int32) error
	CreateTransactionBulk(ctx context.Context, ts []*transaction.Transaction) error
	UpdateTransactionBulk(ctx context.Context, ts []*transaction.Transaction) error
	DeleteTransactionBulk(ctx context.Context, ids []int32) error
	GetAllTransactionCategories(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionCategory, error)
	GetTransactionCategoryById(ctx context.Context, id int32) (*transaction.TransactionCategory, error)
	CreateTransactionCategory(ctx context.Context, tc *transaction.TransactionCategory) (*transaction.TransactionCategory, error)
	UpdateTransactionCategory(ctx context.Context, tc *transaction.TransactionCategory) error
	DeleteTransactionCategory(ctx context.Context, id int32) error
	CreateTransactionCategoryBulk(ctx context.Context, tcs []*transaction.TransactionCategory) error
	UpdateTransactionCategoryBulk(ctx context.Context, tcs []*transaction.TransactionCategory) error
	DeleteTransactionCategoryBulk(ctx context.Context, ids []int32) error
	GetAllTransactionFees(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*transaction.TransactionFee, error)
	GetTransactionFeeById(ctx context.Context, id int32) (*transaction.TransactionFee, error)
	CreateTransactionFee(ctx context.Context, tf *transaction.TransactionFee) (*transaction.TransactionFee, error)
	UpdateTransactionFee(ctx context.Context, tf *transaction.TransactionFee) error
	DeleteTransactionFee(ctx context.Context, id int32) error
	CreateTransactionFeeBulk(ctx context.Context, tfs []*transaction.TransactionFee) error
	UpdateTransactionFeeBulk(ctx context.Context, tfs []*transaction.TransactionFee) error
	DeleteTransactionFeeBulk(ctx context.Context, ids []int32) error
}
type UserCredentialServiceInt interface {
	GetUserCredentialById(ctx context.Context, id int32) (*user_credential.UserCredential, error)
	GetUserCredentialByUsername(ctx context.Context, username string) (*user_credential.UserCredential, error)
	GetAllUserCredentials(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*user_credential.UserCredential, error)
	CreateUserCredential(ctx context.Context, id int32, login, password string) (*user_credential.UserCredential, error)
	UpdateUserCredential(ctx context.Context, uc *user_credential.UserCredential) error
	DeleteUserCredential(ctx context.Context, customerId int32) error
}
