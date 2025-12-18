package agents

import (
	"fmt"
	"math/rand/v2"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

type Individual struct {
	BaseAgent
	Name     string
	Balance  float64
	Contract *economy.EmploymentContract
	Needs    map[string]float64 // Need name -> Level (0-100)
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
	// 1. Employment Logic
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
	} else {
		// Work logic...
		// In a complex sim, receive salary periodically.
		// For now, assume simplified instant payment or generic income handling elsewhere.
	}

	// 2. Consumption logic
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
