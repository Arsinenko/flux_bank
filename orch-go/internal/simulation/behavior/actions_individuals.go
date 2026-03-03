package behavior

import (
	"orch-go/internal/domain/transaction"
	"orch-go/internal/simulation/agents"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"
	"time"

	"github.com/shopspring/decimal"
)

type BuyFoodAction struct{}

func (b *BuyFoodAction) Execute(a *agents.Individual, ctx simcontext.AgentContext) error {
	amount := decimal.NewFromFloat(100)
	item := ctx.Market().FindBy(
		economy.WithMaxPrice(amount),
		economy.WithNameContains("food"),
	)[0]
	if item == nil {
		return nil
	}
	var shopAccount *int32
	for _, agentInterface := range ctx.Agents() {
		if agent, ok := agentInterface.(agents.Agent); ok && agent.Type() == "Shop" {
			if agent.ID() == item.SellerID {
				shopAccount = agent.GetAccountID()
				break
			}
		}
	}
	now := time.Now()
	status := "success"
	_, err := ctx.Services().TransactionService.CreateTransaction(ctx, &transaction.Transaction{
		TransactionID: 0,
		SourceAccount: a.AccountId,
		TargetAccount: shopAccount,
		Amount:        item.Price,
		Currency:      "rub",
		CreatedAt:     &now,
		Status:        &status,
	})
	if err != nil {
		return err
	}
	a.Balance = a.Balance.Sub(item.Price)
	return nil

}

func (b *BuyFoodAction) Score(a *agents.Individual, ctx simcontext.AgentContext) float64 {
	hunger := a.Needs["hunger"]

	canAfford := Linear(a.Balance.InexactFloat64(), 0, 50)

	return ResponseCurve(hunger, 2.0) * canAfford
}

//type TakeLoanAction struct{}
//
//func (t TakeLoanAction) Score(a *agents.Individual, ctx simcontext.AgentContext) float64 {
//	despair := Linear(a.Balance.InexactFloat64(), 50, -100)
//
//	// 2. Инвертируем: чем меньше баланс, тем выше Score
//	score := ResponseCurve(despair, 1.5)
//
//	// Если у агента уже огромные долги, Score можно снизить (страх невозврата)
//	if i.DebtTotal.GreaterThan(decimal.NewFromInt(5000)) {
//		score *= 0.5
//	}
//
//	return score
//}
//
//func (t TakeLoanAction) Execute(a *agents.Individual, ctx simcontext.AgentContext) error {
//	//TODO implement me
//	panic("implement me")
//}
