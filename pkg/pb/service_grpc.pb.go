// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortalClient is the client API for Portal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalClient interface {
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type portalClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalClient(cc grpc.ClientConnInterface) PortalClient {
	return &portalClient{cc}
}

func (c *portalClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalServer is the server API for Portal service.
// All implementations should embed UnimplementedPortalServer
// for forward compatibility
type PortalServer interface {
	GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error)
}

// UnimplementedPortalServer should be embedded to have forward compatible implementations.
type UnimplementedPortalServer struct {
}

func (UnimplementedPortalServer) GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

// UnsafePortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalServer will
// result in compilation errors.
type UnsafePortalServer interface {
	mustEmbedUnimplementedPortalServer()
}

func RegisterPortalServer(s grpc.ServiceRegistrar, srv PortalServer) {
	s.RegisterService(&Portal_ServiceDesc, srv)
}

func _Portal_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Portal_ServiceDesc is the grpc.ServiceDesc for Portal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portal.Portal",
	HandlerType: (*PortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Portal_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro/portal/pkg/pb/service.proto",
}
