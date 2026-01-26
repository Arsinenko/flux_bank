package agents

import (
	"fmt"
	"math/rand/v2"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/customer"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

type Individual struct {
	BaseAgent
	Name       string                      `json:"name"`
	Balance    float64                     `json:"balance"`
	Contract   *economy.EmploymentContract `json:"contract"`
	Needs      map[string]float64          `json:"needs"`
	CustomerID *int32                      `json:"customer_id"`
	AccountID  *int32                      `json:"account_id"`
}

func NewIndividual(name string) *Individual {
	return &Individual{
		BaseAgent: NewBaseAgent(uuid.Nil, "Individual"),
		Name:      name,
		Balance:   100.0, // Starting money
		Needs:     make(map[string]float64),
	}
}

func (i *Individual) OnTick(ctx AgentContext) error {
	// 1. Bank Registration Logic
	if i.CustomerID == nil {
		svcs := ctx.Services()
		// Simplified customer creation, assuming we need more details in a real scenario
		//TODO fill Customer fields
		customer, err := svcs.CustomerService.CreateCustomer(ctx, &customer.Customer{
			FirstName: i.Name,
			LastName:  "",
			Email:     "",
			Phone:     nil,
			BirthDate: nil,
			CreatedAt: nil,
		})
		if err != nil {
			return fmt.Errorf("agent %s failed to create customer: %w", i.Name, err)
		}
		i.CustomerID = &customer.Id

		// Assuming default account type, currency etc.
		account, err := svcs.AccountService.CreateAccount(ctx, &account.Account{
			//TODO fill
		})
		if err != nil {
			return fmt.Errorf("agent %s failed to create account: %w", i.Name, err)
		}
		i.AccountID = account.Id
		fmt.Printf("Individual %s registered in bank. CustomerID: %d, AccountID: %d\n", i.Name, *i.CustomerID, *i.AccountID)
	}

	// 2. Employment Logic
	if i.Contract == nil {
		// Look for job
		lm := ctx.LaborMarket()
		vacancies := lm.GetVacancies()
		if len(vacancies) > 0 {
			// Apply to random vacancy
			idx := rand.IntN(len(vacancies))
			v := vacancies[idx]
			contract, err := lm.ApplyAndHire(v.ID, i.ID())
			if err == nil {
				i.Contract = contract
				fmt.Printf("Individual %s hired by %s\n", i.Name, v.EmployerID)
			}
		}
	}

	// 3. Consumption logic
	i.consume(ctx)

	return nil
}

func (i *Individual) consume(ctx AgentContext) {
	// Simple logic: buy something if we have money and random chance
	if i.Balance > 10.0 && rand.Float64() < 0.2 {
		m := ctx.Market()
		listings := m.GetAllListings()
		if len(listings) > 0 {
			// Buy random thing
			idx := rand.IntN(len(listings))
			l := listings[idx]

			// Attempt purchase 1 unit
			res, err := m.BuyItem(l.ID, 1)
			if err == nil && res.Success {
				i.Balance -= res.Cost
				// fmt.Printf("Individual %s bought %s for %.2f\n", i.Name, l.Name, res.Cost)
			}
		}
	}
}
