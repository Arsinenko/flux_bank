package simulation

import (
	"context"
	"fmt"
	"orch-go/internal/simulation/engine"
	"orch-go/internal/simulation/factory"
	"time"
)

// RunSimulation initializes and runs the simulation.
func RunSimulation(ctx context.Context) error {
	// 1. Initialize Engine
	startTime := time.Now()
	tickInterval := time.Hour // 1 virtual hour per tick? Or minute?
	simEngine := engine.NewSimulationEngine(startTime, tickInterval)

	// 2. Create Agents
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

	// 2.3 Entrepreneurs
	devStudio := agentFactory.CreateEntrepreneur("DevStudio")
	simEngine.AddAgent(devStudio)

	freelancer := agentFactory.CreateSelfEmployed("Alice_Freelancer")
	simEngine.AddAgent(freelancer)

	// 3. Run
	return simEngine.Run(ctx)
}
