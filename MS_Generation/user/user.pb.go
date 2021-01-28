// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MsgRequest struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRequest.Unmarshal(m, b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return xxx_messageInfo_MsgRequest.Size(m)
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

func (m *MsgRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type MsgReply struct {
	Sent                 bool     `protobuf:"varint,1,opt,name=sent,proto3" json:"sent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgReply) Reset()         { *m = MsgReply{} }
func (m *MsgReply) String() string { return proto.CompactTextString(m) }
func (*MsgReply) ProtoMessage()    {}
func (*MsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *MsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgReply.Unmarshal(m, b)
}
func (m *MsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgReply.Marshal(b, m, deterministic)
}
func (m *MsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReply.Merge(m, src)
}
func (m *MsgReply) XXX_Size() int {
	return xxx_messageInfo_MsgReply.Size(m)
}
func (m *MsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReply proto.InternalMessageInfo

func (m *MsgReply) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

func init() {
	proto.RegisterType((*MsgRequest)(nil), "MsgRequest")
	proto.RegisterType((*MsgReply)(nil), "MsgReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe3, 0xe2, 0xf2, 0x2d, 0x4e, 0x0f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x95, 0xe4, 0xb8, 0x38, 0xc0, 0xea, 0x0a, 0x72,
	0x2a, 0x85, 0x84, 0xb8, 0x58, 0x8a, 0x53, 0xf3, 0x4a, 0xc0, 0x4a, 0x38, 0x82, 0xc0, 0x6c, 0x23,
	0x1d, 0x2e, 0x96, 0xd0, 0xe2, 0xd4, 0x22, 0x21, 0x15, 0x2e, 0x8e, 0xe0, 0xd4, 0xbc, 0x94, 0x80,
	0xc4, 0xe2, 0x62, 0x21, 0x6e, 0x3d, 0x84, 0xd1, 0x52, 0x9c, 0x7a, 0x30, 0xfd, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x20, 0x45, 0x7a, 0x8a, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	SendPass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) SendPass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/User/SendPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	SendPass(context.Context, *MsgRequest) (*MsgReply, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) SendPass(ctx context.Context, req *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPass not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_SendPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/SendPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendPass(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPass",
			Handler:    _User_SendPass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
