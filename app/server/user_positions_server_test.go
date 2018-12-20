package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

func Test_UserPositionServiceServer_ListUserPositions(t *testing.T) {
	svr := NewUserPositionServiceServer()

	ctx := context.Background()
	req := &api_pb.ListUserPositionsRequest{}

	resp, err := svr.ListUserPositions(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_UserPositionServiceServer_GetUserPosition(t *testing.T) {
	svr := NewUserPositionServiceServer()

	ctx := context.Background()
	req := &api_pb.GetUserPositionRequest{}

	resp, err := svr.GetUserPosition(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_UserPositionServiceServer_CreateUserPosition(t *testing.T) {
	svr := NewUserPositionServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateUserPositionRequest{}

	resp, err := svr.CreateUserPosition(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_UserPositionServiceServer_UpdateUserPosition(t *testing.T) {
	svr := NewUserPositionServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateUserPositionRequest{}

	resp, err := svr.UpdateUserPosition(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_UserPositionServiceServer_DeleteUserPosition(t *testing.T) {
	svr := NewUserPositionServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteUserPositionRequest{}

	resp, err := svr.DeleteUserPosition(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
