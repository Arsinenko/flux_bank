package agents

import (
	simcontext "orch-go/internal/simulation/context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Agent is the main interface that all simulation agents must implement.
type Agent interface {
	// ID returns the unique identifier of the agent.
	ID() uuid.UUID

	// Type returns the type of the agent (Company, Individual, etc.)
	Type() string

	// OnTick is called every simulation step.
	// agents should decide their actions based on the current state and time.
	OnTick(ctx simcontext.AgentContext) error

	GetAccountID() *int32
	GetCustomerID() *int32
}

// BaseAgent provides common functionality for all agents.
type BaseAgent struct {
	AgentID    uuid.UUID       `json:"id"`
	AgentType  string          `json:"type"`
	Name       string          `json:"name"`
	AccountId  *int32          `json:"account_id"`
	CustomerId *int32          `json:"customer_id"`
	Balance    decimal.Decimal `json:"balance"`
}

func NewBaseAgent(id uuid.UUID, agentType string, name string, balance decimal.Decimal) BaseAgent {
	if id == uuid.Nil {
		id = uuid.New()
	}
	return BaseAgent{
		AgentID:   id,
		AgentType: agentType,
		Name:      name,
		Balance:   balance,
	}
}

func (b *BaseAgent) GetName() string { return b.Name }

func (b *BaseAgent) ID() uuid.UUID {
	return b.AgentID
}

func (b *BaseAgent) Type() string {
	return b.AgentType
}

func (b *BaseAgent) SetCustomerID(id int32) {
	b.CustomerId = &id
}

func (b *BaseAgent) SetAccountID(id int32) {
	b.AccountId = &id
}

func (b *BaseAgent) GetAccountID() *int32 {
	return b.AccountId
}

func (b *BaseAgent) GetCustomerID() *int32 {
	return b.CustomerId
}

func (b *BaseAgent) UpdateBalanceInfo(ctx simcontext.AgentContext) {
	acc, err := ctx.Services().AccountService.GetAccountById(ctx, *b.GetAccountID())
	if err != nil {
		return
	}
	b.Balance = acc.Balance
}
