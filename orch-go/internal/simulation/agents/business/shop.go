package business

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/economy"

	"github.com/google/uuid"
)

// Shop represents a retail business.
type Shop struct {
	agents.BaseAgent
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	// Inventory map[uuid.UUID]int // Products bought from companies to resell?
	// Or maybe Shops produce "Retail Service".
	// Let's assume shops buy wholesale and sell retail.
}

func NewShop(name string) *Shop {
	return &Shop{
		BaseAgent: agents.NewBaseAgent(uuid.Nil, "Shop"),
		Name:      name,
		Balance:   5000.0,
	}
}

func (s *Shop) OnTick(ctx agents.AgentContext) error {
	// Simplified logic: Shops just list goods they magically have for now,
	// or we can make them buy from Companies if we link them.

	// For this stage, let's just make the Shop listing items.
	m := ctx.Market()
	// Check if we have listings, if not create some
	// In a real tick loop we might not want to spam listings.

	// Let's try to maintain a listing.
	// Since we don't query our own listings easily yet without searching registry,
	// we'll just add one blindly for now or improve registry.

	m.AddListing(s.ID(), "Retail Goods", economy.ItemProduct, 15.0, 10)

	return nil
}
