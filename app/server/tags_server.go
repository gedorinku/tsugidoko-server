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

// TagServiceServer is a composite interface of api_pb.TagServiceServer and grapiserver.Server.
type TagServiceServer interface {
	api_pb.TagServiceServer
	grapiserver.Server
}

// NewTagServiceServer creates a new TagServiceServer instance.
func NewTagServiceServer(store di.StoreComponent) TagServiceServer {
	return &tagServiceServerImpl{
		StoreComponent: store,
	}
}

type tagServiceServerImpl struct {
	di.StoreComponent
}

func (s *tagServiceServerImpl) ListTags(ctx context.Context, req *api_pb.ListTagsRequest) (*api_pb.ListTagsResponse, error) {
	ts := s.TagStore(ctx)
	tags, err := ts.ListValidTags()
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		grpclog.Error(err)
		return nil, err
	}

	resp := &api_pb.ListTagsResponse{
		Tags: tagsToResponse(tags),
	}
	return resp, nil
}

func (s *tagServiceServerImpl) GetTag(ctx context.Context, req *api_pb.GetTagRequest) (*api_pb.Tag, error) {
	ts := s.TagStore(ctx)
	t, err := ts.GetTag(int64(req.TagId))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "tag not found")
		}
		grpclog.Error(err)
		return nil, err
	}

	return tagToResponse(t), nil
}

func (s *tagServiceServerImpl) CreateTag(ctx context.Context, req *api_pb.CreateTagRequest) (*api_pb.Tag, error) {
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	ts := s.TagStore(ctx)
	t := &record.Tag{
		Name: req.Tag.Name,
	}
	t, err := ts.CreateTag(t)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	uts := s.UserTagStore(ctx)
	err = uts.AddUserTag(model.UserID(session.UserID), t.ID)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return tagToResponse(t), nil
}

func tagsToResponse(tags []*record.Tag) []*api_pb.Tag {
	res := make([]*api_pb.Tag, 0, len(tags))
	for _, t := range tags {
		res = append(res, tagToResponse(t))
	}
	return res
}

func tagToResponse(tag *record.Tag) *api_pb.Tag {
	return &api_pb.Tag{
		Id:   int32(tag.ID),
		Name: tag.Name,
	}
}
