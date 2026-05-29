package simulationrunner

import (
	"math/rand/v2"
	"orch-go/internal/simulation/actions"
	"orch-go/internal/simulation/simulation_context"
	"time"

	"github.com/shopspring/decimal"
)

func RunSimulation(ctx *simulation_context.SimpleSimulationContext) {
	for tick := 1; ; tick++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		agents := ctx.Agents()
		for _, agent := range agents {
			if tick%10 == 0 {
				_ = actions.GetMoneyAction{}.Execute(ctx, agent)
				continue
			}

			balance, err := actions.GetCurrentMoneyAction{}.Execute(ctx, *agent)
			if err != nil {
				continue
			}

			tenPercent := agent.Salary.Mul(decimal.NewFromFloat(0.1))
			ninetyPercent := agent.Salary.Mul(decimal.NewFromFloat(0.9))

			if balance.LessThan(tenPercent) {
				// Less than 10% of salary: idle or take loan (50/50)
				if rand.N(2) == 1 {
					_ = actions.TakeLoanAction{}.Execute(ctx, agent)
				}
			} else if balance.LessThanOrEqual(ninetyPercent) {
				// More than 10% and less than 90% (inclusive): idle or random transfer (50/50)
				if rand.N(2) == 1 {
					_ = actions.MakeTransfer{}.Execute(ctx, agent)
				}
			} else {
				// More than 90%: idle, deposit, or transfer (33.3% chance each)
				choice := rand.N(3)
				switch choice {
				case 1:
					_ = actions.MakeDepositAction{}.Execute(ctx, agent)
				case 2:
					_ = actions.MakeTransfer{}.Execute(ctx, agent)
				}
			}
		}

		tickDelay := 100 * time.Millisecond
		if delayVal := ctx.Value("tick_delay"); delayVal != nil {
			if d, ok := delayVal.(time.Duration); ok {
				tickDelay = d
			}
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(tickDelay):
		}
	}
}
