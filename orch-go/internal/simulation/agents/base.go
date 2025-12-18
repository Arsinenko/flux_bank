package agents

import (
	"orch-go/internal/simulation/economy"
	"orch-go/internal/simulation/types"

	"github.com/google/uuid"
)

// AgentContext provides access to external systems for the agent (Economy, other agents, etc.)
// Detailed definition will be added as we implement the Economy package.
type AgentContext interface {
	types.SimulationContext
	Market() *economy.MarketRegistry
	LaborMarket() *economy.LaborMarket
}

// Agent is the main interface that all simulation agents must implement.
type Agent interface {
	// ID returns the unique identifier of the agent.
	ID() uuid.UUID

	// Type returns the type of the agent (Company, Individual, etc.)
	Type() string

	// OnTick is called every simulation step.
	// agents should decide their actions based on the current state and time.
	OnTick(ctx AgentContext) error
}

// BaseAgent provides common functionality for all agents.
type BaseAgent struct {
	id        uuid.UUID
	agentType string
}

func NewBaseAgent(id uuid.UUID, agentType string) BaseAgent {
	if id == uuid.Nil {
		id = uuid.New()
	}
	return BaseAgent{
		id:        id,
		agentType: agentType,
	}
}

func (b *BaseAgent) ID() uuid.UUID {
	return b.id
}

func (b *BaseAgent) Type() string {
	return b.agentType
}
