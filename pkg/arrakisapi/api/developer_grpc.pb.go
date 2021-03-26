// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package REDACTEDapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DeveloperServiceClient is the client API for DeveloperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeveloperServiceClient interface {
	EditCheck(ctx context.Context, in *EditCheckRequest, opts ...grpc.CallOption) (*EditCheckResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	LookupShared(ctx context.Context, in *LookupShareRequest, opts ...grpc.CallOption) (*LookupShareResponse, error)
}

type developerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeveloperServiceClient(cc grpc.ClientConnInterface) DeveloperServiceClient {
	return &developerServiceClient{cc}
}

func (c *developerServiceClient) EditCheck(ctx context.Context, in *EditCheckRequest, opts ...grpc.CallOption) (*EditCheckResponse, error) {
	out := new(EditCheckResponse)
	err := c.cc.Invoke(ctx, "/DeveloperService/EditCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/DeveloperService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/DeveloperService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) LookupShared(ctx context.Context, in *LookupShareRequest, opts ...grpc.CallOption) (*LookupShareResponse, error) {
	out := new(LookupShareResponse)
	err := c.cc.Invoke(ctx, "/DeveloperService/LookupShared", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeveloperServiceServer is the server API for DeveloperService service.
// All implementations must embed UnimplementedDeveloperServiceServer
// for forward compatibility
type DeveloperServiceServer interface {
	EditCheck(context.Context, *EditCheckRequest) (*EditCheckResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Share(context.Context, *ShareRequest) (*ShareResponse, error)
	LookupShared(context.Context, *LookupShareRequest) (*LookupShareResponse, error)
	mustEmbedUnimplementedDeveloperServiceServer()
}

// UnimplementedDeveloperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeveloperServiceServer struct {
}

func (UnimplementedDeveloperServiceServer) EditCheck(context.Context, *EditCheckRequest) (*EditCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCheck not implemented")
}
func (UnimplementedDeveloperServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedDeveloperServiceServer) Share(context.Context, *ShareRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedDeveloperServiceServer) LookupShared(context.Context, *LookupShareRequest) (*LookupShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupShared not implemented")
}
func (UnimplementedDeveloperServiceServer) mustEmbedUnimplementedDeveloperServiceServer() {}

// UnsafeDeveloperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeveloperServiceServer will
// result in compilation errors.
type UnsafeDeveloperServiceServer interface {
	mustEmbedUnimplementedDeveloperServiceServer()
}

func RegisterDeveloperServiceServer(s grpc.ServiceRegistrar, srv DeveloperServiceServer) {
	s.RegisterService(&_DeveloperService_serviceDesc, srv)
}

func _DeveloperService_EditCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).EditCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeveloperService/EditCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).EditCheck(ctx, req.(*EditCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeveloperService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeveloperService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_LookupShared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).LookupShared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeveloperService/LookupShared",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).LookupShared(ctx, req.(*LookupShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeveloperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DeveloperService",
	HandlerType: (*DeveloperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditCheck",
			Handler:    _DeveloperService_EditCheck_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _DeveloperService_Validate_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _DeveloperService_Share_Handler,
		},
		{
			MethodName: "LookupShared",
			Handler:    _DeveloperService_LookupShared_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTEDapi/api/developer.proto",
}
