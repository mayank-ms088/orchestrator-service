// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: dummy_data_service.proto

package datamock

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DummyDataServiceClient is the client API for DummyDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyDataServiceClient interface {
	GetMockUserData(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*User, error)
}

type dummyDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyDataServiceClient(cc grpc.ClientConnInterface) DummyDataServiceClient {
	return &dummyDataServiceClient{cc}
}

func (c *dummyDataServiceClient) GetMockUserData(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/datamock.DummyDataService/GetMockUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyDataServiceServer is the server API for DummyDataService service.
// All implementations must embed UnimplementedDummyDataServiceServer
// for forward compatibility
type DummyDataServiceServer interface {
	GetMockUserData(context.Context, *RPCRequest) (*User, error)
	mustEmbedUnimplementedDummyDataServiceServer()
}

// UnimplementedDummyDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDummyDataServiceServer struct {
}

func (UnimplementedDummyDataServiceServer) GetMockUserData(context.Context, *RPCRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMockUserData not implemented")
}
func (UnimplementedDummyDataServiceServer) mustEmbedUnimplementedDummyDataServiceServer() {}

// UnsafeDummyDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyDataServiceServer will
// result in compilation errors.
type UnsafeDummyDataServiceServer interface {
	mustEmbedUnimplementedDummyDataServiceServer()
}

func RegisterDummyDataServiceServer(s grpc.ServiceRegistrar, srv DummyDataServiceServer) {
	s.RegisterService(&DummyDataService_ServiceDesc, srv)
}

func _DummyDataService_GetMockUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyDataServiceServer).GetMockUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamock.DummyDataService/GetMockUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyDataServiceServer).GetMockUserData(ctx, req.(*RPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DummyDataService_ServiceDesc is the grpc.ServiceDesc for DummyDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DummyDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datamock.DummyDataService",
	HandlerType: (*DummyDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMockUserData",
			Handler:    _DummyDataService_GetMockUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dummy_data_service.proto",
}
