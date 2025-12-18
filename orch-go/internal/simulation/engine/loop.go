package engine

import (
	"context"
	"fmt"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/economy"
	"orch-go/internal/simulation/types"
	"sync"
	"time"
)

type SimulationEngine struct {
	Clock       *Clock
	Agents      []agents.Agent
	Market      *economy.MarketRegistry
	LaborMarket *economy.LaborMarket

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

func NewSimulationEngine(startTime time.Time, tickInterval time.Duration) *SimulationEngine {
	return &SimulationEngine{
		Clock:       NewClock(startTime, tickInterval),
		Agents:      make([]agents.Agent, 0),
		Market:      economy.NewMarketRegistry(),
		LaborMarket: economy.NewLaborMarket(),
		state:       StateStopped,
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
			tickInfo := e.Clock.Tick()

			// Create context for this tick
			simCtx := &SimpleSimulationContext{
				Context:     ctx,
				time:        tickInfo.CurrentTime,
				tickNumber:  tickInfo.TickNumber,
				market:      e.Market,
				laborMarket: e.LaborMarket,
			}

			// Notify all agents
			// TODO: Parallelize this if needed
			e.mu.RLock()
			currentAgents := make([]agents.Agent, len(e.Agents))
			copy(currentAgents, e.Agents)
			e.mu.RUnlock()

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

// SimpleSimulationContext implements AgentContext
type SimpleSimulationContext struct {
	context.Context
	time        types.SimulationTime
	tickNumber  uint64
	market      *economy.MarketRegistry
	laborMarket *economy.LaborMarket
}

func (s *SimpleSimulationContext) Time() types.SimulationTime {
	return s.time
}

func (s *SimpleSimulationContext) Tick() uint64 {
	return s.tickNumber
}

func (s *SimpleSimulationContext) Market() *economy.MarketRegistry {
	return s.market
}

func (s *SimpleSimulationContext) LaborMarket() *economy.LaborMarket {
	return s.laborMarket
}
