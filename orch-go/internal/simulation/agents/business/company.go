package business

import (
	"fmt"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

// Company represents a business entity.
type Company struct {
	agents.BaseAgent
	Name      string
	Balance   float64
	Employees []string // Store Employee IDs for now
	// Logic parameters
	TargetEmployees int
}

func NewCompany(name string, targetEmployees int) *Company {
	c := &Company{
		BaseAgent:       agents.NewBaseAgent(uuid.Nil, "Company"),
		Name:            name,
		Balance:         10000.0, // Initial Capital
		TargetEmployees: targetEmployees,
	}
	return c
}

// OnTick implements the Agent interface
func (c *Company) OnTick(ctx agents.AgentContext) error {
	// 1. Check if we need to hire
	if len(c.Employees) < c.TargetEmployees {
		// Post vacancy if not already posted (simplified logic: just post every tick if understaffed and have budget)
		// Real logic would check existing vacancies.

		// For now simple stochastic posting
		// Using the LabotMarket from context
		lm := ctx.LaborMarket()
		lm.PostVacancy(c.ID(), fmt.Sprintf("Worker at %s", c.Name), 500.0) // Fixed salary 500
	}

	// 2. Produce and sell (simplified)
	// Add products to market
	m := ctx.Market()
	m.AddListing(c.ID(), "Basic Product", economy.ItemProduct, 10.0, 5)

	// 3. Pay salaries
	// This would require iterating employees and transferring funds via Bank/Transaction service.
	// Placeholder

	return nil
}
