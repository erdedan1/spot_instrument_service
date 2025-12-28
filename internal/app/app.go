package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"spot_instrument_service/config"
	marketRepo "spot_instrument_service/internal/infrastructure/persistense/market"
	marketSrv "spot_instrument_service/internal/service/market"
	marketGRPC "spot_instrument_service/internal/transport/grpc"

	pbLogger "github.com/erdedan1/shared_for_homework/pkg/interceptors/logger"
	"github.com/erdedan1/shared_for_homework/pkg/interceptors/recovery"
	requestid "github.com/erdedan1/shared_for_homework/pkg/interceptors/request_id"
	pb "github.com/erdedan1/shared_for_homework/proto/spot_instrument_service/gen"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *config.Config
	grpcServer *grpc.Server
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	log.Println("starting service")
	repos := marketRepo.NewInMemory()
	srvs := marketSrv.New(repos)

	go func() {
		if err := a.startGRPCServer(srvs); err != nil {
			log.Println("failed to start grpc server", err)
			os.Exit(1)
		}
	}()
	log.Println("service work")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Println("waiting for shutdown signal")
	<-quit
	log.Println("shutdown signal received")
	a.stopGRPCServer()
	log.Println("service stopped gracefully")
	return nil
}

func (a *App) startGRPCServer(marketService *marketSrv.Service) error {
	zap, _ := zap.NewProduction()
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			requestid.XRequestIDServerInterceptor(),
			pbLogger.LoggerServerInterceptor(zap),
			recovery.RecoveryServerInterceptor(zap),
		),
	)

	a.grpcServer = grpcServer
	grpcHandler := marketGRPC.New(marketService)
	pb.RegisterMarketServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", a.cfg.GRPCServer.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("grpc serve error", err)
		return err
	}
	return nil
}

func (a *App) stopGRPCServer() {
	a.grpcServer.GracefulStop()
	log.Println("grpc server stopped gracefully")
}
