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
	Tags                 []*Tag   `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fbce36946609d703, []int{0}
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

func (m *User) GetTags() []*Tag {
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
	return fileDescriptor_users_fbce36946609d703, []int{1}
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
	return fileDescriptor_users_fbce36946609d703, []int{2}
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
	return fileDescriptor_users_fbce36946609d703, []int{3}
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
	return fileDescriptor_users_fbce36946609d703, []int{4}
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

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_fbce36946609d703) }

var fileDescriptor_users_fbce36946609d703 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xc6, 0x24, 0x30, 0x56, 0x10, 0x19, 0x84, 0x6a, 0x19, 0x24, 0x2c, 0x9f, 0x42,
	0xa5, 0xda, 0xe0, 0x72, 0x2a, 0x37, 0x72, 0xa8, 0x7a, 0x35, 0xf4, 0xc2, 0x25, 0x5a, 0xd7, 0x83,
	0xbb, 0x2a, 0xf1, 0x6e, 0x77, 0xd7, 0x45, 0x08, 0x71, 0xe1, 0x15, 0x78, 0x34, 0xae, 0x1c, 0x79,
	0x10, 0xb4, 0x1b, 0x93, 0x5a, 0x85, 0xba, 0xe9, 0xcd, 0xeb, 0x9d, 0xf9, 0xbf, 0xf9, 0xff, 0xd1,
	0x42, 0xd0, 0x6a, 0x52, 0x3a, 0x95, 0x4a, 0x18, 0x81, 0x51, 0x4d, 0x95, 0x50, 0xbc, 0x39, 0x6f,
	0x53, 0xa3, 0xdb, 0x9a, 0x57, 0xe2, 0x5c, 0x2c, 0x35, 0xa9, 0x4b, 0x52, 0xd1, 0xb3, 0x5a, 0x88,
	0xfa, 0x13, 0x65, 0x4c, 0xf2, 0x8c, 0x35, 0x8d, 0x30, 0xcc, 0x70, 0xd1, 0x74, 0x9d, 0xd1, 0xd3,
	0xee, 0xd6, 0x9d, 0xca, 0xf6, 0x63, 0x46, 0x2b, 0x69, 0xbe, 0x74, 0x97, 0x60, 0x58, 0xdd, 0x15,
	0x26, 0x67, 0xe0, 0x9f, 0x68, 0x52, 0xb8, 0x0b, 0x13, 0x4b, 0x5e, 0xf2, 0x2a, 0xf4, 0x62, 0x6f,
	0x3e, 0x2d, 0xc6, 0xf6, 0x78, 0x5c, 0x21, 0x82, 0xdf, 0xb0, 0x15, 0x85, 0x3b, 0xb1, 0x37, 0x7f,
	0x50, 0xb8, 0x6f, 0x3c, 0x00, 0xdf, 0x4a, 0x84, 0xa3, 0x78, 0x34, 0x0f, 0xf2, 0xe7, 0xe9, 0xcd,
	0x63, 0xa6, 0xef, 0x59, 0x5d, 0xb8, 0xe2, 0x64, 0x17, 0x9e, 0x1c, 0x91, 0x59, 0xb4, 0x4a, 0x51,
	0x63, 0x2c, 0xb3, 0xa0, 0x8b, 0x96, 0xb4, 0x49, 0x8e, 0x61, 0x76, 0x22, 0x2b, 0x66, 0xa8, 0xf7,
	0x13, 0x5f, 0x83, 0x6f, 0x07, 0x70, 0xc3, 0x04, 0x79, 0x3c, 0x84, 0x70, 0x6d, 0xae, 0x3a, 0x79,
	0x01, 0x0f, 0x8f, 0xa8, 0x2f, 0x7e, 0xa3, 0xaf, 0x64, 0x01, 0xb3, 0x85, 0xa2, 0x6b, 0xd4, 0xbf,
	0x66, 0xbd, 0x9e, 0xd9, 0x08, 0xee, 0x4b, 0xa6, 0xf5, 0x67, 0xa1, 0xaa, 0x2e, 0x84, 0xcd, 0x39,
	0xff, 0x35, 0x82, 0xc0, 0xf6, 0xbf, 0x23, 0x75, 0xc9, 0x4f, 0x09, 0x8d, 0xe3, 0xf7, 0x3c, 0xe2,
	0xab, 0xa1, 0xc9, 0xff, 0x9b, 0x47, 0x74, 0xab, 0xd9, 0x64, 0xfa, 0xfd, 0xe7, 0xef, 0x1f, 0x3b,
	0x13, 0xbc, 0x97, 0x59, 0x33, 0xa8, 0x00, 0xae, 0x02, 0xc4, 0xfd, 0xc1, 0xf6, 0xeb, 0x41, 0x6f,
	0x41, 0x7b, 0xec, 0x68, 0xd3, 0x7c, 0x4d, 0x3b, 0x74, 0x49, 0xe3, 0x05, 0x4c, 0xba, 0xa4, 0x71,
	0xef, 0x16, 0x8b, 0x77, 0xa3, 0x85, 0x8e, 0x86, 0xf8, 0xc8, 0xd1, 0x74, 0xf6, 0xb5, 0x5b, 0xdf,
	0x37, 0x94, 0x00, 0x57, 0x1b, 0x1b, 0xb6, 0xf9, 0xcf, 0x66, 0xb7, 0x00, 0xcf, 0x1c, 0x38, 0x48,
	0xc6, 0x6b, 0xf0, 0xa1, 0xb7, 0xf7, 0x36, 0xff, 0xf0, 0xb2, 0xe6, 0xe6, 0xac, 0x2d, 0xd3, 0x53,
	0xb1, 0xca, 0x36, 0x02, 0xd9, 0x46, 0x60, 0x7f, 0x2d, 0x60, 0x1f, 0xe1, 0x1b, 0x26, 0xf9, 0x52,
	0x96, 0xe5, 0xd8, 0xbd, 0xab, 0x83, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x8d, 0xf1, 0xaf,
	0xc9, 0x03, 0x00, 0x00,
}
