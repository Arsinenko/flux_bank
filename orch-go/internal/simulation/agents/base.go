package agents

import (
	simcontext "orch-go/internal/simulation/context"

	"github.com/google/uuid"
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
}

// BaseAgent provides common functionality for all agents.
type BaseAgent struct {
	AgentID   uuid.UUID `json:"id"`
	AgentType string    `json:"type"`
}

func NewBaseAgent(id uuid.UUID, agentType string) BaseAgent {
	if id == uuid.Nil {
		id = uuid.New()
	}
	return BaseAgent{
		AgentID:   id,
		AgentType: agentType,
	}
}

func (b *BaseAgent) ID() uuid.UUID {
	return b.AgentID
}

func (b *BaseAgent) Type() string {
	return b.AgentType
}
