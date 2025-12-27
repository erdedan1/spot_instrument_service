package mapper

import (
	"spot_instrument_service/internal/app/domain/markets"
	"spot_instrument_service/internal/app/dto"
)

func toMarketDTO(m *markets.Market) dto.MarketDTO {
	roles := make([]string, len(m.AllowedRoles()))
	for i, r := range m.AllowedRoles() {
		roles[i] = string(r)
	}

	return dto.MarketDTO{
		ID:      m.ID().String(),
		Name:    m.Name(),
		Enabled: m.Enabled(),
		Roles:   roles,
	}
}

func ToViewMarketsResponse(markets []*markets.Market) *dto.ViewMarketsByRolesResponse {
	dtos := make([]dto.MarketDTO, len(markets))
	for i, m := range markets {
		dtos[i] = toMarketDTO(m)
	}
	return &dto.ViewMarketsByRolesResponse{Markets: dtos}
}

func ToDomainRoles(req dto.ViewMarketsByRolesRequest) ([]markets.Role, error) {
	if len(req.UserRoles) == 0 {
		return nil, dto.ErrEmptyRole
	}

	var roles []markets.Role
	for _, r := range req.UserRoles {
		roles = append(roles, markets.Role(r))
	}

	return roles, nil
}
