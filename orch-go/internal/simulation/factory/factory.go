package factory

import (
	"math/rand"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/customer"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func CreateAgents(ctx *simulation_context.SimpleSimulationContext, count int) error {
	g := errgroup.Group{}
	g.SetLimit(15)
	for i := 0; i < count; i++ {
		g.Go(func() error {
			agent, err := RegisterAgent(ctx)
			if err != nil {
				return err
			}
			ctx.AddAgent(&agent)
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
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
	salary := decimal.NewFromInt(int64(50000 + rand.Intn(100000-50000-1)))
	return agents.Agent{
		CustomerId: createCustomer.Id,
		AccountId:  *createAccount.Id,
		Salary:     salary,
	}, nil

}
