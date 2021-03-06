// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_tags.proto

package api_pb // import "github.com/gedorinku/tsugidoko-server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateUserTagRequest struct {
	TagIds               []int32  `protobuf:"varint,1,rep,packed,name=tag_ids,json=tagIds,proto3" json:"tag_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserTagRequest) Reset()         { *m = UpdateUserTagRequest{} }
func (m *UpdateUserTagRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserTagRequest) ProtoMessage()    {}
func (*UpdateUserTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_tags_086cfe3208ecc60f, []int{0}
}
func (m *UpdateUserTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserTagRequest.Unmarshal(m, b)
}
func (m *UpdateUserTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserTagRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserTagRequest.Merge(dst, src)
}
func (m *UpdateUserTagRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserTagRequest.Size(m)
}
func (m *UpdateUserTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserTagRequest proto.InternalMessageInfo

func (m *UpdateUserTagRequest) GetTagIds() []int32 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type UpdateUserTagResponse struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserTagResponse) Reset()         { *m = UpdateUserTagResponse{} }
func (m *UpdateUserTagResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserTagResponse) ProtoMessage()    {}
func (*UpdateUserTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_tags_086cfe3208ecc60f, []int{1}
}
func (m *UpdateUserTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserTagResponse.Unmarshal(m, b)
}
func (m *UpdateUserTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserTagResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateUserTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserTagResponse.Merge(dst, src)
}
func (m *UpdateUserTagResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserTagResponse.Size(m)
}
func (m *UpdateUserTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserTagResponse proto.InternalMessageInfo

func (m *UpdateUserTagResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateUserTagRequest)(nil), "gedorinku.tsugidoko_server.UpdateUserTagRequest")
	proto.RegisterType((*UpdateUserTagResponse)(nil), "gedorinku.tsugidoko_server.UpdateUserTagResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserTagServiceClient is the client API for UserTagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserTagServiceClient interface {
	UpdateUserTag(ctx context.Context, in *UpdateUserTagRequest, opts ...grpc.CallOption) (*UpdateUserTagResponse, error)
}

type userTagServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserTagServiceClient(cc *grpc.ClientConn) UserTagServiceClient {
	return &userTagServiceClient{cc}
}

func (c *userTagServiceClient) UpdateUserTag(ctx context.Context, in *UpdateUserTagRequest, opts ...grpc.CallOption) (*UpdateUserTagResponse, error) {
	out := new(UpdateUserTagResponse)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserTagService/UpdateUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTagServiceServer is the server API for UserTagService service.
type UserTagServiceServer interface {
	UpdateUserTag(context.Context, *UpdateUserTagRequest) (*UpdateUserTagResponse, error)
}

func RegisterUserTagServiceServer(s *grpc.Server, srv UserTagServiceServer) {
	s.RegisterService(&_UserTagService_serviceDesc, srv)
}

func _UserTagService_UpdateUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTagServiceServer).UpdateUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.UserTagService/UpdateUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTagServiceServer).UpdateUserTag(ctx, req.(*UpdateUserTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserTagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.UserTagService",
	HandlerType: (*UserTagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserTag",
			Handler:    _UserTagService_UpdateUserTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_tags.proto",
}

func init() { proto.RegisterFile("user_tags.proto", fileDescriptor_user_tags_086cfe3208ecc60f) }

var fileDescriptor_user_tags_086cfe3208ecc60f = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0x87, 0x39, 0xff, 0x44, 0x58, 0xd1, 0xc0, 0x61, 0x50, 0x0e, 0x41, 0x49, 0x25, 0x82, 0xb7,
	0xf1, 0xd2, 0x69, 0x67, 0x27, 0x58, 0x45, 0xd3, 0xd8, 0x1c, 0x7b, 0xd9, 0x61, 0x5d, 0xa2, 0x3b,
	0xeb, 0xce, 0x6c, 0x1e, 0xc0, 0xd6, 0xd2, 0xde, 0x97, 0xf2, 0x15, 0x7c, 0x10, 0xd9, 0x4b, 0x88,
	0x28, 0x2a, 0xd8, 0xcf, 0xf7, 0xcd, 0xc7, 0x4f, 0x74, 0x23, 0x41, 0xa8, 0x59, 0x19, 0x2a, 0x7d,
	0x40, 0xc6, 0xbc, 0x30, 0xa0, 0x31, 0x58, 0x37, 0x8d, 0x25, 0x53, 0x34, 0x56, 0xe3, 0x14, 0x6b,
	0x82, 0x30, 0x83, 0x50, 0xec, 0x1b, 0x44, 0x73, 0x0f, 0x52, 0x79, 0x2b, 0x95, 0x73, 0xc8, 0x8a,
	0x2d, 0xba, 0x05, 0x59, 0x88, 0x4f, 0x4b, 0x5f, 0x8a, 0x9d, 0xb1, 0xd7, 0x8a, 0x61, 0x4c, 0x10,
	0x6e, 0x94, 0x19, 0xc1, 0x63, 0x04, 0xe2, 0x7c, 0x57, 0x6c, 0xb0, 0x32, 0xb5, 0xd5, 0xb4, 0x97,
	0x1d, 0xae, 0x1e, 0xad, 0x8f, 0x3a, 0xac, 0xcc, 0xa5, 0xa6, 0xfe, 0x95, 0xe8, 0x7d, 0x03, 0xc8,
	0xa3, 0x23, 0xc8, 0x87, 0x62, 0x2d, 0x79, 0xdb, 0xf3, 0xcd, 0xea, 0xa0, 0xfc, 0x3d, 0xaf, 0x4c,
	0x58, 0x7b, 0x5c, 0xbd, 0x66, 0x62, 0x7b, 0x21, 0xba, 0x86, 0x30, 0xb3, 0x13, 0xc8, 0x9f, 0x33,
	0xb1, 0xf5, 0xe5, 0x43, 0x3e, 0xf8, 0xcb, 0xf5, 0x53, 0x7d, 0x71, 0xfa, 0x0f, 0x62, 0x9e, 0xdf,
	0xef, 0x3d, 0xbd, 0xbd, 0xbf, 0xac, 0x74, 0x2b, 0x21, 0xd3, 0xd0, 0x32, 0xd5, 0x9d, 0x65, 0xc7,
	0x17, 0xd5, 0xed, 0xc0, 0x58, 0xbe, 0x8b, 0x4d, 0x39, 0xc1, 0x07, 0xb9, 0xb4, 0xca, 0xa5, 0xf5,
	0x64, 0x6e, 0x4d, 0x53, 0x9f, 0x2b, 0x6f, 0x6b, 0xdf, 0x34, 0x9d, 0x76, 0xda, 0xe1, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0x93, 0xd5, 0x03, 0xb3, 0x01, 0x00, 0x00,
}
