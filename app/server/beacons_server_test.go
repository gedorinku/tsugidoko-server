package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

func Test_BeaconServiceServer_ListBeacons(t *testing.T) {
	svr := NewBeaconServiceServer()

	ctx := context.Background()
	req := &api_pb.ListBeaconsRequest{}

	resp, err := svr.ListBeacons(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BeaconServiceServer_GetBeacon(t *testing.T) {
	svr := NewBeaconServiceServer()

	ctx := context.Background()
	req := &api_pb.GetBeaconRequest{}

	resp, err := svr.GetBeacon(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BeaconServiceServer_CreateBeacon(t *testing.T) {
	svr := NewBeaconServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateBeaconRequest{}

	resp, err := svr.CreateBeacon(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BeaconServiceServer_UpdateBeacon(t *testing.T) {
	svr := NewBeaconServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateBeaconRequest{}

	resp, err := svr.UpdateBeacon(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_BeaconServiceServer_DeleteBeacon(t *testing.T) {
	svr := NewBeaconServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteBeaconRequest{}

	resp, err := svr.DeleteBeacon(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
