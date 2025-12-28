package market

import (
	"context"
	"errors"
	"spot_instrument_service/internal/domain/markets"
	"spot_instrument_service/internal/dto"
	"spot_instrument_service/internal/mapper"

	"github.com/gofrs/uuid"
)

var ErrNoMarkets = errors.New("markets not found")

type marketRepo interface {
	ViewMarketsByRoles(ctx context.Context, userRoles []markets.Role) ([]*markets.Market, error)
	Delete(id uuid.UUID) error
	Save(m *markets.Market) (*markets.Market, error)
}

type Service struct {
	marketRepo marketRepo
}

func New(marketRepo marketRepo) *Service {
	return &Service{
		marketRepo: marketRepo,
	}
}

func (s *Service) ViewMarketsByRoles(ctx context.Context, userRoles dto.ViewMarketsByRolesRequest) (*dto.ViewMarketsByRolesResponse, error) {
	if err := userRoles.Validate(); err != nil {
		return nil, err
	}
	roles, err := mapper.ToDomainRoles(userRoles)
	if err != nil {
		return nil, err
	}
	markets, err := s.marketRepo.ViewMarketsByRoles(ctx, roles)
	if err != nil {
		return nil, err
	}

	if markets == nil {
		return nil, ErrNoMarkets
	}

	return mapper.ToViewMarketsResponse(markets), nil
}
