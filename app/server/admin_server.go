package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/app/di"
)

// AdminServiceServer is a composite interface of api_pb.AdminServiceServer and grapiserver.Server.
type AdminServiceServer interface {
	api_pb.AdminServiceServer
	grapiserver.Server
}

// NewAdminServiceServer creates a new AdminServiceServer instance.
func NewAdminServiceServer(cfg *config.Config, store di.StoreComponent, cli di.ClientComponent) AdminServiceServer {
	return &adminServiceServerImpl{
		StoreComponent:  store,
		ClientComponent: cli,
		adminSecret:     cfg.AdminSecret,
	}
}

type adminServiceServerImpl struct {
	di.StoreComponent
	di.ClientComponent
	adminSecret string
}

func (s *adminServiceServerImpl) ImportBuildings(ctx context.Context, req *api_pb.ImportBuildingsRequest) (*api_pb.ImportBuildingsResponse, error) {
	if req.GetKey() != s.adminSecret {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	cli := s.GASClient(ctx)
	imp, err := cli.GetBuildings()
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	as := s.AdminStore(ctx)
	err = as.Overwrite(imp.Buildings, imp.ClassRooms, imp.Beacons)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	us := s.UserPositionStore(ctx)
	err = us.ResetUserPosition()
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return &api_pb.ImportBuildingsResponse{
		Message: "import from google spreadsheet was successfully finished!",
	}, nil
}
