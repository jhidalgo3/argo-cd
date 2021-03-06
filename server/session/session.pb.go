// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/session/session.proto

/*
Package session is a generated protocol buffer package.

Session Service

Session Service API performs CRUD actions against session resources

It is generated from these files:
	server/session/session.proto

It has these top-level messages:
	SessionRequest
	SessionResponse
*/
package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "k8s.io/api/core/v1"
import _ "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

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

// SessionQuery is a query for session resources
type SessionRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SessionResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionRequest)(nil), "session.SessionRequest")
	proto.RegisterType((*SessionResponse)(nil), "session.SessionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SessionService service

type SessionServiceClient interface {
	// Create a new JWT for authentication.
	Create(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Create(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := grpc.Invoke(ctx, "/session.SessionService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionService service

type SessionServiceServer interface {
	// Create a new JWT for authentication.
	Create(context.Context, *SessionRequest) (*SessionResponse, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Create(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SessionService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/session/session.proto",
}

func init() { proto.RegisterFile("server/session/session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0x69, 0xe1, 0xed, 0xab, 0x39, 0x28, 0x2c, 0x85, 0x96, 0xa5, 0x07, 0xc9, 0x45, 0x29,
	0xba, 0xa1, 0x7a, 0x11, 0x8f, 0x7a, 0xd1, 0x6b, 0x7b, 0xf3, 0x22, 0x69, 0x3a, 0xa4, 0x71, 0xb7,
	0x99, 0x34, 0x93, 0x5d, 0xef, 0x7e, 0x05, 0x3f, 0x9a, 0x5f, 0xc1, 0x0f, 0x22, 0x9b, 0xfd, 0x03,
	0xa5, 0xe0, 0x29, 0xf3, 0xe4, 0x97, 0x3c, 0x33, 0x3c, 0xc3, 0x66, 0x04, 0xbe, 0x02, 0x2f, 0x08,
	0x88, 0x0c, 0xda, 0xee, 0xcc, 0x9c, 0xc7, 0x80, 0xc9, 0xff, 0x56, 0xa6, 0x63, 0x8d, 0x1a, 0xe3,
	0x9d, 0xa8, 0xab, 0x06, 0xa7, 0x33, 0x8d, 0xa8, 0x0b, 0x10, 0xd2, 0x19, 0x21, 0xad, 0xc5, 0x20,
	0x83, 0x41, 0x4b, 0x2d, 0xe5, 0xf9, 0x3d, 0x65, 0x06, 0x23, 0x55, 0xe8, 0x41, 0x54, 0x0b, 0xa1,
	0xc1, 0x82, 0x97, 0x01, 0x36, 0xed, 0x9b, 0x17, 0x6d, 0xc2, 0xb6, 0x5c, 0x67, 0x0a, 0x77, 0x42,
	0xfa, 0xd8, 0xe2, 0x3d, 0x16, 0x37, 0x6a, 0x23, 0x5c, 0xae, 0xeb, 0xcf, 0x24, 0xa4, 0x73, 0x85,
	0x51, 0xd1, 0x5c, 0x54, 0x0b, 0x59, 0xb8, 0xad, 0x3c, 0xb2, 0xe2, 0xcf, 0xec, 0x6c, 0xd5, 0x4c,
	0xbb, 0x84, 0x7d, 0x09, 0x14, 0x92, 0x94, 0x9d, 0x94, 0x04, 0xde, 0xca, 0x1d, 0x4c, 0x07, 0x17,
	0x83, 0xab, 0xd3, 0x65, 0xaf, 0x6b, 0xe6, 0x24, 0xd1, 0x07, 0xfa, 0xcd, 0x74, 0xd8, 0xb0, 0x4e,
	0xf3, 0x4b, 0x76, 0xde, 0x3b, 0x91, 0x43, 0x4b, 0x90, 0x8c, 0xd9, 0xbf, 0x80, 0x39, 0xd8, 0xd6,
	0xa7, 0x11, 0xb7, 0xfb, 0xbe, 0xe5, 0x0a, 0x7c, 0x65, 0x14, 0x24, 0x6f, 0x6c, 0xf4, 0xe4, 0x41,
	0x06, 0x48, 0x26, 0x59, 0x17, 0xe5, 0xe1, 0x54, 0xe9, 0xf4, 0x18, 0x34, 0x4d, 0x38, 0xff, 0xfc,
	0xfe, 0xf9, 0x1a, 0xce, 0xf8, 0x24, 0x46, 0x56, 0x2d, 0xba, 0x65, 0x90, 0x50, 0xd1, 0xf3, 0x61,
	0x30, 0x7f, 0xbc, 0x7e, 0x9d, 0xff, 0x15, 0xd9, 0xe1, 0x36, 0xd7, 0xa3, 0x18, 0xcd, 0xdd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xd4, 0xc3, 0x1b, 0xe6, 0x01, 0x00, 0x00,
}
