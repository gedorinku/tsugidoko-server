// Code generated by protoc-gen-go. DO NOT EDIT.
// source: class_rooms.proto

package api_pb // import "github.com/gedorinku/tsugidoko-server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/gedorinku/tsugidoko-server/api/type"
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

type ClassRoom struct {
	ClassRoomId          int32           `protobuf:"varint,1,opt,name=class_room_id,json=classRoomId,proto3" json:"class_room_id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TagCounts            []*TagCount     `protobuf:"bytes,5,rep,name=tag_counts,json=tagCounts,proto3" json:"tag_counts,omitempty"`
	Beacons              []*_type.Beacon `protobuf:"bytes,7,rep,name=beacons,proto3" json:"beacons,omitempty"`
	Floor                int32           `protobuf:"varint,8,opt,name=floor,proto3" json:"floor,omitempty"`
	LocalX               float64         `protobuf:"fixed64,9,opt,name=local_x,json=localX,proto3" json:"local_x,omitempty"`
	LocalY               float64         `protobuf:"fixed64,10,opt,name=local_y,json=localY,proto3" json:"local_y,omitempty"`
	Building             *_type.Building `protobuf:"bytes,11,opt,name=building,proto3" json:"building,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClassRoom) Reset()         { *m = ClassRoom{} }
func (m *ClassRoom) String() string { return proto.CompactTextString(m) }
func (*ClassRoom) ProtoMessage()    {}
func (*ClassRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{0}
}
func (m *ClassRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassRoom.Unmarshal(m, b)
}
func (m *ClassRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassRoom.Marshal(b, m, deterministic)
}
func (dst *ClassRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassRoom.Merge(dst, src)
}
func (m *ClassRoom) XXX_Size() int {
	return xxx_messageInfo_ClassRoom.Size(m)
}
func (m *ClassRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ClassRoom proto.InternalMessageInfo

func (m *ClassRoom) GetClassRoomId() int32 {
	if m != nil {
		return m.ClassRoomId
	}
	return 0
}

func (m *ClassRoom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClassRoom) GetTagCounts() []*TagCount {
	if m != nil {
		return m.TagCounts
	}
	return nil
}

func (m *ClassRoom) GetBeacons() []*_type.Beacon {
	if m != nil {
		return m.Beacons
	}
	return nil
}

func (m *ClassRoom) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

func (m *ClassRoom) GetLocalX() float64 {
	if m != nil {
		return m.LocalX
	}
	return 0
}

func (m *ClassRoom) GetLocalY() float64 {
	if m != nil {
		return m.LocalY
	}
	return 0
}

func (m *ClassRoom) GetBuilding() *_type.Building {
	if m != nil {
		return m.Building
	}
	return nil
}

type ListClassRoomsRequest struct {
	TagIds               []int32  `protobuf:"varint,1,rep,packed,name=tag_ids,json=tagIds,proto3" json:"tag_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClassRoomsRequest) Reset()         { *m = ListClassRoomsRequest{} }
func (m *ListClassRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*ListClassRoomsRequest) ProtoMessage()    {}
func (*ListClassRoomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{1}
}
func (m *ListClassRoomsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClassRoomsRequest.Unmarshal(m, b)
}
func (m *ListClassRoomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClassRoomsRequest.Marshal(b, m, deterministic)
}
func (dst *ListClassRoomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClassRoomsRequest.Merge(dst, src)
}
func (m *ListClassRoomsRequest) XXX_Size() int {
	return xxx_messageInfo_ListClassRoomsRequest.Size(m)
}
func (m *ListClassRoomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClassRoomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListClassRoomsRequest proto.InternalMessageInfo

func (m *ListClassRoomsRequest) GetTagIds() []int32 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type ListClassRoomsResponse struct {
	ClassRooms           []*ClassRoom `protobuf:"bytes,1,rep,name=class_rooms,json=classRooms,proto3" json:"class_rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListClassRoomsResponse) Reset()         { *m = ListClassRoomsResponse{} }
func (m *ListClassRoomsResponse) String() string { return proto.CompactTextString(m) }
func (*ListClassRoomsResponse) ProtoMessage()    {}
func (*ListClassRoomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{2}
}
func (m *ListClassRoomsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClassRoomsResponse.Unmarshal(m, b)
}
func (m *ListClassRoomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClassRoomsResponse.Marshal(b, m, deterministic)
}
func (dst *ListClassRoomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClassRoomsResponse.Merge(dst, src)
}
func (m *ListClassRoomsResponse) XXX_Size() int {
	return xxx_messageInfo_ListClassRoomsResponse.Size(m)
}
func (m *ListClassRoomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClassRoomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListClassRoomsResponse proto.InternalMessageInfo

func (m *ListClassRoomsResponse) GetClassRooms() []*ClassRoom {
	if m != nil {
		return m.ClassRooms
	}
	return nil
}

type GetClassRoomRequest struct {
	ClassRoomId          int32    `protobuf:"varint,1,opt,name=class_room_id,json=classRoomId,proto3" json:"class_room_id,omitempty"`
	TagIds               []int32  `protobuf:"varint,2,rep,packed,name=tag_ids,json=tagIds,proto3" json:"tag_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassRoomRequest) Reset()         { *m = GetClassRoomRequest{} }
func (m *GetClassRoomRequest) String() string { return proto.CompactTextString(m) }
func (*GetClassRoomRequest) ProtoMessage()    {}
func (*GetClassRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{3}
}
func (m *GetClassRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassRoomRequest.Unmarshal(m, b)
}
func (m *GetClassRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassRoomRequest.Marshal(b, m, deterministic)
}
func (dst *GetClassRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassRoomRequest.Merge(dst, src)
}
func (m *GetClassRoomRequest) XXX_Size() int {
	return xxx_messageInfo_GetClassRoomRequest.Size(m)
}
func (m *GetClassRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassRoomRequest proto.InternalMessageInfo

func (m *GetClassRoomRequest) GetClassRoomId() int32 {
	if m != nil {
		return m.ClassRoomId
	}
	return 0
}

func (m *GetClassRoomRequest) GetTagIds() []int32 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type CreateClassRoomRequest struct {
	ClassRoom            *ClassRoom `protobuf:"bytes,1,opt,name=class_room,json=classRoom,proto3" json:"class_room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateClassRoomRequest) Reset()         { *m = CreateClassRoomRequest{} }
func (m *CreateClassRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClassRoomRequest) ProtoMessage()    {}
func (*CreateClassRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{4}
}
func (m *CreateClassRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClassRoomRequest.Unmarshal(m, b)
}
func (m *CreateClassRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClassRoomRequest.Marshal(b, m, deterministic)
}
func (dst *CreateClassRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClassRoomRequest.Merge(dst, src)
}
func (m *CreateClassRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClassRoomRequest.Size(m)
}
func (m *CreateClassRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClassRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClassRoomRequest proto.InternalMessageInfo

func (m *CreateClassRoomRequest) GetClassRoom() *ClassRoom {
	if m != nil {
		return m.ClassRoom
	}
	return nil
}

type UpdateClassRoomRequest struct {
	ClassRoom            *ClassRoom `protobuf:"bytes,1,opt,name=class_room,json=classRoom,proto3" json:"class_room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateClassRoomRequest) Reset()         { *m = UpdateClassRoomRequest{} }
func (m *UpdateClassRoomRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateClassRoomRequest) ProtoMessage()    {}
func (*UpdateClassRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{5}
}
func (m *UpdateClassRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClassRoomRequest.Unmarshal(m, b)
}
func (m *UpdateClassRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClassRoomRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateClassRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClassRoomRequest.Merge(dst, src)
}
func (m *UpdateClassRoomRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateClassRoomRequest.Size(m)
}
func (m *UpdateClassRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClassRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClassRoomRequest proto.InternalMessageInfo

func (m *UpdateClassRoomRequest) GetClassRoom() *ClassRoom {
	if m != nil {
		return m.ClassRoom
	}
	return nil
}

type DeleteClassRoomRequest struct {
	ClassRoomId          int32    `protobuf:"varint,1,opt,name=class_room_id,json=classRoomId,proto3" json:"class_room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteClassRoomRequest) Reset()         { *m = DeleteClassRoomRequest{} }
func (m *DeleteClassRoomRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClassRoomRequest) ProtoMessage()    {}
func (*DeleteClassRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_class_rooms_23aedc398cbb4aff, []int{6}
}
func (m *DeleteClassRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClassRoomRequest.Unmarshal(m, b)
}
func (m *DeleteClassRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClassRoomRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteClassRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClassRoomRequest.Merge(dst, src)
}
func (m *DeleteClassRoomRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteClassRoomRequest.Size(m)
}
func (m *DeleteClassRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClassRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClassRoomRequest proto.InternalMessageInfo

func (m *DeleteClassRoomRequest) GetClassRoomId() int32 {
	if m != nil {
		return m.ClassRoomId
	}
	return 0
}

func init() {
	proto.RegisterType((*ClassRoom)(nil), "gedorinku.tsugidoko_server.ClassRoom")
	proto.RegisterType((*ListClassRoomsRequest)(nil), "gedorinku.tsugidoko_server.ListClassRoomsRequest")
	proto.RegisterType((*ListClassRoomsResponse)(nil), "gedorinku.tsugidoko_server.ListClassRoomsResponse")
	proto.RegisterType((*GetClassRoomRequest)(nil), "gedorinku.tsugidoko_server.GetClassRoomRequest")
	proto.RegisterType((*CreateClassRoomRequest)(nil), "gedorinku.tsugidoko_server.CreateClassRoomRequest")
	proto.RegisterType((*UpdateClassRoomRequest)(nil), "gedorinku.tsugidoko_server.UpdateClassRoomRequest")
	proto.RegisterType((*DeleteClassRoomRequest)(nil), "gedorinku.tsugidoko_server.DeleteClassRoomRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClassRoomServiceClient is the client API for ClassRoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClassRoomServiceClient interface {
	ListClassRooms(ctx context.Context, in *ListClassRoomsRequest, opts ...grpc.CallOption) (*ListClassRoomsResponse, error)
	GetClassRoom(ctx context.Context, in *GetClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error)
	CreateClassRoom(ctx context.Context, in *CreateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error)
	UpdateClassRoom(ctx context.Context, in *UpdateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error)
	DeleteClassRoom(ctx context.Context, in *DeleteClassRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type classRoomServiceClient struct {
	cc *grpc.ClientConn
}

func NewClassRoomServiceClient(cc *grpc.ClientConn) ClassRoomServiceClient {
	return &classRoomServiceClient{cc}
}

func (c *classRoomServiceClient) ListClassRooms(ctx context.Context, in *ListClassRoomsRequest, opts ...grpc.CallOption) (*ListClassRoomsResponse, error) {
	out := new(ListClassRoomsResponse)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.ClassRoomService/ListClassRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classRoomServiceClient) GetClassRoom(ctx context.Context, in *GetClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error) {
	out := new(ClassRoom)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.ClassRoomService/GetClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classRoomServiceClient) CreateClassRoom(ctx context.Context, in *CreateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error) {
	out := new(ClassRoom)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.ClassRoomService/CreateClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classRoomServiceClient) UpdateClassRoom(ctx context.Context, in *UpdateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error) {
	out := new(ClassRoom)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.ClassRoomService/UpdateClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classRoomServiceClient) DeleteClassRoom(ctx context.Context, in *DeleteClassRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gedorinku.tsugidoko_server.ClassRoomService/DeleteClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassRoomServiceServer is the server API for ClassRoomService service.
type ClassRoomServiceServer interface {
	ListClassRooms(context.Context, *ListClassRoomsRequest) (*ListClassRoomsResponse, error)
	GetClassRoom(context.Context, *GetClassRoomRequest) (*ClassRoom, error)
	CreateClassRoom(context.Context, *CreateClassRoomRequest) (*ClassRoom, error)
	UpdateClassRoom(context.Context, *UpdateClassRoomRequest) (*ClassRoom, error)
	DeleteClassRoom(context.Context, *DeleteClassRoomRequest) (*empty.Empty, error)
}

func RegisterClassRoomServiceServer(s *grpc.Server, srv ClassRoomServiceServer) {
	s.RegisterService(&_ClassRoomService_serviceDesc, srv)
}

func _ClassRoomService_ListClassRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClassRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassRoomServiceServer).ListClassRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.ClassRoomService/ListClassRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassRoomServiceServer).ListClassRooms(ctx, req.(*ListClassRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassRoomService_GetClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassRoomServiceServer).GetClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.ClassRoomService/GetClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassRoomServiceServer).GetClassRoom(ctx, req.(*GetClassRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassRoomService_CreateClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassRoomServiceServer).CreateClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.ClassRoomService/CreateClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassRoomServiceServer).CreateClassRoom(ctx, req.(*CreateClassRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassRoomService_UpdateClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassRoomServiceServer).UpdateClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.ClassRoomService/UpdateClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassRoomServiceServer).UpdateClassRoom(ctx, req.(*UpdateClassRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassRoomService_DeleteClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClassRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassRoomServiceServer).DeleteClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gedorinku.tsugidoko_server.ClassRoomService/DeleteClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassRoomServiceServer).DeleteClassRoom(ctx, req.(*DeleteClassRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClassRoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gedorinku.tsugidoko_server.ClassRoomService",
	HandlerType: (*ClassRoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClassRooms",
			Handler:    _ClassRoomService_ListClassRooms_Handler,
		},
		{
			MethodName: "GetClassRoom",
			Handler:    _ClassRoomService_GetClassRoom_Handler,
		},
		{
			MethodName: "CreateClassRoom",
			Handler:    _ClassRoomService_CreateClassRoom_Handler,
		},
		{
			MethodName: "UpdateClassRoom",
			Handler:    _ClassRoomService_UpdateClassRoom_Handler,
		},
		{
			MethodName: "DeleteClassRoom",
			Handler:    _ClassRoomService_DeleteClassRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "class_rooms.proto",
}

func init() { proto.RegisterFile("class_rooms.proto", fileDescriptor_class_rooms_23aedc398cbb4aff) }

var fileDescriptor_class_rooms_23aedc398cbb4aff = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xe5, 0xee, 0xdf, 0x6d, 0x7d, 0xbb, 0xff, 0xc6, 0xbc, 0x51, 0xa2, 0x30, 0xa1, 0x28,
	0x1a, 0x5a, 0x41, 0x22, 0x19, 0x41, 0x5c, 0x18, 0x17, 0xd6, 0x0d, 0x34, 0x89, 0x53, 0x00, 0x09,
	0x38, 0x50, 0xdc, 0xc4, 0x0b, 0xd6, 0xd2, 0x38, 0xc4, 0xce, 0xc4, 0x84, 0xb8, 0xc0, 0x15, 0xb8,
	0xf0, 0x29, 0xf8, 0x3c, 0xdc, 0x38, 0xf3, 0x11, 0xf8, 0x00, 0x28, 0x4e, 0x93, 0xa6, 0xa5, 0x0b,
	0xe5, 0xc0, 0x2d, 0xb1, 0xfd, 0x3e, 0xcf, 0xcf, 0xef, 0xfb, 0x48, 0x86, 0x75, 0x2f, 0x24, 0x42,
	0xf4, 0x13, 0xce, 0x87, 0xc2, 0x8a, 0x13, 0x2e, 0x39, 0xd6, 0x03, 0xea, 0xf3, 0x84, 0x45, 0x27,
	0xa9, 0x25, 0x45, 0x1a, 0x30, 0x9f, 0x9f, 0xf0, 0xbe, 0xa0, 0xc9, 0x29, 0x4d, 0xf4, 0xad, 0x80,
	0xf3, 0x20, 0xa4, 0x36, 0x89, 0x99, 0x4d, 0xa2, 0x88, 0x4b, 0x22, 0x19, 0x8f, 0x46, 0x95, 0xfa,
	0xe5, 0xd1, 0xae, 0xfa, 0x1b, 0xa4, 0xc7, 0x36, 0x1d, 0xc6, 0xf2, 0x6c, 0xb4, 0x09, 0x92, 0x04,
	0xc5, 0xc1, 0x75, 0x79, 0x16, 0x53, 0x7b, 0x40, 0x89, 0xc7, 0xa3, 0xd1, 0xd2, 0x46, 0xbe, 0x94,
	0xb2, 0xd0, 0x67, 0x51, 0x90, 0x2f, 0x9a, 0xdf, 0x1b, 0xd0, 0xea, 0x65, 0x80, 0x2e, 0xe7, 0x43,
	0x6c, 0xc2, 0xff, 0x63, 0xda, 0x3e, 0xf3, 0x35, 0x64, 0xa0, 0x6e, 0xd3, 0x6d, 0x7b, 0xc5, 0x89,
	0x23, 0x1f, 0x63, 0xf8, 0x2f, 0x22, 0x43, 0xaa, 0x35, 0x0c, 0xd4, 0x6d, 0xb9, 0xea, 0x1b, 0xf7,
	0x20, 0xf3, 0xee, 0x7b, 0x3c, 0x8d, 0xa4, 0xd0, 0x9a, 0xc6, 0x42, 0xb7, 0xed, 0x6c, 0x5b, 0xe7,
	0xdf, 0xd2, 0x7a, 0x4c, 0x82, 0x5e, 0x76, 0xd8, 0x6d, 0xc9, 0xd1, 0x97, 0xc0, 0xf7, 0x60, 0x29,
	0xe7, 0x15, 0xda, 0x92, 0x52, 0xd8, 0xa9, 0x53, 0xc8, 0x2e, 0x63, 0xed, 0xab, 0xf3, 0x6e, 0x51,
	0x87, 0x37, 0xa1, 0x79, 0x1c, 0x72, 0x9e, 0x68, 0xcb, 0x8a, 0x3b, 0xff, 0xc1, 0x97, 0x60, 0x29,
	0xe4, 0x1e, 0x09, 0xfb, 0x6f, 0xb4, 0x96, 0x81, 0xba, 0xc8, 0x5d, 0x54, 0xbf, 0x4f, 0xc7, 0x1b,
	0x67, 0x1a, 0x54, 0x36, 0x9e, 0xe1, 0x43, 0x58, 0x2e, 0xfa, 0xa4, 0xb5, 0x0d, 0xd4, 0x6d, 0x3b,
	0xd7, 0xfe, 0xcc, 0x32, 0x2a, 0x70, 0xcb, 0x52, 0x73, 0x17, 0x2e, 0x3e, 0x64, 0x42, 0x96, 0xfd,
	0x15, 0x2e, 0x7d, 0x9d, 0x52, 0x21, 0x33, 0xe3, 0xac, 0x5f, 0xcc, 0x17, 0x1a, 0x32, 0x16, 0xba,
	0x4d, 0x77, 0x51, 0x92, 0xe0, 0xc8, 0x17, 0xe6, 0x4b, 0xe8, 0x4c, 0x57, 0x88, 0x98, 0x47, 0x82,
	0xe2, 0xfb, 0xd0, 0xae, 0x04, 0x49, 0x95, 0xb5, 0x9d, 0xab, 0x75, 0x54, 0xa5, 0x88, 0x0b, 0xe5,
	0xfc, 0x84, 0xe9, 0xc2, 0xc6, 0x03, 0x3a, 0x36, 0x28, 0x88, 0xe6, 0x99, 0x7c, 0x85, 0xba, 0x31,
	0x41, 0xfd, 0x02, 0x3a, 0xbd, 0x84, 0x12, 0x49, 0x7f, 0x93, 0x3d, 0x00, 0x18, 0xcb, 0x2a, 0xcd,
	0xb9, 0xa1, 0x5b, 0xa5, 0x75, 0xa6, 0xff, 0x24, 0xf6, 0xff, 0x9d, 0xfe, 0x5d, 0xe8, 0x1c, 0xd0,
	0x90, 0xce, 0xd0, 0x9f, 0xa3, 0x2d, 0xce, 0xcf, 0x26, 0x5c, 0x28, 0x0b, 0x1f, 0xd1, 0xe4, 0x94,
	0x79, 0x14, 0x7f, 0x42, 0xb0, 0x3a, 0x39, 0x49, 0x7c, 0xb3, 0x8e, 0x6b, 0x66, 0x4e, 0x74, 0xe7,
	0x6f, 0x4a, 0xf2, 0xa0, 0x98, 0x9b, 0xef, 0xbf, 0xfd, 0xf8, 0xd2, 0x58, 0xc5, 0x2b, 0x76, 0x25,
	0x2f, 0xf8, 0x23, 0x82, 0x95, 0xea, 0xdc, 0xb1, 0x5d, 0x27, 0x3d, 0x23, 0x21, 0xfa, 0x7c, 0x6d,
	0x35, 0xb7, 0x95, 0xfd, 0x15, 0xbc, 0x55, 0xb5, 0xb7, 0xdf, 0x4e, 0x74, 0xf1, 0x1d, 0xfe, 0x8c,
	0x60, 0x6d, 0x2a, 0x32, 0xb8, 0xf6, 0xb2, 0xb3, 0xf3, 0x35, 0x2f, 0x94, 0xa1, 0xa0, 0x74, 0x73,
	0xa2, 0x27, 0x77, 0x2a, 0xd1, 0xc1, 0x5f, 0x11, 0xac, 0x4d, 0x65, 0xac, 0x1e, 0x68, 0x76, 0x20,
	0xe7, 0x05, 0xda, 0x53, 0x40, 0xb7, 0x9d, 0x9d, 0xf3, 0xba, 0x64, 0x4d, 0x36, 0x6c, 0x82, 0xf5,
	0x03, 0x82, 0xb5, 0xa9, 0xbc, 0xd6, 0xb3, 0xce, 0x0e, 0xb7, 0xde, 0xb1, 0xf2, 0xd7, 0xc4, 0x2a,
	0x5e, 0x13, 0xeb, 0x30, 0x7b, 0x4d, 0x8a, 0x11, 0x5e, 0xaf, 0x1d, 0xe1, 0xbe, 0xf3, 0x7c, 0x37,
	0x60, 0xf2, 0x55, 0x3a, 0xb0, 0x3c, 0x3e, 0xb4, 0x4b, 0x77, 0xbb, 0x74, 0xbf, 0x91, 0xbb, 0x67,
	0x2f, 0xd9, 0x1e, 0x89, 0x59, 0x3f, 0x1e, 0x0c, 0x16, 0x95, 0xd3, 0xad, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x68, 0x92, 0xf0, 0x48, 0x14, 0x07, 0x00, 0x00,
}
