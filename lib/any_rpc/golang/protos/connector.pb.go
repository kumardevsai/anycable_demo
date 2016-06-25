// Code generated by protoc-gen-go.
// source: connector.proto
// DO NOT EDIT!

/*
Package anycable is a generated protocol buffer package.

It is generated from these files:
	connector.proto

It has these top-level messages:
	ConnectionRequest
	ConnectionResponse
*/
package anycable

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ConnectionStatus int32

const (
	ConnectionStatus_ERROR   ConnectionStatus = 0
	ConnectionStatus_SUCCESS ConnectionStatus = 1
)

var ConnectionStatus_name = map[int32]string{
	0: "ERROR",
	1: "SUCCESS",
}
var ConnectionStatus_value = map[string]int32{
	"ERROR":   0,
	"SUCCESS": 1,
}

func (x ConnectionStatus) String() string {
	return proto.EnumName(ConnectionStatus_name, int32(x))
}
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConnectionRequest struct {
	Path    string            `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConnectionRequest) Reset()                    { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()               {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectionRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type ConnectionResponse struct {
	Status ConnectionStatus `protobuf:"varint,1,opt,name=status,enum=anycable.ConnectionStatus" json:"status,omitempty"`
}

func (m *ConnectionResponse) Reset()                    { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()               {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ConnectionRequest)(nil), "anycable.ConnectionRequest")
	proto.RegisterType((*ConnectionResponse)(nil), "anycable.ConnectionResponse")
	proto.RegisterEnum("anycable.ConnectionStatus", ConnectionStatus_name, ConnectionStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Connector service

type ConnectorClient interface {
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := grpc.Invoke(ctx, "/anycable.Connector/Connect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connector service

type ConnectorServer interface {
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.Connector/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "anycable.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Connector_Connect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("connector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0x96, 0x36, 0xe4, 0x8a, 0x20, 0x9c, 0x18, 0xaa, 0xc0, 0x50, 0x65, 0x8a, 0x3a,
	0x64, 0x08, 0x0b, 0xea, 0x48, 0x14, 0xd4, 0x0d, 0xc9, 0x16, 0x0f, 0xe0, 0x86, 0x93, 0x8a, 0xa8,
	0xec, 0x10, 0x3b, 0x48, 0x79, 0x22, 0x5e, 0x13, 0x6a, 0x27, 0x52, 0x05, 0x15, 0xdb, 0x9d, 0xbf,
	0xdf, 0xff, 0x7f, 0xf6, 0xc1, 0x55, 0xa5, 0x95, 0xa2, 0xca, 0xea, 0x26, 0xab, 0x1b, 0x6d, 0x35,
	0x9e, 0x4b, 0xd5, 0x55, 0x72, 0xbb, 0xa7, 0xe4, 0x8b, 0xc1, 0x75, 0xe1, 0xe9, 0x9b, 0x56, 0x9c,
	0x3e, 0x5a, 0x32, 0x16, 0x11, 0xce, 0x6a, 0x69, 0x77, 0x0b, 0xb6, 0x64, 0x69, 0xc8, 0x5d, 0x8d,
	0x8f, 0x10, 0xec, 0x48, 0xbe, 0x52, 0x63, 0x16, 0xe3, 0xe5, 0x24, 0x9d, 0xe7, 0x69, 0x36, 0xb8,
	0x64, 0x7f, 0x1c, 0xb2, 0x8d, 0x97, 0x96, 0xca, 0x36, 0x1d, 0x1f, 0x2e, 0xc6, 0x6b, 0xb8, 0x38,
	0x06, 0x18, 0xc1, 0xe4, 0x9d, 0xba, 0x3e, 0xe6, 0x50, 0xe2, 0x0d, 0x4c, 0x3f, 0xe5, 0xbe, 0xa5,
	0x9f, 0x8c, 0xc3, 0x99, 0x6f, 0xd6, 0xe3, 0x07, 0x96, 0x6c, 0x00, 0x8f, 0x63, 0x4c, 0xad, 0x95,
	0x21, 0xcc, 0x61, 0x66, 0xac, 0xb4, 0xad, 0x71, 0x26, 0x97, 0x79, 0x7c, 0x6a, 0x28, 0xe1, 0x14,
	0xbc, 0x57, 0xae, 0x56, 0x10, 0xfd, 0x66, 0x18, 0xc2, 0xb4, 0xe4, 0xfc, 0x99, 0x47, 0x23, 0x9c,
	0x43, 0x20, 0x5e, 0x8a, 0xa2, 0x14, 0x22, 0x62, 0xb9, 0x80, 0xb0, 0x18, 0x3e, 0x0f, 0x9f, 0x20,
	0xe8, 0x1b, 0xbc, 0xfd, 0xe7, 0xf1, 0xf1, 0xdd, 0x69, 0xe8, 0x47, 0x4e, 0x46, 0xdb, 0x99, 0xdb,
	0xc2, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xfa, 0x2e, 0x1f, 0x98, 0x01, 0x00, 0x00,
}