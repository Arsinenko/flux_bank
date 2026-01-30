package app

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/fee_type"
	"orch-go/internal/domain/transaction"
	"orch-go/internal/infrastructure/repository/account/account_repo"
	"orch-go/internal/infrastructure/repository/atm_repo"
	"orch-go/internal/infrastructure/repository/branch_repo"
	"orch-go/internal/infrastructure/repository/card_repo"
	"orch-go/internal/infrastructure/repository/customer_address_repo"
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

	"golang.org/x/sync/errgroup"
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
	customerAddressClient := pb.NewCustomerAddressServiceClient(conn)
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
	customerAddressRepo := customer_address_repo.NewCustomerAddressRepository(customerAddressClient)
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
		CustomerAddressService: services.NewCustomerAddressService(customerAddressRepo),
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

func InitAccountTypes(
	ctx context.Context,
	names []string,
	service *services.AccountTypeService,
) error {

	exists, err := service.GetAllAccountTypes(ctx, 0, 0, "", false)
	if err != nil {
		return fmt.Errorf("get account types: %w", err)
	}

	existsMap := make(map[string]struct{}, len(exists))
	for _, t := range exists {
		existsMap[t.Name] = struct{}{}
	}

	toCreate := make([]account.AccountType, 0)
	for _, name := range names {
		if _, ok := existsMap[name]; ok {
			continue
		}
		toCreate = append(toCreate, account.FakeAccountType(name))
	}

	if len(toCreate) == 0 {
		return nil
	}

	if err := service.CreateAccountTypeBulk(ctx, toCreate); err != nil {
		return fmt.Errorf("create account types: %w", err)
	}

	return nil
}

func InitFeeTypes(
	ctx context.Context,
	names []string,
	service *services.FeeTypeService,
) error {

	exists, err := service.GetAllFeeTypes(ctx, 0, 0, "", false)
	if err != nil {
		return fmt.Errorf("get fee types: %w", err)
	}

	existsMap := make(map[string]struct{}, len(exists))
	for _, t := range exists {
		if t.Name != nil {
			existsMap[*t.Name] = struct{}{}
		}
	}

	toCreate := make([]*fee_type.FeeType, 0)
	for _, name := range names {
		if _, ok := existsMap[name]; ok {
			continue
		}

		nameCopy := name
		toCreate = append(toCreate, &fee_type.FeeType{
			FeeID:       0,
			Name:        &nameCopy,
			Description: &nameCopy,
		})
	}

	if len(toCreate) == 0 {
		return nil
	}

	if err := service.CreateFeeTypeBulk(ctx, toCreate); err != nil {
		return fmt.Errorf("create fee types: %w", err)
	}

	return nil
}

func InitTransactionCategories(
	ctx context.Context,
	names []string,
	service *services.TransactionService,
) error {

	exists, err := service.GetAllTransactionCategories(ctx, 0, 0, "", false)
	if err != nil {
		return fmt.Errorf("get transaction categories: %w", err)
	}

	existsMap := make(map[string]struct{}, len(exists))
	for _, t := range exists {
		existsMap[t.Name] = struct{}{}
	}

	toCreate := make([]*transaction.TransactionCategory, 0)
	for _, name := range names {
		if _, ok := existsMap[name]; ok {
			continue
		}

		toCreate = append(toCreate, &transaction.TransactionCategory{
			CategoryID: 0,
			Name:       name,
		})
	}

	if len(toCreate) == 0 {
		return nil
	}

	if err := service.CreateTransactionCategoryBulk(ctx, toCreate); err != nil {
		return fmt.Errorf("create transaction categories: %w", err)
	}

	return nil
}

func InitAll(
	ctx context.Context,
	s *services.ServiceContainer,
	accountTypeNames,
	feeTypeNames,
	transactionCategoryNames []string,
) error {

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return InitAccountTypes(ctx, accountTypeNames, s.AccountTypeService)
	})

	g.Go(func() error {
		return InitFeeTypes(ctx, feeTypeNames, s.FeeTypeService)
	})

	g.Go(func() error {
		return InitTransactionCategories(ctx, transactionCategoryNames, s.TransactionService)
	})

	return g.Wait()
}
