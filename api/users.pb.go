// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package api_pb // import "github.com/gedorinku/tsugidoko-server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/gedorinku/tsugidoko-server/api/type"
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

type User struct {
	UserId               uint32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags                 []*_type.Tag `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_2200e63caaa36a6a, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetTags() []*_type.Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type GetCurrentUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserRequest) Reset()         { *m = GetCurrentUserRequest{} }
func (m *GetCurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserRequest) ProtoMessage()    {}
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_2200e63caaa36a6a, []int{1}
}
func (m *GetCurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetCurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserRequest.Merge(dst, src)
}
func (m *GetCurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserRequest.Size(m)
}
func (m *GetCurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserRequest proto.InternalMessageInfo

type UpdateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_2200e63caaa36a6a, []int{2}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_2200e63caaa36a6a, []int{3}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_2200e63caaa36a6a, []int{4}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "gedorinku.tsugidoko_server.User")
	proto.RegisterType((*GetCurrentUserRequest)(nil), "gedorinku.tsugidoko_server.GetCurrentUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "gedorinku.tsugidoko_server.UpdateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "gedorinku.tsugidoko_server.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "gedorinku.tsugidoko_server.CreateUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.UserService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentUser",
			Handler:    _UserService_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_2200e63caaa36a6a) }

var fileDescriptor_users_2200e63caaa36a6a = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xb5, 0xb4, 0xf0, 0xa2, 0x4e, 0xf4, 0x21, 0xb4, 0x2a, 0x70, 0xa8, 0x22, 0x0e,
	0x65, 0xd2, 0x62, 0x08, 0x1c, 0xd0, 0xb8, 0xd1, 0xc3, 0xb4, 0x6b, 0x61, 0x17, 0x2e, 0x95, 0xbb,
	0x3c, 0x8c, 0x35, 0x1a, 0x7b, 0xb6, 0x33, 0x34, 0x21, 0x2e, 0x7c, 0x05, 0x3e, 0x1a, 0x57, 0x8e,
	0x7c, 0x10, 0x64, 0x27, 0x74, 0xd1, 0x60, 0x59, 0xb9, 0xe5, 0xc5, 0xfe, 0xbf, 0xdf, 0xfb, 0xff,
	0x9f, 0x0c, 0x71, 0x65, 0xc9, 0xd8, 0x4c, 0x1b, 0xe5, 0x14, 0x26, 0x82, 0x0a, 0x65, 0x64, 0x79,
	0x56, 0x65, 0xce, 0x56, 0x42, 0x16, 0xea, 0x4c, 0x2d, 0x2d, 0x99, 0x0b, 0x32, 0xc9, 0x63, 0xa1,
	0x94, 0xf8, 0x44, 0x8c, 0x6b, 0xc9, 0x78, 0x59, 0x2a, 0xc7, 0x9d, 0x54, 0x65, 0xa3, 0x4c, 0x1e,
	0x35, 0xa7, 0xa1, 0x5a, 0x55, 0x1f, 0x18, 0xad, 0xb5, 0xbb, 0x6c, 0x0e, 0x77, 0xdd, 0xa5, 0x26,
	0xe6, 0xb8, 0xa8, 0xeb, 0x74, 0x0d, 0xfd, 0x13, 0x4b, 0x06, 0xf7, 0x60, 0xe8, 0xe9, 0x4b, 0x59,
	0x4c, 0xa2, 0x69, 0x34, 0x1b, 0x2d, 0x06, 0xbe, 0x3c, 0x2e, 0x10, 0xa1, 0x5f, 0xf2, 0x35, 0x4d,
	0x76, 0xa6, 0xd1, 0xec, 0xde, 0x22, 0x7c, 0xe3, 0x2b, 0xe8, 0x3b, 0x2e, 0xec, 0xa4, 0x37, 0xed,
	0xcd, 0xe2, 0xfc, 0x49, 0x76, 0xf3, 0xa8, 0x99, 0xc7, 0x65, 0xef, 0xb8, 0x58, 0x04, 0x45, 0xba,
	0x07, 0x0f, 0x8f, 0xc8, 0xcd, 0x2b, 0x63, 0xa8, 0x74, 0x1e, 0xbc, 0xa0, 0xf3, 0x8a, 0xac, 0x4b,
	0x8f, 0x61, 0x7c, 0xa2, 0x0b, 0xee, 0xa8, 0xf5, 0x13, 0x5f, 0x42, 0xdf, 0x4f, 0x11, 0x26, 0x8a,
	0xf3, 0x69, 0x17, 0x27, 0xc8, 0xc2, 0xed, 0xf4, 0x29, 0xec, 0x1e, 0x51, 0xbb, 0xf9, 0x8d, 0xe6,
	0xd2, 0x39, 0x8c, 0xe7, 0x86, 0xae, 0x51, 0xff, 0x38, 0x8e, 0x5a, 0x8e, 0x13, 0xb8, 0xab, 0xb9,
	0xb5, 0x9f, 0x95, 0x29, 0x9a, 0x24, 0x36, 0x75, 0xfe, 0xb3, 0x07, 0xb1, 0xd7, 0xbf, 0x25, 0x73,
	0x21, 0x4f, 0x09, 0x5d, 0xe0, 0xb7, 0x3c, 0xe2, 0xf3, 0xae, 0xc9, 0xff, 0x99, 0x47, 0x72, 0xab,
	0xd9, 0x74, 0xf4, 0xed, 0xc7, 0xaf, 0xef, 0x3b, 0x43, 0xbc, 0xc3, 0xbc, 0x19, 0x34, 0x00, 0x57,
	0x01, 0xe2, 0x41, 0xa7, 0xfc, 0x7a, 0xd0, 0x5b, 0xd0, 0x1e, 0x04, 0xda, 0x28, 0xaf, 0x69, 0x87,
	0x21, 0x69, 0x3c, 0x87, 0x61, 0x93, 0x34, 0xee, 0xdf, 0x62, 0xf1, 0xff, 0x68, 0x93, 0x40, 0x43,
	0xbc, 0x1f, 0x68, 0x96, 0x7d, 0x69, 0xd6, 0xf7, 0x15, 0x35, 0xc0, 0xd5, 0xc6, 0xba, 0x6d, 0xfe,
	0xb5, 0xd9, 0x2d, 0xc0, 0xe3, 0x00, 0x8e, 0xd3, 0x41, 0x0d, 0x3e, 0x8c, 0xf6, 0xdf, 0xe4, 0xef,
	0x9f, 0x09, 0xe9, 0x3e, 0x56, 0xab, 0xec, 0x54, 0xad, 0xd9, 0xa6, 0x01, 0xdb, 0x34, 0x38, 0xa8,
	0x1b, 0xf8, 0xd7, 0xf8, 0x9a, 0x6b, 0xb9, 0xd4, 0xab, 0xd5, 0x20, 0x3c, 0xae, 0x17, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x01, 0xea, 0x85, 0xd2, 0x03, 0x00, 0x00,
}
