package grpc

import (
	"context"
	"spot_instrument_service/internal/dto"
	"spot_instrument_service/internal/mapper"

	pb "github.com/erdedan1/shared_for_homework/proto/spot_instrument_service/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type marketSrv interface {
	ViewMarketsByRoles(ctx context.Context, userRoles dto.ViewMarketsByRolesRequest) (*dto.ViewMarketsByRolesResponse, error)
}

type GRPCService struct {
	marketSrv marketSrv
	pb.MarketServiceServer
}

func New(marketSrv marketSrv) *GRPCService {
	return &GRPCService{
		marketSrv: marketSrv,
	}
}

func (s *GRPCService) ViewMarketsByRoles(ctx context.Context, req *pb.ViewMarketsRequest) (*pb.ViewMarketsResponse, error) {
	dtoReq := mapper.ToDTOViewMarketsRequest(req)

	markets, err := s.marketSrv.ViewMarketsByRoles(ctx, dtoReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(markets.Markets) == 0 {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &pb.ViewMarketsResponse{
		Markets: mapper.ToProtoMarkets(*markets),
	}, nil
}
