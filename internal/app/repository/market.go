package repository

import (
	"context"
	"spot_instrument_service/internal/app/domain/markets"

	"github.com/gofrs/uuid"
)

type Market interface {
	ViewMarketsByRoles(ctx context.Context, userRoles []markets.Role) ([]*markets.Market, error)
	Save(m *markets.Market) (*markets.Market, error)
	Delete(id uuid.UUID) error
}
