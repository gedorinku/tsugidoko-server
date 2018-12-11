package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

func Test_ClassRoomServiceServer_ListClassRooms(t *testing.T) {
	svr := NewClassRoomServiceServer()

	ctx := context.Background()
	req := &api_pb.ListClassRoomsRequest{}

	resp, err := svr.ListClassRooms(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_ClassRoomServiceServer_GetClassRoom(t *testing.T) {
	svr := NewClassRoomServiceServer()

	ctx := context.Background()
	req := &api_pb.GetClassRoomRequest{}

	resp, err := svr.GetClassRoom(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_ClassRoomServiceServer_CreateClassRoom(t *testing.T) {
	svr := NewClassRoomServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateClassRoomRequest{}

	resp, err := svr.CreateClassRoom(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_ClassRoomServiceServer_UpdateClassRoom(t *testing.T) {
	svr := NewClassRoomServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateClassRoomRequest{}

	resp, err := svr.UpdateClassRoom(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_ClassRoomServiceServer_DeleteClassRoom(t *testing.T) {
	svr := NewClassRoomServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteClassRoomRequest{}

	resp, err := svr.DeleteClassRoom(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
