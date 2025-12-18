package economy

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type ItemType string

const (
	ItemProduct ItemType = "Product"
	ItemService ItemType = "Service"
)

// Listing represents an item (good or service) listed for sale.
type Listing struct {
	ID          uuid.UUID
	SellerID    uuid.UUID
	Name        string
	Type        ItemType
	Price       float64
	Quantity    int // -1 for unlimited (services)
	Description string
}

// MarketRegistry manages all active listings.
type MarketRegistry struct {
	mu       sync.RWMutex
	Listings map[uuid.UUID]*Listing
}

func NewMarketRegistry() *MarketRegistry {
	return &MarketRegistry{
		Listings: make(map[uuid.UUID]*Listing),
	}
}

func (m *MarketRegistry) AddListing(sellerID uuid.UUID, name string, itemType ItemType, price float64, quantity int) uuid.UUID {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := uuid.New()
	listing := &Listing{
		ID:       id,
		SellerID: sellerID,
		Name:     name,
		Type:     itemType,
		Price:    price,
		Quantity: quantity,
	}
	m.Listings[id] = listing
	return id
}

func (m *MarketRegistry) GetListing(id uuid.UUID) (*Listing, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if l, ok := m.Listings[id]; ok {
		return l, nil
	}
	return nil, errors.New("listing not found")
}

func (m *MarketRegistry) GetAllListings() []*Listing {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Listing, 0, len(m.Listings))
	for _, l := range m.Listings {
		result = append(result, l)
	}
	return result
}

// PurchaseResult indicates the outcome of a purchase attempt
type PurchaseResult struct {
	Success bool
	Cost    float64
	Message string
}

// BuyItem attempts to purchase an item.
// Note: Actual money transfer should handle by TransactionService, this just updates quantity.
func (m *MarketRegistry) BuyItem(listingID uuid.UUID, quantity int) (*PurchaseResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	listing, ok := m.Listings[listingID]
	if !ok {
		return nil, errors.New("listing not found")
	}

	if listing.Quantity != -1 && listing.Quantity < quantity {
		return &PurchaseResult{Success: false, Message: "Not enough quantity"}, nil
	}

	totalCost := listing.Price * float64(quantity)

	if listing.Quantity != -1 {
		listing.Quantity -= quantity
		if listing.Quantity == 0 {
			delete(m.Listings, listingID)
		}
	}

	return &PurchaseResult{
		Success: true,
		Cost:    totalCost,
		Message: "Purchase successful",
	}, nil
}
