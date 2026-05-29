package simulationrunner

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "orch-go/api/generated"
	account_repo "orch-go/internal/infrastructure/repository/account/account_repo"
	deposit_repo "orch-go/internal/infrastructure/repository/deposit_repo"
	loan_repo "orch-go/internal/infrastructure/repository/loan_repo"
	transaction_repo "orch-go/internal/infrastructure/repository/transaction_repo"
	"orch-go/internal/services"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
)

// Mock AccountServiceClient
type mockAccountClient struct {
	pb.AccountServiceClient
	mu       sync.Mutex
	balances map[int32]string
	getCnt   int
	updCnt   int
	onUpdate func(accountId int32, oldBal, newBal string)
}

func (m *mockAccountClient) GetById(ctx context.Context, in *pb.GetAccountByIdRequest, opts ...grpc.CallOption) (*pb.AccountModel, error) {
	m.mu.Lock()
	m.getCnt++
	bal, ok := m.balances[in.AccountId]
	if !ok {
		bal = "0"
	}
	m.mu.Unlock()

	var customerId int32 = in.AccountId // simple mapping
	var typeId int32 = 1
	var isActive bool = true
	return &pb.AccountModel{
		AccountId:  in.AccountId,
		CustomerId: &customerId,
		TypeId:     &typeId,
		Iban:       "IBAN123",
		Balance:    &bal,
		IsActive:   &isActive,
	}, nil
}

func (m *mockAccountClient) Update(ctx context.Context, in *pb.UpdateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.mu.Lock()
	oldBal := m.balances[in.AccountId]
	m.balances[in.AccountId] = *in.Balance
	m.updCnt++
	m.mu.Unlock()

	if m.onUpdate != nil {
		m.onUpdate(in.AccountId, oldBal, *in.Balance)
	}

	return &emptypb.Empty{}, nil
}

func (m *mockAccountClient) Add(ctx context.Context, in *pb.AddAccountRequest, opts ...grpc.CallOption) (*pb.AccountModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.updCnt++
	accountId := int32(m.updCnt + 100)

	bal := "0"
	if in.Balance != nil {
		bal = *in.Balance
	}
	m.balances[accountId] = bal

	return &pb.AccountModel{
		AccountId:  accountId,
		CustomerId: in.CustomerId,
		TypeId:     in.TypeId,
		Iban:       in.Iban,
		Balance:    &bal,
		IsActive:   in.IsActive,
	}, nil
}

// Mock LoanServiceClient
type mockLoanClient struct {
	pb.LoanServiceClient
	mu    sync.Mutex
	loans []*pb.AddLoanRequest
}

func (m *mockLoanClient) Add(ctx context.Context, in *pb.AddLoanRequest, opts ...grpc.CallOption) (*pb.LoanModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.loans = append(m.loans, in)
	return &pb.LoanModel{
		LoanId:       1,
		CustomerId:   in.CustomerId,
		Principal:    in.Principal,
		InterestRate: in.InterestRate,
		Status:       in.Status,
	}, nil
}

// Mock LoanPaymentServiceClient
type mockLoanPaymentClient struct {
	pb.LoanPaymentServiceClient
	mu       sync.Mutex
	payments []*pb.AddLoanPaymentRequest
}

func (m *mockLoanPaymentClient) Add(ctx context.Context, in *pb.AddLoanPaymentRequest, opts ...grpc.CallOption) (*pb.LoanPaymentModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.payments = append(m.payments, in)
	return &pb.LoanPaymentModel{
		PaymentId: 1,
		LoanId:    in.LoanId,
		Amount:    in.Amount,
		IsPaid:    in.IsPaid,
	}, nil
}

// Mock TransactionServiceClient
type mockTransactionClient struct {
	pb.TransactionServiceClient
	mu  sync.Mutex
	txs []*pb.AddTransactionRequest
}

func (m *mockTransactionClient) Add(ctx context.Context, in *pb.AddTransactionRequest, opts ...grpc.CallOption) (*pb.TransactionModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.txs = append(m.txs, in)
	return &pb.TransactionModel{
		TransactionId: 1,
		SourceAccount: in.SourceAccount,
		TargetAccount: in.TargetAccount,
		Amount:        in.Amount,
		Currency:      in.Currency,
		Status:        in.Status,
	}, nil
}

// Mock DepositServiceClient
type mockDepositClient struct {
	pb.DepositServiceClient
	mu       sync.Mutex
	deposits []*pb.AddDepositRequest
}

func (m *mockDepositClient) Add(ctx context.Context, in *pb.AddDepositRequest, opts ...grpc.CallOption) (*pb.DepositModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.deposits = append(m.deposits, in)
	return &pb.DepositModel{
		DepositId:    1,
		CustomerId:   in.CustomerId,
		Amount:       in.Amount,
		InterestRate: in.InterestRate,
		Status:       in.Status,
	}, nil
}

func TestRunSimulation(t *testing.T) {
	// Setup ctx with short tick delay so test runs instantaneously
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "tick_delay", time.Microsecond)
	defer cancel()

	// Initialize mock clients
	accClient := &mockAccountClient{balances: make(map[int32]string)}
	loanClient := &mockLoanClient{}
	loanPaymentClient := &mockLoanPaymentClient{}
	txClient := &mockTransactionClient{}
	depClient := &mockDepositClient{}

	// When Agent 2 gets doubled on tick 10 (balance increase), cancel to stop simulation.
	accClient.onUpdate = func(accountId int32, oldBal, newBal string) {
		oldD, _ := decimal.NewFromString(oldBal)
		newD, _ := decimal.NewFromString(newBal)
		if newD.GreaterThan(oldD) && accountId == 2 {
			cancel()
		}
	}

	// Set initial balances
	accClient.balances[1] = "5000"  // Agent 1: Low balance (5%)
	accClient.balances[2] = "50000" // Agent 2: Medium balance (50%)
	accClient.balances[3] = "95000" // Agent 3: High balance (95%)

	// Repositories
	accRepo := account_repo.NewRepository(accClient)
	loanRepo := loan_repo.NewLoanRepository(loanClient)
	loanPaymentRepo := loan_repo.NewLoanPaymentRepository(loanPaymentClient)
	txRepo := transaction_repo.NewTransactionRepository(txClient)
	depRepo := deposit_repo.NewRepository(depClient)

	// Service container
	serviceContainer := &services.ServiceContainer{
		AccountService:     services.NewAccountService(accRepo),
		LoanService:        services.NewLoanService(loanRepo, loanPaymentRepo),
		TransactionService: services.NewTransactionService(txRepo, transaction_repo.TransactionCategoryRepository{}, transaction_repo.TransactionFeeRepository{}, accRepo),
		DepositService:     services.NewDepositService(depRepo),
	}

	// Simulation context
	rawCtx := simulation_context.NewSimpleSimulationContext(ctx, serviceContainer)
	simCtx, ok := rawCtx.(*simulation_context.SimpleSimulationContext)
	if !ok {
		t.Fatalf("Failed to cast to *SimpleSimulationContext")
	}

	// Create and add agents
	agent1 := &agents.Agent{CustomerId: 1, AccountId: 1, Salary: decimal.NewFromInt(100000)}
	agent2 := &agents.Agent{CustomerId: 2, AccountId: 2, Salary: decimal.NewFromInt(100000)}
	agent3 := &agents.Agent{CustomerId: 3, AccountId: 3, Salary: decimal.NewFromInt(100000)}

	simCtx.AddAgent(agent1)
	simCtx.AddAgent(agent2)
	simCtx.AddAgent(agent3)

	// Run simulation (this block will finish when the context is cancelled at the 10th tick)
	RunSimulation(simCtx)

	// Verify Agent 1 (Low Balance < 10%):
	// Check that we attempted to take a loan (at least one Loan addition should be recorded)
	loanClient.mu.Lock()
	numLoans := len(loanClient.loans)
	loanClient.mu.Unlock()
	if numLoans == 0 {
		t.Errorf("Expected Agent 1 to take a loan, but got 0 loans")
	}

	// Verify Agent 2 (Medium Balance 10%-90%):
	// Check that we attempted to make a transfer (at least one Transaction addition should be recorded)
	txClient.mu.Lock()
	numTxs := len(txClient.txs)
	txClient.mu.Unlock()
	if numTxs == 0 {
		t.Errorf("Expected Agent 2 to make a transfer, but got 0 transactions")
	}

	// Verify Agent 3 (High Balance > 90%):
	// Check that we made a deposit (at least one Deposit addition should be recorded)
	depClient.mu.Lock()
	numDeposits := len(depClient.deposits)
	depClient.mu.Unlock()
	if numDeposits == 0 {
		t.Errorf("Expected Agent 3 to make a deposit, but got 0 deposits")
	}

	// Verify GetMoneyAction (Tick 10):
	// GetMoneyAction doubles the account balance.
	// Initial Agent 1: 5000 -> doubled to 10000.
	// Initial Agent 2: 50000 -> doubled to 100000.
	accClient.mu.Lock()
	finalBal1, _ := decimal.NewFromString(accClient.balances[1])
	finalBal2, _ := decimal.NewFromString(accClient.balances[2])
	accClient.mu.Unlock()

	if !finalBal1.Equal(decimal.NewFromInt(10000)) {
		t.Errorf("Expected Agent 1's balance to double to exactly 10,000 on tick 10, got %s", finalBal1)
	}
	if !finalBal2.Equal(decimal.NewFromInt(100000)) {
		t.Errorf("Expected Agent 2's balance to double to exactly 100,000 on tick 10, got %s", finalBal2)
	}
}
