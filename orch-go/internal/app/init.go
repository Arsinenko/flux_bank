package app

import (
	pb "orch-go/api/generated"
	"orch-go/internal/infrastructure/repository/account/account_repo"
	"orch-go/internal/infrastructure/repository/atm_repo"
	"orch-go/internal/infrastructure/repository/branch_repo"
	"orch-go/internal/infrastructure/repository/card_repo"
	"orch-go/internal/infrastructure/repository/customer_repo"
	"orch-go/internal/infrastructure/repository/deposit_repo"
	"orch-go/internal/infrastructure/repository/exchange_rate_repo"
	"orch-go/internal/infrastructure/repository/fee_type_repo"
	"orch-go/internal/infrastructure/repository/loan_repo"
	"orch-go/internal/infrastructure/repository/login_log_repo"
	"orch-go/internal/infrastructure/repository/notification_repo"
	"orch-go/internal/infrastructure/repository/payment_template_repo"
	"orch-go/internal/infrastructure/repository/transaction_repo"
	"orch-go/internal/infrastructure/repository/user_credential_repo"
	"orch-go/internal/services"

	"google.golang.org/grpc"
)

func InitServices(conn *grpc.ClientConn) *services.ServiceContainer {
	//Init gRPC clients
	accountClient := pb.NewAccountServiceClient(conn)
	accountTypeClient := pb.NewAccountTypeServiceClient(conn)
	atmClient := pb.NewAtmServiceClient(conn)
	branchClient := pb.NewBranchServiceClient(conn)
	cardClient := pb.NewCardServiceClient(conn)
	customerClient := pb.NewCustomerServiceClient(conn)
	depositClient := pb.NewDepositServiceClient(conn)
	exchangeRateClient := pb.NewExchangeRateServiceClient(conn)
	feeTypeClient := pb.NewFeeTypeServiceClient(conn)
	loanClient := pb.NewLoanServiceClient(conn)
	loanPaymentClient := pb.NewLoanPaymentServiceClient(conn)
	loginLogClient := pb.NewLoginLogServiceClient(conn)
	notificationClient := pb.NewNotificationServiceClient(conn)
	paymentTemplateClient := pb.NewPaymentTemplateServiceClient(conn)
	transactionClient := pb.NewTransactionServiceClient(conn)
	transactionCategoryClient := pb.NewTransactionCategoryServiceClient(conn)
	transactionFeeClient := pb.NewTransactionFeeServiceClient(conn)
	userCredentialClient := pb.NewUserCredentialServiceClient(conn)

	// Init infrastructure repositories
	accountRepo := account_repo.NewRepository(accountClient)
	accountTypeRepo := account_repo.NewAccountTypeRepository(accountTypeClient)
	atmRepo := atm_repo.NewRepository(atmClient)
	branchRepo := branch_repo.NewRepository(branchClient)
	cardRepo := card_repo.NewRepository(cardClient)
	customerRepo := customer_repo.NewRepository(customerClient)
	depositRepo := deposit_repo.NewRepository(depositClient)
	exchangeRateRepo := exchange_rate_repo.NewRepository(exchangeRateClient)
	feeTypeRepo := fee_type_repo.NewRepository(feeTypeClient)
	loanRepo := loan_repo.NewLoanRepository(loanClient)
	loanPaymentRepo := loan_repo.NewLoanPaymentRepository(loanPaymentClient)
	loginLogRepo := login_log_repo.NewRepository(loginLogClient)
	notificationRepo := notification_repo.NewRepository(notificationClient)
	paymentTemplateRepo := payment_template_repo.NewRepository(paymentTemplateClient)
	transactionRepo := transaction_repo.NewTransactionRepository(transactionClient)
	transactionCategoryRepo := transaction_repo.NewTransactionCategoryRepository(transactionCategoryClient)
	transactionFeeRepo := transaction_repo.NewTransactionFeeRepository(transactionFeeClient)
	userCredentialRepo := user_credential_repo.NewRepository(userCredentialClient)

	// Init services
	return &services.ServiceContainer{
		AccountService:         services.NewAccountService(accountRepo),
		AccountTypeService:     services.NewAccountTypeService(accountTypeRepo),
		AtmService:             services.NewAtmService(atmRepo),
		BranchService:          services.NewBranchService(branchRepo),
		CardService:            services.NewCardService(cardRepo),
		CustomerService:        services.NewCustomerService(customerRepo),
		DepositService:         services.NewDepositService(depositRepo),
		ExchangeRateService:    services.NewExchangeRateService(exchangeRateRepo),
		FeeTypeService:         services.NewFeeTypeService(feeTypeRepo),
		LoanService:            services.NewLoanService(loanRepo, loanPaymentRepo),
		LoginLogService:        services.NewLoginLogService(loginLogRepo),
		NotificationService:    services.NewNotificationService(notificationRepo),
		PaymentTemplateService: services.NewPaymentTemplateService(paymentTemplateRepo),
		TransactionService:     services.NewTransactionService(transactionRepo, transactionCategoryRepo, transactionFeeRepo),
		UserCredentialService:  services.NewUserCredentialService(userCredentialRepo),
	}
}
