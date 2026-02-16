package engine

import (
	"context"
	"fmt"
	"orch-go/internal/services"
	"orch-go/internal/simulation/agents"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"
	"sync"
	"time"
)

type SimulationEngine struct {
	Clock            *Clock
	Agents           []agents.Agent
	Market           *economy.MarketRegistry
	LaborMarket      *economy.LaborMarket
	ServiceContainer *services.ServiceContainer

	// Synchronization
	mu    sync.RWMutex
	state EngineState
}

type EngineState int

const (
	StateStopped EngineState = iota
	StateRunning
	StatePaused
)

func NewSimulationEngine(startTime time.Time, tickInterval time.Duration, serviceContainer *services.ServiceContainer) *SimulationEngine {
	return &SimulationEngine{
		Clock:            NewClock(startTime, tickInterval),
		Agents:           make([]agents.Agent, 0),
		Market:           economy.NewMarketRegistry(),
		LaborMarket:      economy.NewLaborMarket(),
		ServiceContainer: serviceContainer,
		state:            StateStopped,
	}
}

func (e *SimulationEngine) AddAgent(a agents.Agent) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Agents = append(e.Agents, a)
}

// Run starts the simulation loop. It blocks until context is cancelled/done.
func (e *SimulationEngine) Run(ctx context.Context) error {
	e.mu.Lock()
	e.state = StateRunning
	e.mu.Unlock()

	fmt.Println("Simulation Engine Started")
	defer func() {
		e.mu.Lock()
		e.state = StateStopped
		e.mu.Unlock()
		fmt.Println("Simulation Engine Stopped")
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Execute one tick
			e.Clock.Tick()

			e.mu.RLock()
			currentAgents := make([]agents.Agent, len(e.Agents))
			copy(currentAgents, e.Agents)
			e.mu.RUnlock()

			// Convert []agents.Agent to []interface{}
			agentInterfaces := make([]interface{}, len(currentAgents))
			for i, a := range currentAgents {
				agentInterfaces[i] = a
			}

			// Create context for this tick
			simCtx := simcontext.NewSimpleSimulationContext(
				ctx,
				e.ServiceContainer,
				e.Market,
				e.LaborMarket,
				agentInterfaces,
			)

			// Notify all agents
			// TODO: Parallelize this if needed
			var wg sync.WaitGroup
			for _, a := range currentAgents {
				wg.Add(1)
				go func(agent agents.Agent) {
					defer wg.Done()
					if err := agent.OnTick(simCtx); err != nil {
						fmt.Printf("Error in agent %s: %v\n", agent.ID(), err)
					}
				}(a)
			}
			wg.Wait()

			// Optional: Sleep to match real time speed if needed,
			// but for now we run as fast as possible or controlled by caller/sleep
			time.Sleep(10 * time.Millisecond) // throttling to prevent CPU spin
		}
	}
}
