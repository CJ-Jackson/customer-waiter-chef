// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chef.proto

package protocol

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

type WaiterRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaiterRequest) Reset()         { *m = WaiterRequest{} }
func (m *WaiterRequest) String() string { return proto.CompactTextString(m) }
func (*WaiterRequest) ProtoMessage()    {}
func (*WaiterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e705f2f0ff9c45, []int{0}
}

func (m *WaiterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaiterRequest.Unmarshal(m, b)
}
func (m *WaiterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaiterRequest.Marshal(b, m, deterministic)
}
func (m *WaiterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaiterRequest.Merge(m, src)
}
func (m *WaiterRequest) XXX_Size() int {
	return xxx_messageInfo_WaiterRequest.Size(m)
}
func (m *WaiterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaiterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaiterRequest proto.InternalMessageInfo

func (m *WaiterRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type ChefReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChefReply) Reset()         { *m = ChefReply{} }
func (m *ChefReply) String() string { return proto.CompactTextString(m) }
func (*ChefReply) ProtoMessage()    {}
func (*ChefReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e705f2f0ff9c45, []int{1}
}

func (m *ChefReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChefReply.Unmarshal(m, b)
}
func (m *ChefReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChefReply.Marshal(b, m, deterministic)
}
func (m *ChefReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChefReply.Merge(m, src)
}
func (m *ChefReply) XXX_Size() int {
	return xxx_messageInfo_ChefReply.Size(m)
}
func (m *ChefReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChefReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChefReply proto.InternalMessageInfo

func (m *ChefReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*WaiterRequest)(nil), "protocol.WaiterRequest")
	proto.RegisterType((*ChefReply)(nil), "protocol.ChefReply")
}

func init() {
	proto.RegisterFile("chef.proto", fileDescriptor_14e705f2f0ff9c45)
}

var fileDescriptor_14e705f2f0ff9c45 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x4d,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x9a, 0x5c,
	0xbc, 0xe1, 0x89, 0x99, 0x25, 0xa9, 0x45, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12,
	0x5c, 0xec, 0x45, 0x10, 0xa6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xca,
	0xc5, 0xe9, 0x9c, 0x91, 0x9a, 0x16, 0x94, 0x5a, 0x90, 0x53, 0x09, 0x52, 0x96, 0x9b, 0x5a, 0x5c,
	0x9c, 0x98, 0x9e, 0x0a, 0x53, 0x06, 0xe5, 0x1a, 0x79, 0x72, 0xb1, 0x80, 0x94, 0x09, 0x39, 0x72,
	0xf1, 0xfb, 0x17, 0xa5, 0xa4, 0x16, 0xb9, 0x15, 0xe5, 0xe7, 0x42, 0xac, 0x10, 0x12, 0xd7, 0x83,
	0xd9, 0xab, 0x87, 0x62, 0xa9, 0x94, 0x30, 0x42, 0x02, 0x6e, 0x85, 0x12, 0x43, 0x12, 0x1b, 0x58,
	0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x85, 0x01, 0x3d, 0xb1, 0xbb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefClient is the client API for Chef service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefClient interface {
	OrderFromWaiter(ctx context.Context, in *WaiterRequest, opts ...grpc.CallOption) (*ChefReply, error)
}

type chefClient struct {
	cc grpc.ClientConnInterface
}

func NewChefClient(cc grpc.ClientConnInterface) ChefClient {
	return &chefClient{cc}
}

func (c *chefClient) OrderFromWaiter(ctx context.Context, in *WaiterRequest, opts ...grpc.CallOption) (*ChefReply, error) {
	out := new(ChefReply)
	err := c.cc.Invoke(ctx, "/protocol.Chef/OrderFromWaiter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefServer is the server API for Chef service.
type ChefServer interface {
	OrderFromWaiter(context.Context, *WaiterRequest) (*ChefReply, error)
}

// UnimplementedChefServer can be embedded to have forward compatible implementations.
type UnimplementedChefServer struct {
}

func (*UnimplementedChefServer) OrderFromWaiter(ctx context.Context, req *WaiterRequest) (*ChefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderFromWaiter not implemented")
}

func RegisterChefServer(s *grpc.Server, srv ChefServer) {
	s.RegisterService(&_Chef_serviceDesc, srv)
}

func _Chef_OrderFromWaiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefServer).OrderFromWaiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Chef/OrderFromWaiter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefServer).OrderFromWaiter(ctx, req.(*WaiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chef_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Chef",
	HandlerType: (*ChefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderFromWaiter",
			Handler:    _Chef_OrderFromWaiter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chef.proto",
}