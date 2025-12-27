package market

import (
	"context"
	"errors"
	"spot_instrument_service/internal/app/dto"
	"spot_instrument_service/internal/app/mapper"
	"spot_instrument_service/internal/app/repository"
)

var ErrNoMarkets = errors.New("markets not found")

type Service struct {
	marketRepo repository.Market
}

func New(marketRepo repository.Market) *Service {
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
