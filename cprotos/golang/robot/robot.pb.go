// Code generated by protoc-gen-go. DO NOT EDIT.
// source: robot/robot.proto

package robot

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RobotServiceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotServiceRequest) Reset()         { *m = RobotServiceRequest{} }
func (m *RobotServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RobotServiceRequest) ProtoMessage()    {}
func (*RobotServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d65d06a1694be51, []int{0}
}

func (m *RobotServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotServiceRequest.Unmarshal(m, b)
}
func (m *RobotServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotServiceRequest.Marshal(b, m, deterministic)
}
func (m *RobotServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotServiceRequest.Merge(m, src)
}
func (m *RobotServiceRequest) XXX_Size() int {
	return xxx_messageInfo_RobotServiceRequest.Size(m)
}
func (m *RobotServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RobotServiceRequest proto.InternalMessageInfo

type RobotServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotServiceResponse) Reset()         { *m = RobotServiceResponse{} }
func (m *RobotServiceResponse) String() string { return proto.CompactTextString(m) }
func (*RobotServiceResponse) ProtoMessage()    {}
func (*RobotServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d65d06a1694be51, []int{1}
}

func (m *RobotServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotServiceResponse.Unmarshal(m, b)
}
func (m *RobotServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotServiceResponse.Marshal(b, m, deterministic)
}
func (m *RobotServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotServiceResponse.Merge(m, src)
}
func (m *RobotServiceResponse) XXX_Size() int {
	return xxx_messageInfo_RobotServiceResponse.Size(m)
}
func (m *RobotServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RobotServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RobotServiceRequest)(nil), "robot.RobotServiceRequest")
	proto.RegisterType((*RobotServiceResponse)(nil), "robot.RobotServiceResponse")
}

func init() { proto.RegisterFile("robot/robot.proto", fileDescriptor_9d65d06a1694be51) }

var fileDescriptor_9d65d06a1694be51 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xca, 0x4f, 0xca,
	0x2f, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x28,
	0x97, 0x70, 0x10, 0x88, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x1a, 0x94, 0x5a, 0x58, 0x9a,
	0x5a, 0x5c, 0xa2, 0x24, 0xc6, 0x25, 0x82, 0x2a, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0xe4,
	0xc5, 0xc5, 0x0a, 0x16, 0x17, 0x72, 0xe4, 0x62, 0xf1, 0xcd, 0x2f, 0x4b, 0x15, 0x92, 0xd2, 0x83,
	0x18, 0x8a, 0xc5, 0x10, 0x29, 0x69, 0xac, 0x72, 0x10, 0x93, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x0e,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd4, 0x89, 0xd7, 0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotClient is the client API for Robot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotClient interface {
	Move(ctx context.Context, in *RobotServiceRequest, opts ...grpc.CallOption) (*RobotServiceResponse, error)
}

type robotClient struct {
	cc *grpc.ClientConn
}

func NewRobotClient(cc *grpc.ClientConn) RobotClient {
	return &robotClient{cc}
}

func (c *robotClient) Move(ctx context.Context, in *RobotServiceRequest, opts ...grpc.CallOption) (*RobotServiceResponse, error) {
	out := new(RobotServiceResponse)
	err := c.cc.Invoke(ctx, "/robot.Robot/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServer is the server API for Robot service.
type RobotServer interface {
	Move(context.Context, *RobotServiceRequest) (*RobotServiceResponse, error)
}

func RegisterRobotServer(s *grpc.Server, srv RobotServer) {
	s.RegisterService(&_Robot_serviceDesc, srv)
}

func _Robot_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.Robot/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).Move(ctx, req.(*RobotServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Robot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robot.Robot",
	HandlerType: (*RobotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Move",
			Handler:    _Robot_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robot/robot.proto",
}
