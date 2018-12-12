package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// BeaconServiceServer is a composite interface of api_pb.BeaconServiceServer and grapiserver.Server.
type BeaconServiceServer interface {
	api_pb.BeaconServiceServer
	grapiserver.Server
}

// NewBeaconServiceServer creates a new BeaconServiceServer instance.
func NewBeaconServiceServer() BeaconServiceServer {
	return &beaconServiceServerImpl{}
}

type beaconServiceServerImpl struct {
}

func (s *beaconServiceServerImpl) ListBeacons(ctx context.Context, req *api_pb.ListBeaconsRequest) (*api_pb.ListBeaconsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *beaconServiceServerImpl) GetBeacon(ctx context.Context, req *api_pb.GetBeaconRequest) (*api_pb.Beacon, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *beaconServiceServerImpl) CreateBeacon(ctx context.Context, req *api_pb.CreateBeaconRequest) (*api_pb.Beacon, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *beaconServiceServerImpl) UpdateBeacon(ctx context.Context, req *api_pb.UpdateBeaconRequest) (*api_pb.Beacon, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *beaconServiceServerImpl) DeleteBeacon(ctx context.Context, req *api_pb.DeleteBeaconRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
