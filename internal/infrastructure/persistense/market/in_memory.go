package markets

import (
	"context"
	"spot_instrument_service/internal/domain/markets"
	"sync"

	"github.com/gofrs/uuid"
)

type InMemory struct {
	markets map[uuid.UUID]*markets.Market
	mu      *sync.RWMutex
}

func NewInMemory() *InMemory {
	return &InMemory{
		markets: make(map[uuid.UUID]*markets.Market),
		mu:      &sync.RWMutex{},
	}
}

func (r *InMemory) ViewMarketsByRoles(ctx context.Context, userRoles []markets.Role) ([]*markets.Market, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	setRoles := make(map[markets.Role]struct{}, len(userRoles))
	for _, role := range userRoles {
		setRoles[role] = struct{}{}
	}
	var result []*markets.Market

	for _, m := range r.markets {
		if !m.Enabled() || m.DeletedAt() != nil {
			continue
		}

		for _, ar := range m.AllowedRoles() {
			if _, ok := setRoles[ar]; ok {
				result = append(result, m)
				break
			}
		}
	}

	return result, nil
}

func (r *InMemory) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.markets, id)
	return nil
}

func (r *InMemory) Save(m *markets.Market) (*markets.Market, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.markets[m.ID()] = m
	return m, nil
}
