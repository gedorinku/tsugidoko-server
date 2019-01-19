package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

func Test_TagServiceServer_ListTags(t *testing.T) {
	svr := NewTagServiceServer()

	ctx := context.Background()
	req := &api_pb.ListTagsRequest{}

	resp, err := svr.ListTags(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TagServiceServer_GetTag(t *testing.T) {
	svr := NewTagServiceServer()

	ctx := context.Background()
	req := &api_pb.GetTagRequest{}

	resp, err := svr.GetTag(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TagServiceServer_CreateTag(t *testing.T) {
	svr := NewTagServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateTagRequest{}

	resp, err := svr.CreateTag(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TagServiceServer_UpdateTag(t *testing.T) {
	svr := NewTagServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateTagRequest{}

	resp, err := svr.UpdateTag(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TagServiceServer_DeleteTag(t *testing.T) {
	svr := NewTagServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteTagRequest{}

	resp, err := svr.DeleteTag(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
