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
	"github.com/gedorinku/tsugidoko-server/model"
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
			return classRoomsToResponse(nil, nil), nil
		}
		grpclog.Error(err)
		return nil, err
	}

	ts := s.TagStore(ctx)
	tags, err := ts.TagsMap()
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return classRoomsToResponse(rooms, tags), nil
}

func (s *classRoomServiceServerImpl) GetClassRoom(ctx context.Context, req *api_pb.GetClassRoomRequest) (*api_pb.ClassRoom, error) {
	cs := s.ClassRoomStore(ctx)
	room, err := cs.GetClassRoom(int64(req.ClassRoomId), conv.Int32SliceToInt64Slice(req.GetTagIds()))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Class room not found")
		}
		grpclog.Error(err)
		return nil, err
	}

	ts := s.TagStore(ctx)
	tags, err := ts.TagsMap()
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return classRoomToResponse(room, tags), nil
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

func classRoomsToResponse(rooms []*record.ClassRoom, tags map[int64]*model.Tag) *api_pb.ListClassRoomsResponse {
	resp := make([]*api_pb.ClassRoom, 0, len(rooms))
	for _, r := range rooms {
		resp = append(resp, classRoomToResponse(r, tags))
	}

	return &api_pb.ListClassRoomsResponse{
		ClassRooms: resp,
	}
}

func classRoomToResponse(room *record.ClassRoom, tags map[int64]*model.Tag) *api_pb.ClassRoom {
	resp := &api_pb.ClassRoom{
		ClassRoomId: int32(room.ID),
		Name:        room.Name,
		Floor:       int32(room.Floor),
		LocalX:      room.LocalX,
		LocalY:      room.LocalY,
	}
	if r := room.R; r != nil {
		resp.TagCounts = classRoomTagsToResponse(r.ClassRoomTags, tags)
		resp.Beacons = beaconsToResponse(r.Beacons)
		resp.Building = buildingToResponse(r.Building)
	}

	return resp
}

func buildingToResponse(building *record.Building) *type_pb.Building {
	if building == nil {
		return nil
	}
	return &type_pb.Building{
		Id:        int32(building.ID),
		Name:      building.Name,
		Latitude:  building.Latitude,
		Longitude: building.Longitude,
	}
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

func classRoomTagsToResponse(roomTags []*record.ClassRoomTag, tags map[int64]*model.Tag) []*api_pb.TagCount {
	resp := make([]*api_pb.TagCount, 0, len(roomTags))
	for _, rt := range roomTags {
		if rt.R == nil || rt.R.Tag == nil {
			continue
		}
		r := rt.R
		tc := &api_pb.TagCount{
			Tag:   tagToResponse(tags[r.Tag.ID]),
			Count: int32(rt.Count),
		}
		resp = append(resp, tc)
	}

	return resp
}
