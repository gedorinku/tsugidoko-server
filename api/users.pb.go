// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

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

type User struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_c35bb4c972e73530, []int{0}
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

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_c35bb4c972e73530, []int{1}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (dst *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(dst, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_c35bb4c972e73530, []int{2}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (dst *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(dst, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
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
	return fileDescriptor_users_c35bb4c972e73530, []int{3}
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
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_c35bb4c972e73530, []int{4}
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

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

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
	return fileDescriptor_users_c35bb4c972e73530, []int{5}
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

func init() {
	proto.RegisterType((*User)(nil), "gedorinku.tsugidoko_server.User")
	proto.RegisterType((*ListUsersRequest)(nil), "gedorinku.tsugidoko_server.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "gedorinku.tsugidoko_server.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "gedorinku.tsugidoko_server.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "gedorinku.tsugidoko_server.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "gedorinku.tsugidoko_server.UpdateUserRequest")
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
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
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

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
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

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_c35bb4c972e73530) }

var fileDescriptor_users_c35bb4c972e73530 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x4b, 0xe3, 0x40,
	0x14, 0xc7, 0x49, 0xb7, 0xdb, 0xb2, 0xaf, 0x6c, 0x69, 0x87, 0x5d, 0x36, 0x64, 0x77, 0x21, 0x64,
	0x2f, 0x5d, 0xa1, 0x89, 0xa4, 0xe2, 0x41, 0x6f, 0x7a, 0x90, 0xa2, 0xa7, 0x48, 0x2f, 0x5e, 0x4a,
	0xd2, 0x3c, 0xe3, 0x50, 0x9b, 0x49, 0x67, 0x26, 0x05, 0x11, 0x2f, 0xfd, 0x0a, 0x7e, 0x34, 0xbf,
	0x82, 0x9f, 0xc2, 0x93, 0xcc, 0x24, 0xb6, 0x95, 0xd2, 0xda, 0x83, 0xb7, 0xbc, 0xbc, 0xf7, 0xfe,
	0xbf, 0xf9, 0xff, 0xe1, 0x41, 0x23, 0x17, 0xc8, 0x85, 0x9b, 0x71, 0x26, 0x19, 0xb1, 0x12, 0x8c,
	0x19, 0xa7, 0xe9, 0x38, 0x77, 0xa5, 0xc8, 0x13, 0x1a, 0xb3, 0x31, 0x1b, 0x0a, 0xe4, 0x33, 0xe4,
	0xd6, 0x9f, 0x84, 0xb1, 0xe4, 0x16, 0xbd, 0x30, 0xa3, 0x5e, 0x98, 0xa6, 0x4c, 0x86, 0x92, 0xb2,
	0xb4, 0xdc, 0xb4, 0x7e, 0x97, 0x5d, 0x5d, 0x45, 0xf9, 0xb5, 0x87, 0x93, 0x4c, 0xde, 0x15, 0x4d,
	0xa7, 0x07, 0xd5, 0x81, 0x40, 0x4e, 0x7e, 0x41, 0x5d, 0xd1, 0x86, 0x34, 0x36, 0x0d, 0xdb, 0xe8,
	0x7c, 0x0f, 0x6a, 0xaa, 0xec, 0xc7, 0x84, 0x40, 0x35, 0x0d, 0x27, 0x68, 0x56, 0x6c, 0xa3, 0xf3,
	0x2d, 0xd0, 0xdf, 0x0e, 0x81, 0xd6, 0x05, 0x15, 0x52, 0x2d, 0x8a, 0x00, 0xa7, 0x39, 0x0a, 0xe9,
	0x9c, 0x43, 0x7b, 0xe5, 0x9f, 0xc8, 0x58, 0x2a, 0x90, 0x1c, 0xc2, 0x57, 0xed, 0xc1, 0x34, 0xec,
	0x2f, 0x9d, 0x86, 0x6f, 0xbb, 0x9b, 0x4d, 0xb8, 0x6a, 0x33, 0x28, 0xc6, 0x9d, 0xff, 0xd0, 0x3c,
	0x43, 0xad, 0x55, 0xca, 0x6f, 0x7c, 0x9f, 0xd3, 0x87, 0xf6, 0x29, 0xc7, 0x50, 0xe2, 0xea, 0xf4,
	0x01, 0x54, 0x55, 0x5b, 0x8f, 0xee, 0x82, 0xd5, 0xd3, 0x4a, 0x6a, 0x90, 0xc5, 0x9f, 0x21, 0xe5,
	0xbf, 0x54, 0xa0, 0xa1, 0xca, 0x4b, 0xe4, 0x33, 0x3a, 0x42, 0x32, 0x85, 0x7a, 0x69, 0x88, 0xec,
	0x6d, 0x93, 0x78, 0xef, 0xda, 0xfa, 0x10, 0xe7, 0x98, 0xf3, 0xa7, 0xe7, 0xc7, 0x0a, 0x21, 0x2d,
	0x4f, 0x27, 0xe7, 0xdd, 0x97, 0x29, 0x3d, 0x10, 0x01, 0xb0, 0x0c, 0x86, 0x74, 0xb7, 0x29, 0xad,
	0x05, 0xb8, 0x03, 0xf8, 0x87, 0x06, 0x37, 0x9d, 0x5a, 0x01, 0x3e, 0xd2, 0xbe, 0xc9, 0xdc, 0x00,
	0x58, 0x66, 0xb8, 0x9d, 0xba, 0x96, 0xf5, 0x0e, 0xd4, 0x7f, 0x9a, 0xfa, 0xd7, 0xff, 0xb9, 0x6a,
	0xd7, 0x7d, 0xf3, 0x5c, 0x3c, 0xe2, 0xc4, 0xbf, 0xda, 0x4f, 0xa8, 0xbc, 0xc9, 0x23, 0x77, 0xc4,
	0x26, 0xde, 0x42, 0xd2, 0x5b, 0x48, 0x76, 0x0b, 0x49, 0x75, 0x2f, 0xc7, 0x61, 0x46, 0x87, 0x59,
	0x14, 0xd5, 0xf4, 0x39, 0xf4, 0x5e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xff, 0xf4, 0x6e, 0xff, 0x74,
	0x03, 0x00, 0x00,
}
