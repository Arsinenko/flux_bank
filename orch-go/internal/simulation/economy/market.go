package economy

import (
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	Price       decimal.Decimal
	Quantity    int // -1 for unlimited (services)
	Description string
}

type ListingPredicate func(listing *Listing) bool

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

func (m *MarketRegistry) AddListing(sellerID uuid.UUID, name string, itemType ItemType, price decimal.Decimal, quantity int) uuid.UUID {
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

func (m *MarketRegistry) UpdateListing(id uuid.UUID, price decimal.Decimal, quantity int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	listing, ok := m.Listings[id]
	if !ok {
		return
	}
	listing.Price = price
	listing.Quantity = quantity
}

func (m *MarketRegistry) FindBy(predicates ...ListingPredicate) []*Listing {
	m.mu.RLock()
	defer m.mu.RUnlock()

	results := make([]*Listing, 0)
	for _, item := range m.Listings {
		match := true
		for _, predicate := range predicates {
			if !predicate(item) {
				match = false
				break
			}
		}
		if match {
			results = append(results, item)
		}
	}
	return results
}

func WithMinPrice(price decimal.Decimal) ListingPredicate {
	return func(l *Listing) bool { return l.Price.GreaterThanOrEqual(price) }
}

func WithMaxPrice(price decimal.Decimal) ListingPredicate {
	return func(l *Listing) bool { return l.Price.LessThanOrEqual(price) }
}

func WithType(itemType ItemType) ListingPredicate {
	return func(l *Listing) bool { return l.Type == itemType }
}

func WithNameContains(name string) ListingPredicate {
	return func(l *Listing) bool {
		return strings.Contains(strings.ToLower(l.Name), strings.ToLower(name))
	}
}
