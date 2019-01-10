package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	type_pb "github.com/gedorinku/tsugidoko-server/api/type"
	"github.com/gedorinku/tsugidoko-server/app/conv"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// ClassRoomServiceServer is a composite interface of api_pb.ClassRoomServiceServer and grapiserver.Server.
type ClassRoomServiceServer interface {
	api_pb.ClassRoomServiceServer
	grapiserver.Server
}

// NewClassRoomServiceServer creates a new ClassRoomServiceServer instance.
func NewClassRoomServiceServer(store di.StoreComponent) ClassRoomServiceServer {
	return &classRoomServiceServerImpl{
		store,
	}
}

type classRoomServiceServerImpl struct {
	di.StoreComponent
}

func (s *classRoomServiceServerImpl) ListClassRooms(ctx context.Context, req *api_pb.ListClassRoomsRequest) (*api_pb.ListClassRoomsResponse, error) {
	cs := s.ClassRoomStore(ctx)
	rooms, err := cs.ListClassRoom(conv.Int32SliceToInt64Slice(req.GetTagIds()))
	if err != nil {
		if err == sql.ErrNoRows {
			return classRoomsToResponse(nil), nil
		}
		grpclog.Error(err)
		return nil, err
	}

	return classRoomsToResponse(rooms), nil
}

func (s *classRoomServiceServerImpl) GetClassRoom(ctx context.Context, req *api_pb.GetClassRoomRequest) (*api_pb.ClassRoom, error) {
	cs := s.ClassRoomStore(ctx)
	room, err := cs.GetClassRoom(int64(req.ClassRoomId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Class room not found")
		}
		grpclog.Error(err)
		return nil, err
	}

	return classRoomToResponse(room), nil
}

func (s *classRoomServiceServerImpl) CreateClassRoom(ctx context.Context, req *api_pb.CreateClassRoomRequest) (*api_pb.ClassRoom, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) UpdateClassRoom(ctx context.Context, req *api_pb.UpdateClassRoomRequest) (*api_pb.ClassRoom, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) DeleteClassRoom(ctx context.Context, req *api_pb.DeleteClassRoomRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func classRoomsToResponse(rooms []*record.ClassRoom) *api_pb.ListClassRoomsResponse {
	resp := make([]*api_pb.ClassRoom, 0, len(rooms))
	for _, r := range rooms {
		resp = append(resp, classRoomToResponse(r))
	}

	return &api_pb.ListClassRoomsResponse{
		ClassRooms: resp,
	}
}

func classRoomToResponse(room *record.ClassRoom) *api_pb.ClassRoom {
	resp := &api_pb.ClassRoom{
		ClassRoomId: int32(room.ID),
		Name:        room.Name,
		Latitude:    room.Latitude,
		Longitude:   room.Longitude,
		BuildingId:  int32(room.BuildingID),
		Floor:       int32(room.Floor),
		LocalX:      room.LocalX,
		LocalY:      room.LocalY,
	}
	if room.R != nil {
		resp.TagCounts = classRoomTagsToResponse(room.R.ClassRoomTags)
		resp.Beacons = beaconsToResponse(room.R.Beacons)
	}

	return resp
}

func beaconsToResponse(beacons []*record.Beacon) []*type_pb.Beacon {
	resp := make([]*type_pb.Beacon, 0, len(beacons))
	for _, b := range beacons {
		resp = append(resp, beaconToResponse(b))
	}

	return resp
}

func beaconToResponse(beacon *record.Beacon) *type_pb.Beacon {
	return &type_pb.Beacon{
		Bssid: beacon.Bssid,
	}
}

func classRoomTagsToResponse(roomTags []*record.ClassRoomTag) []*type_pb.TagCount {
	resp := make([]*type_pb.TagCount, 0, len(roomTags))
	for _, rt := range roomTags {
		if rt.R == nil || rt.R.Tag == nil {
			continue
		}
		tc := &type_pb.TagCount{
			Tag:   tagToResponse(rt.R.Tag),
			Count: int32(rt.Count),
		}
		resp = append(resp, tc)
	}

	return resp
}
