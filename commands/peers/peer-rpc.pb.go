// Code generated by protoc-gen-go.
// source: commands/peers/peer-rpc.proto
// DO NOT EDIT!

/*
Package peercommands is a generated protocol buffer package.

It is generated from these files:
	commands/peers/peer-rpc.proto

It has these top-level messages:
	PeerAddReq
	PeerAddResp
	PeerGenericResp
	PeerDeleteReq
	EtcdConfigReq
*/
package peercommands

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

type PeerAddReq struct {
	Name      string   `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=Addresses,json=addresses" json:"Addresses,omitempty"`
	Client    bool     `protobuf:"varint,3,opt,name=Client,json=client" json:"Client,omitempty"`
}

func (m *PeerAddReq) Reset()                    { *m = PeerAddReq{} }
func (m *PeerAddReq) String() string            { return proto.CompactTextString(m) }
func (*PeerAddReq) ProtoMessage()               {}
func (*PeerAddReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PeerAddResp struct {
	OpRet   int32  `protobuf:"varint,1,opt,name=OpRet,json=opRet" json:"OpRet,omitempty"`
	OpError string `protobuf:"bytes,2,opt,name=OpError,json=opError" json:"OpError,omitempty"`
	UUID    string `protobuf:"bytes,3,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
}

func (m *PeerAddResp) Reset()                    { *m = PeerAddResp{} }
func (m *PeerAddResp) String() string            { return proto.CompactTextString(m) }
func (*PeerAddResp) ProtoMessage()               {}
func (*PeerAddResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PeerGenericResp struct {
	OpRet   int32  `protobuf:"varint,1,opt,name=OpRet,json=opRet" json:"OpRet,omitempty"`
	OpError string `protobuf:"bytes,2,opt,name=OpError,json=opError" json:"OpError,omitempty"`
}

func (m *PeerGenericResp) Reset()                    { *m = PeerGenericResp{} }
func (m *PeerGenericResp) String() string            { return proto.CompactTextString(m) }
func (*PeerGenericResp) ProtoMessage()               {}
func (*PeerGenericResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PeerDeleteReq struct {
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *PeerDeleteReq) Reset()                    { *m = PeerDeleteReq{} }
func (m *PeerDeleteReq) String() string            { return proto.CompactTextString(m) }
func (*PeerDeleteReq) ProtoMessage()               {}
func (*PeerDeleteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EtcdConfigReq struct {
	PeerName       string `protobuf:"bytes,1,opt,name=PeerName,json=peerName" json:"PeerName,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	InitialCluster string `protobuf:"bytes,3,opt,name=InitialCluster,json=initialCluster" json:"InitialCluster,omitempty"`
	ClusterState   string `protobuf:"bytes,4,opt,name=ClusterState,json=clusterState" json:"ClusterState,omitempty"`
	Client         bool   `protobuf:"varint,5,opt,name=Client,json=client" json:"Client,omitempty"`
	DeletePeer     bool   `protobuf:"varint,6,opt,name=DeletePeer,json=deletePeer" json:"DeletePeer,omitempty"`
}

func (m *EtcdConfigReq) Reset()                    { *m = EtcdConfigReq{} }
func (m *EtcdConfigReq) String() string            { return proto.CompactTextString(m) }
func (*EtcdConfigReq) ProtoMessage()               {}
func (*EtcdConfigReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*PeerAddReq)(nil), "peercommands.PeerAddReq")
	proto.RegisterType((*PeerAddResp)(nil), "peercommands.PeerAddResp")
	proto.RegisterType((*PeerGenericResp)(nil), "peercommands.PeerGenericResp")
	proto.RegisterType((*PeerDeleteReq)(nil), "peercommands.PeerDeleteReq")
	proto.RegisterType((*EtcdConfigReq)(nil), "peercommands.EtcdConfigReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PeerService service

type PeerServiceClient interface {
	ValidateAdd(ctx context.Context, in *PeerAddReq, opts ...grpc.CallOption) (*PeerAddResp, error)
	ValidateDelete(ctx context.Context, in *PeerDeleteReq, opts ...grpc.CallOption) (*PeerGenericResp, error)
	ExportAndStoreETCDConfig(ctx context.Context, in *EtcdConfigReq, opts ...grpc.CallOption) (*PeerGenericResp, error)
}

type peerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerServiceClient(cc *grpc.ClientConn) PeerServiceClient {
	return &peerServiceClient{cc}
}

func (c *peerServiceClient) ValidateAdd(ctx context.Context, in *PeerAddReq, opts ...grpc.CallOption) (*PeerAddResp, error) {
	out := new(PeerAddResp)
	err := grpc.Invoke(ctx, "/peercommands.PeerService/ValidateAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerServiceClient) ValidateDelete(ctx context.Context, in *PeerDeleteReq, opts ...grpc.CallOption) (*PeerGenericResp, error) {
	out := new(PeerGenericResp)
	err := grpc.Invoke(ctx, "/peercommands.PeerService/ValidateDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerServiceClient) ExportAndStoreETCDConfig(ctx context.Context, in *EtcdConfigReq, opts ...grpc.CallOption) (*PeerGenericResp, error) {
	out := new(PeerGenericResp)
	err := grpc.Invoke(ctx, "/peercommands.PeerService/ExportAndStoreETCDConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeerService service

type PeerServiceServer interface {
	ValidateAdd(context.Context, *PeerAddReq) (*PeerAddResp, error)
	ValidateDelete(context.Context, *PeerDeleteReq) (*PeerGenericResp, error)
	ExportAndStoreETCDConfig(context.Context, *EtcdConfigReq) (*PeerGenericResp, error)
}

func RegisterPeerServiceServer(s *grpc.Server, srv PeerServiceServer) {
	s.RegisterService(&_PeerService_serviceDesc, srv)
}

func _PeerService_ValidateAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).ValidateAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peercommands.PeerService/ValidateAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).ValidateAdd(ctx, req.(*PeerAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerService_ValidateDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).ValidateDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peercommands.PeerService/ValidateDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).ValidateDelete(ctx, req.(*PeerDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerService_ExportAndStoreETCDConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EtcdConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).ExportAndStoreETCDConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peercommands.PeerService/ExportAndStoreETCDConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).ExportAndStoreETCDConfig(ctx, req.(*EtcdConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peercommands.PeerService",
	HandlerType: (*PeerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateAdd",
			Handler:    _PeerService_ValidateAdd_Handler,
		},
		{
			MethodName: "ValidateDelete",
			Handler:    _PeerService_ValidateDelete_Handler,
		},
		{
			MethodName: "ExportAndStoreETCDConfig",
			Handler:    _PeerService_ExportAndStoreETCDConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("commands/peers/peer-rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0x51, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0x6d, 0xa1, 0x05, 0x06, 0xa8, 0xc9, 0xc6, 0x98, 0x15, 0x45, 0x49, 0x1f, 0x8c, 0x2f,
	0x62, 0xa2, 0x9f, 0x80, 0xb4, 0x8d, 0xe1, 0x05, 0xb4, 0x08, 0xf1, 0xb5, 0xb6, 0xa3, 0xd9, 0xa4,
	0x74, 0xd7, 0xed, 0x72, 0xb9, 0xef, 0x77, 0x9f, 0xeb, 0x92, 0xdb, 0x6d, 0xa1, 0x57, 0x8e, 0x5c,
	0x2e, 0xb9, 0x97, 0x66, 0xff, 0x33, 0xd3, 0xdf, 0xce, 0xfc, 0x67, 0x61, 0x9a, 0xf2, 0xfd, 0x3e,
	0x29, 0xb2, 0xf2, 0x8b, 0x40, 0x94, 0xf5, 0xf7, 0xb3, 0x14, 0xe9, 0x5c, 0x48, 0xae, 0x38, 0x19,
	0x19, 0x7d, 0x2a, 0xf1, 0x77, 0x00, 0x3f, 0xb4, 0x5e, 0x64, 0x59, 0x8c, 0xff, 0x09, 0x81, 0xee,
	0x2a, 0xd9, 0x23, 0xb5, 0x66, 0xd6, 0xa7, 0x41, 0xdc, 0x2d, 0xf4, 0x99, 0xbc, 0x83, 0x81, 0xce,
	0x4a, 0x2c, 0x4b, 0x2c, 0xa9, 0x3d, 0xeb, 0xe8, 0xc4, 0x20, 0x39, 0x05, 0xc8, 0x6b, 0x70, 0x83,
	0x9c, 0x61, 0xa1, 0x68, 0x47, 0xff, 0xd3, 0x8f, 0xdd, 0xb4, 0x52, 0xfe, 0x4f, 0x18, 0x36, 0xdc,
	0x52, 0x90, 0x57, 0xe0, 0xac, 0x45, 0x8c, 0xaa, 0x22, 0x3b, 0xb1, 0xc3, 0x8d, 0x20, 0x14, 0x7a,
	0x6b, 0x11, 0x49, 0xc9, 0xa5, 0x06, 0x9b, 0x1b, 0x7b, 0xbc, 0x96, 0xa6, 0x91, 0xed, 0x76, 0x19,
	0x56, 0x50, 0xdd, 0xc8, 0x41, 0x9f, 0xfd, 0x05, 0xbc, 0x34, 0xc8, 0xef, 0x58, 0xa0, 0x64, 0xe9,
	0x73, 0xb0, 0xfe, 0x07, 0x18, 0x1b, 0x44, 0x88, 0x39, 0x2a, 0x34, 0x03, 0x7b, 0x60, 0xeb, 0x5b,
	0xea, 0x71, 0x6d, 0x16, 0xfa, 0x37, 0x16, 0x8c, 0x23, 0x95, 0x66, 0x01, 0x2f, 0xfe, 0xb2, 0x7f,
	0xa6, 0x62, 0x02, 0x7d, 0xf3, 0x4b, 0xcb, 0x96, 0xbe, 0x38, 0xea, 0xc6, 0x2e, 0xbb, 0x65, 0xd7,
	0x47, 0xf0, 0x96, 0x05, 0x53, 0x2c, 0xc9, 0x83, 0xfc, 0x50, 0x2a, 0x94, 0xc7, 0x19, 0x3c, 0x76,
	0x16, 0x25, 0x3e, 0x8c, 0x8e, 0xc7, 0x8d, 0x4a, 0x14, 0xd2, 0x6e, 0x55, 0x35, 0x4a, 0x5b, 0xb1,
	0x96, 0xb9, 0x4e, 0xdb, 0x5c, 0xf2, 0x1e, 0xa0, 0x1e, 0xc1, 0x74, 0x46, 0xdd, 0x2a, 0x07, 0x59,
	0x13, 0xf9, 0x7a, 0x6b, 0xd5, 0xee, 0x6f, 0x50, 0x5e, 0xb1, 0x14, 0x49, 0x08, 0xc3, 0x5d, 0x92,
	0xb3, 0x4c, 0x33, 0xf5, 0x42, 0x08, 0x9d, 0xb7, 0x9f, 0xc0, 0xfc, 0x7e, 0xff, 0x93, 0x37, 0x8f,
	0x64, 0x4a, 0xe1, 0xbf, 0x20, 0x2b, 0xf0, 0x4e, 0x94, 0xfa, 0x76, 0xf2, 0xf6, 0xb2, 0xbc, 0xb1,
	0x76, 0x32, 0xbd, 0x4c, 0xb6, 0x56, 0xa7, 0x79, 0xbf, 0x81, 0x46, 0xd7, 0x82, 0x4b, 0xb5, 0x28,
	0xb2, 0x8d, 0xe2, 0x12, 0xa3, 0x5f, 0x41, 0x58, 0x1b, 0xff, 0x90, 0x7c, 0xb6, 0x92, 0x27, 0xc9,
	0x7f, 0xdc, 0xea, 0xa5, 0x7f, 0xbb, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x6f, 0x4f, 0x8c, 0x0a,
	0x03, 0x00, 0x00,
}
