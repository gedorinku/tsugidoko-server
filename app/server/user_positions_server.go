package server

import (
	"context"
	"database/sql"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserPositionServiceServer is a composite interface of api_pb.UserPositionServiceServer and grapiserver.Server.
type UserPositionServiceServer interface {
	api_pb.UserPositionServiceServer
	grapiserver.Server
}

// NewUserPositionServiceServer creates a new UserPositionServiceServer instance.
func NewUserPositionServiceServer(store di.StoreComponent) UserPositionServiceServer {
	return &userPositionServiceServerImpl{
		StoreComponent: store,
	}
}

type userPositionServiceServerImpl struct {
	di.StoreComponent
}

func (s *userPositionServiceServerImpl) GetUserPosition(ctx context.Context, req *api_pb.GetUserPositionRequest) (*api_pb.UserPosition, error) {
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	us := s.UserPositionStore(ctx)
	pos, err := us.GetUserPosition(model.UserID(session.ID))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return userPositionToResponse(nil), nil
		}
		grpclog.Error(err)
		return nil, err
	}

	return userPositionToResponse(pos), nil
}

func (s *userPositionServiceServerImpl) UpdateUserPosition(ctx context.Context, req *api_pb.UpdateUserPositionRequest) (*api_pb.UserPosition, error) {
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	us := s.UserPositionStore(ctx)
	pos, err := us.UpdateUserPosition(model.UserID(session.ID), req.Bssid, req.GetIsValid())
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return userPositionToResponse(nil), status.Error(codes.NotFound, err.Error())
		}
		grpclog.Errorf("%+v", err)
		return nil, err
	}

	return userPositionToResponse(pos), nil
}

func userPositionToResponse(pos *record.UserPosition) *api_pb.UserPosition {
	var room *api_pb.ClassRoom
	if pos != nil && pos.R.ClassRoom != nil {
		room = classRoomToResponse(pos.R.ClassRoom)
	}
	resp := &api_pb.UserPosition{
		ClassRoom: room,
		IsValid:   room != nil,
	}

	return resp
}
