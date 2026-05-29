package simulationrunner

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"orch-go/internal/services"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/factory"
	"orch-go/internal/simulation/simulation_context"
)

type SimulationManager struct {
	mu       sync.Mutex
	cancel   context.CancelFunc
	simCtx   *simulation_context.SimpleSimulationContext
	running  bool
	count    int
	filePath string
}

func NewSimulationManager(filePath string) *SimulationManager {
	return &SimulationManager{
		filePath: filePath,
		count:    5, // default agent count
	}
}

var DefaultManager = NewSimulationManager("agents.json")

func (m *SimulationManager) Start(services *services.ServiceContainer, count int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.running {
		return nil
	}

	if count <= 0 {
		count = m.count
	} else {
		m.count = count
	}

	ctx, cancel := context.WithCancel(context.Background())
	simCtxRaw := simulation_context.NewSimpleSimulationContext(ctx, services)
	simCtx, ok := simCtxRaw.(*simulation_context.SimpleSimulationContext)
	if !ok {
		cancel()
		return fmt.Errorf("failed to cast simulation context")
	}

	// Try loading agents from JSON
	var loadedAgents []*agents.Agent
	if data, err := os.ReadFile(m.filePath); err == nil {
		_ = json.Unmarshal(data, &loadedAgents)
	}

	var finalAgents []*agents.Agent
	if len(loadedAgents) >= count {
		finalAgents = loadedAgents[:count]
	} else {
		finalAgents = loadedAgents
		for i := len(loadedAgents); i < count; i++ {
			agent, err := factory.RegisterAgent(simCtx)
			if err != nil {
				cancel()
				return fmt.Errorf("failed to create agent: %w", err)
			}
			finalAgents = append(finalAgents, &agent)
		}
	}

	for _, a := range finalAgents {
		simCtx.AddAgent(a)
	}

	m.simCtx = simCtx
	m.cancel = cancel
	m.running = true

	go RunSimulation(simCtx)

	return nil
}

func (m *SimulationManager) Pause() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.running {
		return nil
	}

	m.cancel()
	m.running = false

	agentsList := m.simCtx.Agents()
	data, err := json.MarshalIndent(agentsList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal agents: %w", err)
	}

	err = os.WriteFile(m.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save agents to file: %w", err)
	}

	return nil
}

func (m *SimulationManager) SetAgentCount(services *services.ServiceContainer, count int) error {
	if count <= 0 {
		return fmt.Errorf("agent count must be positive")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.count = count

	if m.running {
		m.cancel()
		m.running = false

		// Save current agents
		agentsList := m.simCtx.Agents()
		data, err := json.MarshalIndent(agentsList, "", "  ")
		if err == nil {
			_ = os.WriteFile(m.filePath, data, 0644)
		}

		// Re-initialize with new count
		ctx, cancel := context.WithCancel(context.Background())
		simCtxRaw := simulation_context.NewSimpleSimulationContext(ctx, services)
		simCtx, ok := simCtxRaw.(*simulation_context.SimpleSimulationContext)
		if !ok {
			cancel()
			return fmt.Errorf("failed to cast simulation context")
		}

		// Load agents from file
		var loadedAgents []*agents.Agent
		if fileData, err := os.ReadFile(m.filePath); err == nil {
			_ = json.Unmarshal(fileData, &loadedAgents)
		}

		var finalAgents []*agents.Agent
		if len(loadedAgents) >= count {
			finalAgents = loadedAgents[:count]
		} else {
			finalAgents = loadedAgents
			for i := len(loadedAgents); i < count; i++ {
				agent, err := factory.RegisterAgent(simCtx)
				if err != nil {
					cancel()
					return fmt.Errorf("failed to create agent: %w", err)
				}
				finalAgents = append(finalAgents, &agent)
			}
		}

		for _, a := range finalAgents {
			simCtx.AddAgent(a)
		}

		m.simCtx = simCtx
		m.cancel = cancel
		m.running = true

		go RunSimulation(simCtx)
	}

	return nil
}

func (m *SimulationManager) GetStatus() (running bool, count int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.running, m.count
}
