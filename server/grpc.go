package server

import (
	"context"
	"fmt"
	"github.com/tidepool-org/devices/api"
	"github.com/tidepool-org/devices/config"
	"github.com/tidepool-org/devices/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	grpcReflection "google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

const serviceName = "api.Devices"

type Params struct {
	Cfg   *config.Config
	Cgms  repo.CgmsRepo
	Pumps repo.PumpsRepo
}

type GrpcServer struct {
	api.UnimplementedDevicesServer

	grpcServer   *grpc.Server
	healthServer *health.Server

	cfg   *config.Config
	cgms  repo.CgmsRepo
	pumps repo.PumpsRepo
}

var _ api.DevicesServer = &GrpcServer{}

func NewGrpcServer(p *Params) *GrpcServer {
	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	srvr := &GrpcServer{
		grpcServer:   grpcServer,
		healthServer: healthServer,
		cfg:          p.Cfg,
		cgms:         p.Cgms,
		pumps:        p.Pumps,
	}

	api.RegisterDevicesServer(grpcServer, srvr)
	grpcHealth.RegisterHealthServer(grpcServer, healthServer)
	grpcReflection.Register(grpcServer)

	return srvr
}

func (s *GrpcServer) GetPumpById(ctx context.Context, request *api.GetPumpByIdRequest) (*api.GetPumpByIdResponse, error) {
	pump, err := s.pumps.GetById(ctx, request.Id)
	if err != nil {
		log.Println(fmt.Sprintf("error while getting pump by id: %v", err))
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}
	if pump == nil {
		return nil, status.Errorf(codes.NotFound, "pump not found")
	}

	return &api.GetPumpByIdResponse{
		Pump: pump,
	}, nil
}

func (s *GrpcServer) ListPumps(ctx context.Context, request *api.ListPumpsRequest) (*api.ListPumpsResponse, error) {
	pumps, err := s.pumps.List(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error while getting list of pumps: %v", err))
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}

	return &api.ListPumpsResponse{
		Pumps: pumps,
	}, nil
}

func (s *GrpcServer) GetCgmById(ctx context.Context, request *api.GetCgmByIdRequest) (*api.GetCgmByIdResponse, error) {
	cgm, err := s.cgms.GetById(ctx, request.Id)
	if err != nil {
		log.Println(fmt.Sprintf("error while getting cgm by id: %v", err))
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}
	if cgm == nil {
		return nil, status.Errorf(codes.NotFound, "cgm not found")
	}

	return &api.GetCgmByIdResponse{
		Cgm: cgm,
	}, nil
}

func (s *GrpcServer) ListCgms(ctx context.Context, request *api.ListCgmsRequest) (*api.ListCgmsResponse, error) {
	cgms, err := s.cgms.List(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error while getting list of cgms: %v", err))
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}

	return &api.ListCgmsResponse{
		Cgms: cgms,
	}, nil
}

func (s *GrpcServer) Run(ctx context.Context, lis net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		<-ctx.Done()
		if err := s.Stop(); err != nil {
			log.Println(fmt.Sprintf("error while shutting down the grpc server: %v", err))
		}
	}()

	log.Println(fmt.Sprintf("serving grpc requests on %v", lis.Addr().String()))
	// blocks until the grpc server exits
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Println(fmt.Sprintf("failed to start grpc server: %v", err))
		return
	}

	log.Println("grpc server was successfully shutdown")
}

func (s *GrpcServer) Stop() error {
	s.SetNotServing()
	s.grpcServer.GracefulStop()
	return nil
}

func (s *GrpcServer) SetServing() {
	s.healthServer.SetServingStatus(serviceName, grpcHealth.HealthCheckResponse_SERVING)
}

func (s *GrpcServer) SetNotServing() {
	s.healthServer.SetServingStatus(serviceName, grpcHealth.HealthCheckResponse_NOT_SERVING)
}
