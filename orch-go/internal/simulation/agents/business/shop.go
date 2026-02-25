package business

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/bank"
	simcontext "orch-go/internal/simulation/context"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Shop represents a retail business.
type Shop struct {
	agents.BaseAgent
	// Inventory map[uuid.UUID]int // Products bought from companies to resell?
	// Or maybe Shops produce "Retail Service".
	// Let's assume shops buy wholesale and sell retail.
}

func NewShop(name string) *Shop {
	return &Shop{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "Shop", name, decimal.NewFromInt(100)),
	}
}

func (s *Shop) OnTick(ctx simcontext.AgentContext) error {
	if s.CustomerId == nil {
		svcs := ctx.Services()
		err := bank.RegisterAgent(ctx, svcs, s)
		if err != nil {
			return err
		}
	}

	// Simplified logic: Shops just list goods they magically have for now,
	// or we can make them buy from Companies if we link them.

	// For this stage, let's just make the Shop listing items.
	m := ctx.Market()
	// Check if we have listings, if not create some
	// In a real tick loop we might not want to spam listings.

	// Let's try to maintain a listing.
	// Since we don't query our own listings easily yet without searching registry,
	// we'll just add one blindly for now or improve registry.

	m.AddListing(s.ID(), "Retail Goods", economy.ItemProduct, decimal.NewFromFloat32(10.0), 10)

	return nil
}
