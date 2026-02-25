package business

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/bank"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Entrepreneur represents a small business owner (IP).
type Entrepreneur struct {
	agents.BaseAgent
	Balance decimal.Decimal
	// Can verify small number of employees
}

func NewEntrepreneur(name string) *Entrepreneur {
	return &Entrepreneur{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "Entrepreneur", name),
		Balance:   decimal.NewFromInt(100),
	}
}

func (e *Entrepreneur) OnTick(ctx simcontext.AgentContext) error {
	if e.CustomerId == nil {
		svcs := ctx.Services()
		err := bank.RegisterAgent(ctx, svcs, e)
		if err != nil {
			return err
		}
	}

	// Small business logic
	// Maybe produce service
	m := ctx.Market()
	// Check listings...
	m.AddListing(e.ID(), "Small Service", economy.ItemService, decimal.NewFromFloat32(10.0), -1)
	return nil
}

// SelfEmployed represents an individual working for themselves.
type SelfEmployed struct {
	agents.BaseAgent
	Name       string
	Balance    decimal.Decimal
	CustomerID *int32
	AccountID  *int32
}

func NewSelfEmployed(name string) *SelfEmployed {
	return &SelfEmployed{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "SelfEmployed", name),
		Balance:   decimal.NewFromInt(500),
	}
}

func (s *SelfEmployed) OnTick(ctx simcontext.AgentContext) error {
	if s.CustomerID == nil {
		svcs := ctx.Services()
		err := bank.RegisterAgent(ctx, svcs, s)
		if err != nil {
			return err
		}
	}

	// Freelancer logic
	// Provide service on demand
	m := ctx.Market()
	m.AddListing(s.ID(), "Freelance Work", economy.ItemService, decimal.NewFromFloat32(10.0), -1)
	return nil
}

func (s *SelfEmployed) UpdateBalanceInfo(ctx simcontext.AgentContext) {
	acc, err := ctx.Services().AccountService.GetAccountById(ctx, *s.GetAccountID())
	if err != nil {
		return
	}
	s.Balance = acc.Balance
}
