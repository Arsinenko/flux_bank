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
}
