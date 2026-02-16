package context

import (
	"context"
	"orch-go/internal/services"
	"orch-go/internal/simulation/economy"
)

type AgentContext interface {
	context.Context
	Services() *services.ServiceContainer
	Market() *economy.MarketRegistry
	LaborMarket() *economy.LaborMarket
	Agents() []interface{}
}

// SimpleSimulationContext is a basic implementation of SimulationContext.
type SimpleSimulationContext struct {
	context.Context
	services    *services.ServiceContainer
	market      *economy.MarketRegistry
	laborMarket *economy.LaborMarket
	agents      []interface{}
}

// NewSimpleSimulationContext creates a new SimpleSimulationContext.
func NewSimpleSimulationContext(
	ctx context.Context,
	services *services.ServiceContainer,
	market *economy.MarketRegistry,
	laborMarket *economy.LaborMarket,
	agents []interface{},
) *SimpleSimulationContext {
	return &SimpleSimulationContext{
		Context:     ctx,
		services:    services,
		market:      market,
		laborMarket: laborMarket,
		agents:      agents,
	}
}

func (s *SimpleSimulationContext) Services() *services.ServiceContainer {
	return s.services
}

func (s *SimpleSimulationContext) Market() *economy.MarketRegistry {
	return s.market
}

func (s *SimpleSimulationContext) LaborMarket() *economy.LaborMarket {
	return s.laborMarket
}

func (s *SimpleSimulationContext) Agents() []interface{} {
	return s.agents
}
