package simulation_context

import (
	"context"
	"orch-go/internal/services"
	"orch-go/internal/simulation/agents"
	"sync"
)

type SimulationContext interface {
	Services() *services.ServiceContainer
	Agents() []*agents.Agent
}

type SimpleSimulationContext struct {
	context.Context
	services *services.ServiceContainer
	agents   []*agents.Agent
	mu       sync.RWMutex
}

func NewSimpleSimulationContext(ctx context.Context, services *services.ServiceContainer) SimulationContext {
	return &SimpleSimulationContext{
		Context:  ctx,
		services: services,
	}
}

// Agents implements [SimulationContext].
func (s *SimpleSimulationContext) Agents() []*agents.Agent {
	//return copy of agents
	s.mu.RLock()
	defer s.mu.RUnlock()

	agentsCopy := make([]*agents.Agent, len(s.agents))
	copy(agentsCopy, s.agents)

	return agentsCopy
}

func (s *SimpleSimulationContext) AddAgent(agent *agents.Agent) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.agents = append(s.agents, agent)
}

// Services implements [SimulationContext].
func (s *SimpleSimulationContext) Services() *services.ServiceContainer {
	return s.services
}
