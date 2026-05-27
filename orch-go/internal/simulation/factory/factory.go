package factory

import (
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/customer"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
	"time"
)

func CreateAgents(ctx *simulation_context.SimpleSimulationContext, count int) error {
	for i := 0; i < count; i++ {
		agent, err := RegisterAgent(ctx)
		if err != nil {
			return err
		}
		ctx.AddAgent(&agent)
	}
	return nil // TODO: with errgroup
}

func RegisterAgent(ctx *simulation_context.SimpleSimulationContext) (agents.Agent, error) {
	c := customer.FakeCustomer(time.Now())
	createCustomer, err := ctx.Services().CustomerService.CreateCustomer(ctx, c)
	if err != nil {
		return agents.Agent{}, err
	}
	a := account.FakeAccount(createCustomer.Id, 1)
	createAccount, err := ctx.Services().AccountService.CreateAccount(ctx, a)
	if err != nil {
		return agents.Agent{}, err
	}
	return agents.Agent{
		CustomerId: createCustomer.Id,
		AccountId:  *createAccount.Id,
	}, nil

}
