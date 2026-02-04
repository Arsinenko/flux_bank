package app

import (
	"context"
	"fmt"
	"math/rand"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/atm"
	"orch-go/internal/domain/branch"
	"orch-go/internal/domain/card"
	"orch-go/internal/domain/customer"
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
	"time"

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
	//CreateTestCustomers(ctx, *s.CustomerService)
	CreateTestAccounts(ctx, s)
	CreateTestCards(ctx, s)
	CreateTestBranches(ctx, s)
	CreateTestAtms(ctx, s)

	return g.Wait()
}

func CreateTestCustomers(ctx context.Context, service services.CustomerService) {
	var customers []*customer.Customer
	for i := 0; i < 100; i++ {
		customers = append(customers, customer.FakeCustomer(time.Now()))
	}
	err := service.CreateCustomerBulk(ctx, customers)
	if err != nil {
		fmt.Printf("create customers: %w", err.Error())
		return
	}
}

func CreateTestAccounts(ctx context.Context, container *services.ServiceContainer) {
	customers, err := container.CustomerService.GetAllCustomers(ctx, 0, 0, "", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	accTypes, err := container.AccountTypeService.GetAllAccountTypes(ctx, 0, 0, "", false)
	if err != nil {
		fmt.Println("get account types: ", err.Error())
		return
	}

	var accTypeIds []int32
	for i := 0; i < len(accTypes); i++ {
		accTypeIds = append(accTypeIds, *accTypes[i].Id)
	}

	var accounts []*account.Account
	for _, c := range customers {
		a := account.FakeAccount(c.Id, accTypeIds[rand.Intn(len(accTypes))])
		accounts = append(accounts, a)
	}
	err = container.AccountService.CreateAccountBulk(ctx, accounts)
	if err != nil {
		fmt.Println("create accounts: %w", err.Error())
		return
	}

}

func CreateTestCards(ctx context.Context, container *services.ServiceContainer) {
	accounts, err := container.AccountService.GetAllAccounts(ctx, 0, 0, "", false)
	if err != nil {
		fmt.Printf("get accounts: %w", err.Error())
		return
	}

	var accIds []int32
	for i := 0; i < len(accounts); i++ {
		accIds = append(accIds, *accounts[i].Id)
	}

	var cards []*card.Card
	for _, a := range accIds {
		cards = append(cards, card.FakeCard(0, a, time.Now(), "active"))
	}
	err = container.CardService.CreateCardBulk(ctx, cards)
	if err != nil {
		fmt.Printf("create cards: %w", err.Error())
		return
	}

}

func CreateTestBranches(ctx context.Context, container *services.ServiceContainer) {
	var branches []*branch.Branch
	for i := 0; i < 20; i++ {
		b := branch.FakeBranch()
		branches = append(branches, &b)
	}
	err := container.BranchService.CreateBranchBulk(ctx, branches)
	if err != nil {
		fmt.Printf("create branches: %w", err.Error())
		return
	}

}

func CreateTestAtms(ctx context.Context, container *services.ServiceContainer) {
	branches, err := container.BranchService.GetAllBranches(ctx, 0, 0, "", false)
	if err != nil {
		fmt.Printf("get branches: %w", err.Error())
		return
	}

	var branchIds []int32
	for _, b := range branches {
		branchIds = append(branchIds, *b.BranchID)

	}

	var atms []*atm.Atm
	for _, id := range branchIds {
		for i := 0; i < 3; i++ {
			a := atm.FakeAtm(id)
			atms = append(atms, &a)
		}
	}

	err = container.AtmService.CreateAtmBulk(ctx, atms)
	if err != nil {
		fmt.Printf("create atms: %w", err.Error())
		return
	}

}
