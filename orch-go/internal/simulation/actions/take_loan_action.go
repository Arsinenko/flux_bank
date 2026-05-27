package actions

import (
	"orch-go/internal/domain/loan"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/simulation_context"
	"time"

	"github.com/shopspring/decimal"
)

type TakeLoanAction struct{}

func (a TakeLoanAction) Execute(ctx *simulation_context.SimpleSimulationContext, agent *agents.Agent) error {
	account, err := ctx.Services().AccountService.GetAccountById(ctx, agent.AccountId)
	if err != nil {
		return err
	}
	if account.Balance.LessThan(agent.Salary.Mul(decimal.NewFromFloat(0.1))) {
		return nil
	}
	loanAmount := agent.Salary.String()

	createLoan := &loan.Loan{
		LoanID:       0,
		CustomerID:   &agent.CustomerId,
		Principal:    &loanAmount,
		InterestRate: &[]string{"0.1"}[0],
		StartDate:    &[]time.Time{time.Now()}[0],
		EndDate:      &[]time.Time{time.Now().Add(time.Hour * 24 * 365)}[0],
		Status:       &[]string{"active"}[0],
	}
	createLoan, err = ctx.Services().LoanService.CreateLoan(ctx, createLoan)
	if err != nil {
		return err
	}

	var payments []loan.LoanPayment
	paymentAmount, _ := decimal.NewFromString(*createLoan.Principal)
	paymentAmount = paymentAmount.Div(decimal.NewFromInt(3))
	isPaid := false
	for i, payementDate := 0, time.Now().Add(time.Hour*24*30); i < 3; i, payementDate = i+1, payementDate.Add(time.Hour*24*30) {
		currentDate := payementDate
		payments = append(payments, loan.LoanPayment{
			LoanID:      &createLoan.LoanID,
			Amount:      paymentAmount,
			PaymentDate: &currentDate,
			IsPaid:      &isPaid,
		})
	}

	for _, payment := range payments {
		_, err = ctx.Services().LoanService.CreateLoanPayment(ctx, &payment)
		if err != nil {
			return err
		}
	}
	return nil
}
