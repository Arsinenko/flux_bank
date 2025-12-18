package business

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

// Entrepreneur represents a small business owner (IP).
type Entrepreneur struct {
	agents.BaseAgent
	Name    string
	Balance float64
	// Can verify small number of employees
}

func NewEntrepreneur(name string) *Entrepreneur {
	return &Entrepreneur{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "Entrepreneur"),
		Name:      name,
		Balance:   2000.0,
	}
}

func (e *Entrepreneur) OnTick(ctx agents.AgentContext) error {
	// Small business logic
	// Maybe produce service
	m := ctx.Market()
	// Check listings...
	m.AddListing(e.ID(), "Small Service", economy.ItemService, 50.0, -1)
	return nil
}

// SelfEmployed represents an individual working for themselves.
type SelfEmployed struct {
	agents.BaseAgent
	Name    string
	Balance float64
}

func NewSelfEmployed(name string) *SelfEmployed {
	return &SelfEmployed{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "SelfEmployed"),
		Name:      name,
		Balance:   500.0,
	}
}

func (s *SelfEmployed) OnTick(ctx agents.AgentContext) error {
	// Freelancer logic
	// Provide service on demand
	m := ctx.Market()
	m.AddListing(s.ID(), "Freelance Work", economy.ItemService, 25.0, -1)
	return nil
}
