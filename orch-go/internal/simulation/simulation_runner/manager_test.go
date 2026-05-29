package simulationrunner

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	pb "orch-go/api/generated"
	account_repo "orch-go/internal/infrastructure/repository/account/account_repo"
	customer_repo "orch-go/internal/infrastructure/repository/customer_repo"
	deposit_repo "orch-go/internal/infrastructure/repository/deposit_repo"
	loan_repo "orch-go/internal/infrastructure/repository/loan_repo"
	transaction_repo "orch-go/internal/infrastructure/repository/transaction_repo"
	"orch-go/internal/services"
	"orch-go/internal/simulation/agents"

	"google.golang.org/grpc"
)

// Mock CustomerServiceClient
type mockCustomerClient struct {
	pb.CustomerServiceClient
	mu      sync.Mutex
	custCnt int
}

func (m *mockCustomerClient) Add(ctx context.Context, in *pb.AddCustomerRequest, opts ...grpc.CallOption) (*pb.CustomerModel, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.custCnt++

	var birthDate *pb.DateOnly
	if in.BirthDate != nil {
		birthDate = &pb.DateOnly{
			Year:  in.BirthDate.Year,
			Month: in.BirthDate.Month,
			Day:   in.BirthDate.Day,
		}
	}

	return &pb.CustomerModel{
		CustomerId: int32(m.custCnt),
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		Email:      in.Email,
		Phone:      in.Phone,
		BirthDate:  birthDate,
	}, nil
}

func TestSimulationManager(t *testing.T) {
	// Setup temporary file path for manager JSON
	tempFile := "test_agents.json"
	defer os.Remove(tempFile)

	// Initialize mock clients
	accClient := &mockAccountClient{balances: make(map[int32]string)}
	loanClient := &mockLoanClient{}
	loanPaymentClient := &mockLoanPaymentClient{}
	txClient := &mockTransactionClient{}
	depClient := &mockDepositClient{}
	custClient := &mockCustomerClient{}

	// Setup custom rand seed for factory logic inside RegisterAgent
	rand.Seed(time.Now().UnixNano())

	// Repositories
	accRepo := account_repo.NewRepository(accClient)
	loanRepo := loan_repo.NewLoanRepository(loanClient)
	loanPaymentRepo := loan_repo.NewLoanPaymentRepository(loanPaymentClient)
	txRepo := transaction_repo.NewTransactionRepository(txClient)
	depRepo := deposit_repo.NewRepository(depClient)
	custRepo := customer_repo.NewRepository(custClient)

	// Service container
	serviceContainer := &services.ServiceContainer{
		AccountService:     services.NewAccountService(accRepo),
		CustomerService:    services.NewCustomerService(custRepo),
		LoanService:        services.NewLoanService(loanRepo, loanPaymentRepo),
		TransactionService: services.NewTransactionService(txRepo, transaction_repo.TransactionCategoryRepository{}, transaction_repo.TransactionFeeRepository{}, accRepo),
		DepositService:     services.NewDepositService(depRepo),
	}

	manager := NewSimulationManager(tempFile)

	// Step 1: Start manager with 2 agents. Since agents.json doesn't exist, it should create them.
	err := manager.Start(serviceContainer, 2)
	if err != nil {
		t.Fatalf("failed to start manager: %v", err)
	}

	running, count := manager.GetStatus()
	if !running {
		t.Errorf("expected running to be true")
	}
	if count != 2 {
		t.Errorf("expected count to be 2, got %d", count)
	}

	// Wait a tiny bit and Pause the manager
	time.Sleep(10 * time.Millisecond)
	err = manager.Pause()
	if err != nil {
		t.Fatalf("failed to pause manager: %v", err)
	}

	running, count = manager.GetStatus()
	if running {
		t.Errorf("expected running to be false after pause")
	}

	// Verify that the file test_agents.json exists and has 2 agents
	fileData, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("failed to read agents file: %v", err)
	}

	var savedAgents []*agents.Agent
	err = json.Unmarshal(fileData, &savedAgents)
	if err != nil {
		t.Fatalf("failed to unmarshal saved agents: %v", err)
	}

	if len(savedAgents) != 2 {
		t.Errorf("expected 2 saved agents, got %d", len(savedAgents))
	}

	// Step 2: Increase count to 4 agents.
	// Since we are not running, it should just update the count in manager.
	err = manager.SetAgentCount(serviceContainer, 4)
	if err != nil {
		t.Fatalf("failed to set agent count: %v", err)
	}

	running, count = manager.GetStatus()
	if running {
		t.Errorf("expected running to still be false")
	}
	if count != 4 {
		t.Errorf("expected count to be updated to 4, got %d", count)
	}

	// Step 3: Start simulation again.
	// It should load the 2 saved agents from the file, and create 2 additional agents (total 4).
	err = manager.Start(serviceContainer, 4)
	if err != nil {
		t.Fatalf("failed to start manager with count 4: %v", err)
	}

	running, count = manager.GetStatus()
	if !running {
		t.Errorf("expected running to be true")
	}
	if count != 4 {
		t.Errorf("expected count to be 4, got %d", count)
	}

	// Wait a tiny bit and Pause again
	time.Sleep(10 * time.Millisecond)
	err = manager.Pause()
	if err != nil {
		t.Fatalf("failed to pause manager: %v", err)
	}

	// Check file again: it should now have 4 agents saved!
	fileData, err = os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("failed to read agents file second time: %v", err)
	}

	err = json.Unmarshal(fileData, &savedAgents)
	if err != nil {
		t.Fatalf("failed to unmarshal saved agents: %v", err)
	}

	if len(savedAgents) != 4 {
		t.Errorf("expected 4 saved agents, got %d", len(savedAgents))
	}
}
