package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
)

// AdminServiceServer is a composite interface of api_pb.AdminServiceServer and grapiserver.Server.
type AdminServiceServer interface {
	api_pb.AdminServiceServer
	grapiserver.Server
}

// NewAdminServiceServer creates a new AdminServiceServer instance.
func NewAdminServiceServer(store di.StoreComponent) AdminServiceServer {
	return &adminServiceServerImpl{
		StoreComponent: store,
	}
}

type adminServiceServerImpl struct {
	di.StoreComponent
}

func (s *adminServiceServerImpl) ImportBuildings(ctx context.Context, req *api_pb.ImportBuildingsRequest) (*api_pb.ImportBuildingsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
