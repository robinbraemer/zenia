// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	ApplyNamespace(ctx context.Context, in *ApplyNamespaceRequest, opts ...grpc.CallOption) (*ApplyNamespaceResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) ApplyNamespace(ctx context.Context, in *ApplyNamespaceRequest, opts ...grpc.CallOption) (*ApplyNamespaceResponse, error) {
	out := new(ApplyNamespaceResponse)
	err := c.cc.Invoke(ctx, "/zenia.authz.admin.v1.AdminService/ApplyNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations should embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	ApplyNamespace(context.Context, *ApplyNamespaceRequest) (*ApplyNamespaceResponse, error)
}

// UnimplementedAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) ApplyNamespace(context.Context, *ApplyNamespaceRequest) (*ApplyNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNamespace not implemented")
}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_ApplyNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ApplyNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenia.authz.admin.v1.AdminService/ApplyNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ApplyNamespace(ctx, req.(*ApplyNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zenia.authz.admin.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyNamespace",
			Handler:    _AdminService_ApplyNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zenia/authz/admin/v1/admin.proto",
}
