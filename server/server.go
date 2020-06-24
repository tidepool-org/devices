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
	"net"
	"sync"
)

const serviceName = "api.Devices"

type Params struct {
	Cfg   *config.Config
	Cgms  repo.CgmsRepo
	Pumps repo.PumpsRepo
}

type Server struct {
	api.UnimplementedDevicesServer
	grpcServer   *grpc.Server
	healthServer *health.Server

	cfg   *config.Config
	cgms  repo.CgmsRepo
	pumps repo.PumpsRepo
}

var _ api.DevicesServer = &Server{}

func New(p *Params) *Server {
	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	srvr := &Server{
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

func (s *Server) GetPumpById(ctx context.Context, request *api.GetPumpByIdRequest) (*api.GetPumpByIdResponse, error) {
	pump, err := s.pumps.GetById(ctx, request.Id)
	if err != nil {
		fmt.Printf("error while getting pump by id: %v", err)
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}
	if pump == nil {
		return nil, status.Errorf(codes.NotFound, "pump not found")
	}

	return &api.GetPumpByIdResponse{
		Pump: pump,
	}, nil
}

func (s *Server) ListPumps(ctx context.Context, request *api.ListPumpsRequest) (*api.ListPumpsResponse, error) {
	pumps, err := s.pumps.List(ctx)
	if err != nil {
		fmt.Printf("error while getting list of pumps: %v", err)
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}

	return &api.ListPumpsResponse{
		Pumps: pumps,
	}, nil
}

func (s *Server) GetCgmById(ctx context.Context, request *api.GetCgmByIdRequest) (*api.GetCgmByIdResponse, error) {
	cgm, err := s.cgms.GetById(ctx, request.Id)
	if err != nil {
		fmt.Printf("error while getting cgm by id: %v", err)
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}
	if cgm == nil {
		return nil, status.Errorf(codes.NotFound, "cgm not found")
	}

	return &api.GetCgmByIdResponse{
		Cgm: cgm,
	}, nil
}

func (s *Server) ListCgms(ctx context.Context, request *api.ListCgmsRequest) (*api.ListCgmsResponse, error) {
	cgms, err := s.cgms.List(ctx)
	if err != nil {
		fmt.Printf("error while getting list of cgms: %v", err)
		return nil, status.Errorf(codes.Internal, "unexpected error occurred")
	}

	return &api.ListCgmsResponse{
		Cgms: cgms,
	}, nil
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", s.cfg.ServerPort))
	if err != nil {
		fmt.Printf("error while opening listener: %v", err)
		return
	}

	go func() {
		select {
		case <-ctx.Done():
			if err := s.stop(); err != nil {
				fmt.Printf("error while shutting down the server: %v", err)
			}
		}
	}()

	s.healthServer.SetServingStatus(serviceName, grpcHealth.HealthCheckResponse_SERVING)
	fmt.Println(fmt.Sprintf("server started on port %v", s.cfg.ServerPort))

	// blocks until the grpc server exits
	if err := s.grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to start grpc server: %v", err)
		return
	}

	fmt.Printf("server was successfully shutdown")
}

func (s *Server) stop() error {
	s.healthServer.SetServingStatus(serviceName, grpcHealth.HealthCheckResponse_NOT_SERVING)
	s.grpcServer.GracefulStop()
	return nil
}
