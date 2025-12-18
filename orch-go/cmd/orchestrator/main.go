package main

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/config"
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
	"orch-go/internal/simulation"
	"sync"

	"google.golang.org/grpc"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
var wg = sync.WaitGroup{}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	//Init gRPC clients
	conn, err := grpc.NewClient(cfg.Core.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	accountClient := pb.NewAccountServiceClient(conn)
	accountTypeClient := pb.NewAccountTypeServiceClient(conn)
	atmClient := pb.NewAtmServiceClient(conn)
	branchClient := pb.NewBranchServiceClient(conn)
	cardClient := pb.NewCardServiceClient(conn)
	customerAddressClient := pb.NewCustomerAddressServiceClient(conn)
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
	_ = services.NewAccountService(accountRepo)
	_ = services.NewAccountTypeService(accountTypeRepo)
	_ = services.NewAtmService(atmRepo)
	_ = services.NewBranchService(branchRepo)
	_ = services.NewCardService(cardRepo)
	_ = services.NewCustomerService(customerRepo)
	// TODO: CustomerAddressService?
	_ = customerAddressRepo // Placeholder usage
	_ = services.NewDepositService(depositRepo)
	_ = services.NewExchangeRateService(exchangeRateRepo)
	_ = services.NewFeeTypeService(feeTypeRepo)
	_ = services.NewLoanService(loanRepo, loanPaymentRepo)
	_ = services.NewLoginLogService(loginLogRepo)
	_ = services.NewNotificationService(notificationRepo)
	_ = services.NewPaymentTemplateService(paymentTemplateRepo)
	_ = services.NewTransactionService(transactionRepo, transactionCategoryRepo, transactionFeeRepo)
	_ = services.NewUserCredentialService(userCredentialRepo)
	wg.Add(1)
	go func(ctx context.Context) {
		err := simulation.RunSimulation(ctx)
		if err != nil {
			panic(err)
		}
	}(context.Background())
	wg.Wait()

}
