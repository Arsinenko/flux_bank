package business

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/bank"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

// Entrepreneur represents a small business owner (IP).
type Entrepreneur struct {
	agents.BaseAgent
	Name       string
	Balance    float64
	CustomerID *int32
	AccountID  *int32
	// Can verify small number of employees
}

func NewEntrepreneur(name string) *Entrepreneur {
	return &Entrepreneur{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "Entrepreneur"),
		Name:      name,
		Balance:   2000.0,
	}
}

func (e *Entrepreneur) OnTick(ctx simcontext.AgentContext) error {
	if e.CustomerID == nil {
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
	m.AddListing(e.ID(), "Small Service", economy.ItemService, 50.0, -1)
	return nil
}

func (e *Entrepreneur) SetCustomerID(id int32) {
	e.CustomerID = &id
}

func (e *Entrepreneur) SetAccountID(id int32) {
	e.AccountID = &id
}

func (e *Entrepreneur) GetName() string {
	return e.Name
}

// SelfEmployed represents an individual working for themselves.
type SelfEmployed struct {
	agents.BaseAgent
	Name       string
	Balance    float64
	CustomerID *int32
	AccountID  *int32
}

func NewSelfEmployed(name string) *SelfEmployed {
	return &SelfEmployed{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "SelfEmployed"),
		Name:      name,
		Balance:   500.0,
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
	m.AddListing(s.ID(), "Freelance Work", economy.ItemService, 25.0, -1)
	return nil
}

func (s *SelfEmployed) SetCustomerID(id int32) {
	s.CustomerID = &id
}

func (s *SelfEmployed) SetAccountID(id int32) {
	s.AccountID = &id
}

func (s *SelfEmployed) GetName() string {
	return s.Name
}
