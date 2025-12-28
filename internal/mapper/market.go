package mapper

import (
	"spot_instrument_service/internal/domain/markets"
	"spot_instrument_service/internal/dto"

	pb "github.com/erdedan1/shared_for_homework/proto/spot_instrument_service/gen"
)

// мне не нравится ноунеймные функции, но как будто так правильнее делать(в отдельный файл)
func toMarketDTO(m *markets.Market) dto.MarketDTO {
	roles := make([]string, len(m.AllowedRoles()))
	for i, r := range m.AllowedRoles() {
		roles[i] = string(r)
	}

	return dto.MarketDTO{
		ID:           m.ID().String(),
		Name:         m.Name(),
		Enabled:      m.Enabled(),
		AllowedRoles: roles,
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

func toProtoMarket(m dto.MarketDTO) *pb.Market {
	roles := make([]string, len(m.AllowedRoles))
	for i, r := range m.AllowedRoles {
		roles[i] = string(r)
	}

	return &pb.Market{
		Id:           m.ID,
		Name:         m.Name,
		Enabled:      m.Enabled,
		AllowedRoles: roles,
	}
}

func ToProtoMarkets(markets dto.ViewMarketsByRolesResponse) []*pb.Market {
	res := make([]*pb.Market, 0, len(markets.Markets))
	for _, m := range markets.Markets {
		res = append(res, toProtoMarket(m))
	}
	return res
}

func ToDTOViewMarketsRequest(req *pb.ViewMarketsRequest) dto.ViewMarketsByRolesRequest {
	return dto.ViewMarketsByRolesRequest{
		UserRoles: req.UserRoles,
	}
}
