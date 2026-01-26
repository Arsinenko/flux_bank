package simulation

import (
	"context"
	"fmt"
	"orch-go/internal/services"
	"orch-go/internal/simulation/engine"
	"orch-go/internal/simulation/factory"
	"time"
)

// RunSimulation initializes and runs the simulation.
func RunSimulation(ctx context.Context, serviceContainer *services.ServiceContainer) error {
	// 1. Initialize Engine
	startTime := time.Now()
	tickInterval := time.Hour // 1 virtual hour per tick? Or minute?
	simEngine := engine.NewSimulationEngine(startTime, tickInterval, serviceContainer)

	// 2. Load or Create Agents
	loadedAgents, err := LoadAgents()
	if err != nil {
		fmt.Printf("Could not load agents, starting fresh: %v\n", err)
	}

	if len(loadedAgents) > 0 {
		fmt.Printf("Loaded %d agents from state file.\n", len(loadedAgents))
		for _, agent := range loadedAgents {
			simEngine.AddAgent(agent)
		}
	} else {
		fmt.Println("No saved state found, creating new agents.")
		// Create Agents
		agentFactory := factory.NewAgentFactory()

		// 2.1 Companies
		techCorp := agentFactory.CreateCompany("TechCorp", 10)
		simEngine.AddAgent(techCorp)

		superMart := agentFactory.CreateShop("SuperMart")
		simEngine.AddAgent(superMart)

		// 2.2 Individuals
		for i := 0; i < 20; i++ {
			name := fmt.Sprintf("Individual_%d", i)
			person := agentFactory.CreateIndividual(name)
			simEngine.AddAgent(person)
		}
	}

	// 3. Run
	runErr := simEngine.Run(ctx)

	// 4. Save state on exit
	fmt.Println("Simulation ended, saving state...")
	agentsToSave := simEngine.Agents
	if saveErr := SaveAgents(agentsToSave); saveErr != nil {
		fmt.Printf("Error saving agent state: %v\n", saveErr)
	} else {
		fmt.Println("Agent state saved successfully.")
	}

	return runErr
}
