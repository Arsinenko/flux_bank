package behavior

import (
	"orch-go/internal/simulation/agents"
	simcontext "orch-go/internal/simulation/context"
)

type Action interface {
	Score(a *agents.Individual, ctx simcontext.AgentContext) float64
	Execute(a *agents.Individual, ctx simcontext.AgentContext) error
}
