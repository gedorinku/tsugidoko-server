// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sessions.proto

package api_pb // import "github.com/gedorinku/tsugidoko-server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	UesrId               uint32   `protobuf:"varint,2,opt,name=uesr_id,json=uesrId,proto3" json:"uesr_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ba0c6d694e9f51d5, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetUesrId() uint32 {
	if m != nil {
		return m.UesrId
	}
	return 0
}

type GetSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ba0c6d694e9f51d5, []int{1}
}
func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (dst *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(dst, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

func (m *GetSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type CreateSessionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ba0c6d694e9f51d5, []int{2}
}
func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(dst, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionRequest) Reset()         { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()    {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ba0c6d694e9f51d5, []int{3}
}
func (m *DeleteSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionRequest.Unmarshal(m, b)
}
func (m *DeleteSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionRequest.Merge(dst, src)
}
func (m *DeleteSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionRequest.Size(m)
}
func (m *DeleteSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionRequest proto.InternalMessageInfo

func (m *DeleteSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "gedorinku.tsugidoko_server.Session")
	proto.RegisterType((*GetSessionRequest)(nil), "gedorinku.tsugidoko_server.GetSessionRequest")
	proto.RegisterType((*CreateSessionRequest)(nil), "gedorinku.tsugidoko_server.CreateSessionRequest")
	proto.RegisterType((*DeleteSessionRequest)(nil), "gedorinku.tsugidoko_server.DeleteSessionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.SessionService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*Session, error)
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*empty.Empty, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sessions.proto",
}

func init() { proto.RegisterFile("sessions.proto", fileDescriptor_sessions_ba0c6d694e9f51d5) }

var fileDescriptor_sessions_ba0c6d694e9f51d5 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4b, 0x6b, 0xea, 0x40,
	0x14, 0xc7, 0x89, 0xf7, 0xa2, 0x37, 0x07, 0x14, 0xee, 0x20, 0x5e, 0xc9, 0x7d, 0x20, 0xb9, 0x1b,
	0x11, 0x9c, 0x88, 0xa5, 0x9b, 0x76, 0xd5, 0x37, 0x6e, 0xe3, 0xae, 0x1b, 0x49, 0xcc, 0x69, 0x3a,
	0xa8, 0x99, 0x74, 0x66, 0x62, 0x91, 0xb6, 0x50, 0xba, 0xee, 0xae, 0x1f, 0xad, 0x5f, 0xa1, 0x1f,
	0xa4, 0xe4, 0x61, 0xc4, 0xaa, 0x69, 0x77, 0x39, 0x39, 0x8f, 0xdf, 0x39, 0xff, 0xf9, 0x43, 0x4d,
	0xa2, 0x94, 0x8c, 0x07, 0x92, 0x86, 0x82, 0x2b, 0x4e, 0x0c, 0x1f, 0x3d, 0x2e, 0x58, 0x30, 0x89,
	0xa8, 0x92, 0x91, 0xcf, 0x3c, 0x3e, 0xe1, 0x23, 0x89, 0x62, 0x8e, 0xc2, 0xf8, 0xe3, 0x73, 0xee,
	0x4f, 0xd1, 0x72, 0x42, 0x66, 0x39, 0x41, 0xc0, 0x95, 0xa3, 0x56, 0x9d, 0xc6, 0xef, 0x2c, 0x9b,
	0x44, 0x6e, 0x74, 0x65, 0xe1, 0x2c, 0x54, 0x8b, 0x34, 0x69, 0x1e, 0x41, 0x65, 0x98, 0x82, 0xc8,
	0x5f, 0x80, 0x8c, 0x39, 0x62, 0x5e, 0x53, 0x6b, 0x69, 0x6d, 0xdd, 0xd6, 0xb3, 0x3f, 0x03, 0x8f,
	0xfc, 0x82, 0x4a, 0x84, 0x52, 0xc4, 0xb9, 0x52, 0x4b, 0x6b, 0x57, 0xed, 0x72, 0x1c, 0x0e, 0x3c,
	0xb3, 0x0f, 0x3f, 0x2f, 0x50, 0x65, 0x53, 0x6c, 0xbc, 0x89, 0x50, 0xaa, 0x4f, 0x86, 0x99, 0xe7,
	0x50, 0x3f, 0x11, 0xe8, 0x28, 0xfc, 0xd0, 0x46, 0xe0, 0x7b, 0xe0, 0xcc, 0x30, 0x6b, 0x48, 0xbe,
	0x89, 0x01, 0x3f, 0x42, 0x47, 0xca, 0x5b, 0x2e, 0x52, 0xb2, 0x6e, 0xe7, 0xb1, 0xb9, 0x0f, 0xf5,
	0x53, 0x9c, 0xe2, 0xc6, 0x9c, 0x62, 0x7c, 0xff, 0xf9, 0x1b, 0xd4, 0xb2, 0x8e, 0x21, 0x8a, 0x39,
	0x1b, 0x23, 0x79, 0xd4, 0x00, 0x56, 0x67, 0x90, 0x2e, 0xdd, 0xad, 0x37, 0xdd, 0x38, 0xd7, 0xf8,
	0x5f, 0x54, 0x9e, 0xd5, 0x9a, 0xff, 0x9e, 0x5e, 0xdf, 0x5e, 0x4a, 0x4d, 0xd2, 0xb0, 0x96, 0x6f,
	0x6b, 0xdd, 0xad, 0xb6, 0x7c, 0x20, 0xf7, 0x50, 0x5d, 0x13, 0x85, 0xf4, 0x8a, 0xa6, 0x6e, 0xd3,
	0xef, 0x6b, 0x7b, 0xd4, 0x93, 0x3d, 0x6a, 0xa6, 0x9e, 0xef, 0x71, 0xa0, 0x75, 0xc8, 0x02, 0xaa,
	0x6b, 0x52, 0x16, 0xd3, 0xb7, 0xa9, 0x6e, 0x34, 0x68, 0x6a, 0x35, 0xba, 0xb4, 0x1a, 0x3d, 0x8b,
	0xad, 0xb6, 0x3c, 0xbc, 0xb3, 0xe3, 0xf0, 0xe3, 0xfe, 0x65, 0xcf, 0x67, 0xea, 0x3a, 0x72, 0xe9,
	0x98, 0xcf, 0xac, 0x9c, 0x6a, 0xe5, 0xd4, 0x6e, 0x4a, 0x8d, 0x0d, 0x7e, 0xe8, 0x84, 0x6c, 0x14,
	0xba, 0x6e, 0x39, 0x61, 0xec, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x64, 0x00, 0x30, 0x21, 0x28,
	0x03, 0x00, 0x00,
}
