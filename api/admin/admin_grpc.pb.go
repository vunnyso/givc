// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: admin.proto

package admin

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

const (
	AdminService_RegisterService_FullMethodName   = "/admin.AdminService/RegisterService"
	AdminService_StartApplication_FullMethodName  = "/admin.AdminService/StartApplication"
	AdminService_PauseApplication_FullMethodName  = "/admin.AdminService/PauseApplication"
	AdminService_ResumeApplication_FullMethodName = "/admin.AdminService/ResumeApplication"
	AdminService_StopApplication_FullMethodName   = "/admin.AdminService/StopApplication"
	AdminService_Poweroff_FullMethodName          = "/admin.AdminService/Poweroff"
	AdminService_Reboot_FullMethodName            = "/admin.AdminService/Reboot"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	RegisterService(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error)
	StartApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	PauseApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	ResumeApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	StopApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	Poweroff(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Reboot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) RegisterService(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error) {
	out := new(RegistryResponse)
	err := c.cc.Invoke(ctx, AdminService_RegisterService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) StartApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, AdminService_StartApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PauseApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, AdminService_PauseApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ResumeApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, AdminService_ResumeApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) StopApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, AdminService_StopApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Poweroff(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_Poweroff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Reboot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_Reboot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	RegisterService(context.Context, *RegistryRequest) (*RegistryResponse, error)
	StartApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	PauseApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	ResumeApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	StopApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	Poweroff(context.Context, *Empty) (*Empty, error)
	Reboot(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) RegisterService(context.Context, *RegistryRequest) (*RegistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedAdminServiceServer) StartApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartApplication not implemented")
}
func (UnimplementedAdminServiceServer) PauseApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseApplication not implemented")
}
func (UnimplementedAdminServiceServer) ResumeApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeApplication not implemented")
}
func (UnimplementedAdminServiceServer) StopApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopApplication not implemented")
}
func (UnimplementedAdminServiceServer) Poweroff(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poweroff not implemented")
}
func (UnimplementedAdminServiceServer) Reboot(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reboot not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RegisterService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RegisterService(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_StartApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).StartApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_StartApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).StartApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PauseApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PauseApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PauseApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PauseApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ResumeApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ResumeApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ResumeApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ResumeApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_StopApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).StopApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_StopApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).StopApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Poweroff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Poweroff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_Poweroff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Poweroff(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_Reboot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Reboot(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _AdminService_RegisterService_Handler,
		},
		{
			MethodName: "StartApplication",
			Handler:    _AdminService_StartApplication_Handler,
		},
		{
			MethodName: "PauseApplication",
			Handler:    _AdminService_PauseApplication_Handler,
		},
		{
			MethodName: "ResumeApplication",
			Handler:    _AdminService_ResumeApplication_Handler,
		},
		{
			MethodName: "StopApplication",
			Handler:    _AdminService_StopApplication_Handler,
		},
		{
			MethodName: "Poweroff",
			Handler:    _AdminService_Poweroff_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _AdminService_Reboot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
