// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tags.proto

package api_pb // import "github.com/gedorinku/tsugidoko-server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
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

type Tag struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{0}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type TagCount struct {
	Tag                  *Tag     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagCount) Reset()         { *m = TagCount{} }
func (m *TagCount) String() string { return proto.CompactTextString(m) }
func (*TagCount) ProtoMessage()    {}
func (*TagCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{1}
}
func (m *TagCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagCount.Unmarshal(m, b)
}
func (m *TagCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagCount.Marshal(b, m, deterministic)
}
func (dst *TagCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagCount.Merge(dst, src)
}
func (m *TagCount) XXX_Size() int {
	return xxx_messageInfo_TagCount.Size(m)
}
func (m *TagCount) XXX_DiscardUnknown() {
	xxx_messageInfo_TagCount.DiscardUnknown(m)
}

var xxx_messageInfo_TagCount proto.InternalMessageInfo

func (m *TagCount) GetTag() *Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *TagCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ListTagsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTagsRequest) Reset()         { *m = ListTagsRequest{} }
func (m *ListTagsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTagsRequest) ProtoMessage()    {}
func (*ListTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{2}
}
func (m *ListTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTagsRequest.Unmarshal(m, b)
}
func (m *ListTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTagsRequest.Marshal(b, m, deterministic)
}
func (dst *ListTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTagsRequest.Merge(dst, src)
}
func (m *ListTagsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTagsRequest.Size(m)
}
func (m *ListTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTagsRequest proto.InternalMessageInfo

type ListTagsResponse struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTagsResponse) Reset()         { *m = ListTagsResponse{} }
func (m *ListTagsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTagsResponse) ProtoMessage()    {}
func (*ListTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{3}
}
func (m *ListTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTagsResponse.Unmarshal(m, b)
}
func (m *ListTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTagsResponse.Marshal(b, m, deterministic)
}
func (dst *ListTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTagsResponse.Merge(dst, src)
}
func (m *ListTagsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTagsResponse.Size(m)
}
func (m *ListTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTagsResponse proto.InternalMessageInfo

func (m *ListTagsResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type GetTagRequest struct {
	TagId                int32    `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagRequest) Reset()         { *m = GetTagRequest{} }
func (m *GetTagRequest) String() string { return proto.CompactTextString(m) }
func (*GetTagRequest) ProtoMessage()    {}
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{4}
}
func (m *GetTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagRequest.Unmarshal(m, b)
}
func (m *GetTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagRequest.Marshal(b, m, deterministic)
}
func (dst *GetTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagRequest.Merge(dst, src)
}
func (m *GetTagRequest) XXX_Size() int {
	return xxx_messageInfo_GetTagRequest.Size(m)
}
func (m *GetTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagRequest proto.InternalMessageInfo

func (m *GetTagRequest) GetTagId() int32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

type CreateTagRequest struct {
	Tag                  *Tag     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTagRequest) Reset()         { *m = CreateTagRequest{} }
func (m *CreateTagRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTagRequest) ProtoMessage()    {}
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_cd0cb669a6f70d77, []int{5}
}
func (m *CreateTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTagRequest.Unmarshal(m, b)
}
func (m *CreateTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTagRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTagRequest.Merge(dst, src)
}
func (m *CreateTagRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTagRequest.Size(m)
}
func (m *CreateTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTagRequest proto.InternalMessageInfo

func (m *CreateTagRequest) GetTag() *Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "gedorinku.tsugidoko_server.Tag")
	proto.RegisterType((*TagCount)(nil), "gedorinku.tsugidoko_server.TagCount")
	proto.RegisterType((*ListTagsRequest)(nil), "gedorinku.tsugidoko_server.ListTagsRequest")
	proto.RegisterType((*ListTagsResponse)(nil), "gedorinku.tsugidoko_server.ListTagsResponse")
	proto.RegisterType((*GetTagRequest)(nil), "gedorinku.tsugidoko_server.GetTagRequest")
	proto.RegisterType((*CreateTagRequest)(nil), "gedorinku.tsugidoko_server.CreateTagRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
}

type tagServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagServiceClient(cc *grpc.ClientConn) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.TagService/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.TagService/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.TagService/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	GetTag(context.Context, *GetTagRequest) (*Tag, error)
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.TagService/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.TagService/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.TagService/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTags",
			Handler:    _TagService_ListTags_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagService_GetTag_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _TagService_CreateTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tags.proto",
}

func init() { proto.RegisterFile("tags.proto", fileDescriptor_tags_cd0cb669a6f70d77) }

var fileDescriptor_tags_cd0cb669a6f70d77 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x15, 0xa7, 0x0e, 0xed, 0x40, 0x4b, 0x18, 0x01, 0x8a, 0x0c, 0x12, 0xd1, 0x1e, 0x50,
	0x11, 0x60, 0x43, 0x7a, 0x83, 0x03, 0x12, 0x15, 0xaa, 0x90, 0x38, 0xb9, 0x3e, 0x71, 0x89, 0x26,
	0xf1, 0xb2, 0xac, 0x9a, 0xec, 0x1a, 0xef, 0xb8, 0x12, 0x42, 0x5c, 0x78, 0x05, 0x1e, 0x8d, 0x17,
	0xe0, 0xc0, 0x83, 0xa0, 0xdd, 0x0d, 0x01, 0x2a, 0x91, 0x86, 0x9b, 0xc7, 0xfb, 0xef, 0xf7, 0xcf,
	0xff, 0xcb, 0x06, 0x60, 0x52, 0x2e, 0x6f, 0x5a, 0xcb, 0x16, 0x33, 0x25, 0x6b, 0xdb, 0x6a, 0x73,
	0xd6, 0xe5, 0xec, 0x3a, 0xa5, 0x6b, 0x7b, 0x66, 0xa7, 0x4e, 0xb6, 0xe7, 0xb2, 0xcd, 0xee, 0x2a,
	0x6b, 0xd5, 0x42, 0x16, 0xd4, 0xe8, 0x82, 0x8c, 0xb1, 0x4c, 0xac, 0xad, 0x59, 0xdd, 0xcc, 0xee,
	0xac, 0x4e, 0xc3, 0x34, 0xeb, 0xde, 0x15, 0x72, 0xd9, 0xf0, 0xc7, 0x78, 0x28, 0x5e, 0x40, 0xbf,
	0x22, 0x85, 0x07, 0x90, 0xe8, 0x7a, 0xd4, 0x1b, 0xf7, 0x0e, 0xd3, 0x32, 0xd1, 0x35, 0x22, 0xec,
	0x18, 0x5a, 0xca, 0x51, 0x32, 0xee, 0x1d, 0xee, 0x95, 0xe1, 0x19, 0x6f, 0x42, 0xca, 0x96, 0x69,
	0x31, 0xea, 0x07, 0x59, 0x1c, 0xc4, 0x29, 0xec, 0x56, 0xa4, 0x8e, 0x6d, 0x67, 0x18, 0x9f, 0x42,
	0x9f, 0x49, 0x05, 0xcc, 0xd5, 0xc9, 0xbd, 0xfc, 0xdf, 0x1b, 0xe7, 0x15, 0xa9, 0xd2, 0x6b, 0x3d,
	0x74, 0xee, 0xef, 0x06, 0xa7, 0xb4, 0x8c, 0x83, 0xb8, 0x01, 0xd7, 0xdf, 0x68, 0xc7, 0x15, 0x29,
	0x57, 0xca, 0x0f, 0x9d, 0x74, 0x2c, 0x4e, 0x60, 0xf8, 0xfb, 0x95, 0x6b, 0xac, 0x71, 0x12, 0x8f,
	0x60, 0xc7, 0x37, 0x34, 0xea, 0x8d, 0xfb, 0xdb, 0x18, 0x06, 0xb1, 0xb8, 0x0f, 0xfb, 0x27, 0xd2,
	0x73, 0x56, 0x64, 0xbc, 0x05, 0x03, 0x26, 0x35, 0x5d, 0xe7, 0x4f, 0x99, 0xd4, 0xeb, 0x5a, 0xbc,
	0x82, 0xe1, 0x71, 0x2b, 0x89, 0xe5, 0x1f, 0xd2, 0xff, 0x0f, 0x38, 0xf9, 0x9e, 0x00, 0x54, 0xa4,
	0x4e, 0x65, 0x7b, 0xae, 0xe7, 0x12, 0x19, 0x76, 0x7f, 0xc5, 0xc0, 0x87, 0x9b, 0x00, 0x17, 0xf2,
	0x67, 0x8f, 0xb6, 0x13, 0xc7, 0x66, 0xc4, 0xfe, 0x97, 0x6f, 0x3f, 0xbe, 0x26, 0x57, 0x30, 0x2d,
	0x7c, 0x66, 0x5c, 0xc0, 0x20, 0x66, 0xc6, 0x07, 0x9b, 0x30, 0x7f, 0xf5, 0x92, 0x5d, 0x96, 0x4f,
	0xdc, 0x0e, 0x26, 0x43, 0x3c, 0x08, 0x26, 0xc5, 0xa7, 0xd8, 0xe2, 0x67, 0x34, 0xb0, 0xb7, 0x6e,
	0x0e, 0x37, 0xee, 0x7d, 0xb1, 0xe0, 0xcb, 0x3d, 0x31, 0x78, 0x5e, 0x13, 0x31, 0xd8, 0x33, 0x5f,
	0xf1, 0xcb, 0xc9, 0xdb, 0x27, 0x4a, 0xf3, 0xfb, 0x6e, 0x96, 0xcf, 0xed, 0xb2, 0x58, 0x03, 0x8a,
	0x35, 0xe0, 0x71, 0x04, 0xf8, 0xff, 0xe3, 0x39, 0x35, 0x7a, 0xda, 0xcc, 0x66, 0x83, 0xf0, 0xf9,
	0x1f, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x71, 0x7f, 0xbc, 0xb5, 0x63, 0x03, 0x00, 0x00,
}
