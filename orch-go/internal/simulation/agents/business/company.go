package business

import (
	"fmt"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/bank"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Company represents a business entity.
type Company struct {
	agents.BaseAgent
	Balance         decimal.Decimal `json:"balance"`
	TargetEmployees int             `json:"target_employees"`
}

func NewCompany(name string, targetEmployees int) *Company {
	c := &Company{
		BaseAgent:       agents.NewBaseAgent(uuid.Nil, "Company", name),
		Balance:         decimal.NewFromInt(100), // Initial Capital
		TargetEmployees: targetEmployees,
	}
	return c
}

// OnTick implements the Agent interface
func (c *Company) OnTick(ctx simcontext.AgentContext) error {
	if c.CustomerId == nil {
		svcs := ctx.Services()
		err := bank.RegisterAgent(ctx, svcs, c)
		if err != nil {
			return err
		}
	}

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
	m.AddListing(c.ID(), "Basic Product", economy.ItemProduct, decimal.NewFromFloat32(10.0), 5)

	// 3. Pay salaries
	// This would require iterating employees and transferring funds via Bank/Transaction service.
	// Placeholder

	return nil
}

func (c *Company) SendSalary() {

}

func (c *Company) GetEmployees(ctx simcontext.AgentContext) []agents.Agent {
	lm := ctx.LaborMarket()
	var employees []agents.Agent
	for _, contract := range lm.Contracts {
		if contract.EmployerID == c.ID() {
			employees = append(employees, contract.EmployeeID)
		}
	}
}
