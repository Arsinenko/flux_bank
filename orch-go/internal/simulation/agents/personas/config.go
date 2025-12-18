package personas

import "time"

type Config struct {
	Name             string
	ActionInterval   time.Duration
	CheckBalanceProb float32
	TransferProb     float32
	DepositProb      float32
	LoanProb         float32
	PaymentProb      float32
}

var SimpleAgent = Config{
	Name:             "SimpleAgent",
	ActionInterval:   time.Second * 5,
	CheckBalanceProb: 0.5,
	TransferProb:     0.5,
}

var SaverPersona = Config{
	Name:             "Saver",
	ActionInterval:   time.Second * 10,
	CheckBalanceProb: 0.6,
	DepositProb:      0.3,
	TransferProb:     0.1,
}

var SpenderPersona = Config{
	Name:             "Spender",
	ActionInterval:   time.Second * 3,
	CheckBalanceProb: 0.2,
	TransferProb:     0.5,
	PaymentProb:      0.3,
}

var InvestorPersona = Config{
	Name:             "Investor",
	ActionInterval:   time.Second * 15,
	CheckBalanceProb: 0.4,
	DepositProb:      0.3,
	LoanProb:         0.3,
}
