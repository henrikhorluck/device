// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DeviceHelperClient is the client API for DeviceHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceHelperClient interface {
	// Push and apply new VPN configuration.
	Configure(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*ConfigureResponse, error)
	// Delete VPN configuration and shut down connections.
	Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error)
	// Install the newest version of naisdevice.
	Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error)
}

type deviceHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceHelperClient(cc grpc.ClientConnInterface) DeviceHelperClient {
	return &deviceHelperClient{cc}
}

func (c *deviceHelperClient) Configure(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceHelper/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceHelperClient) Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error) {
	out := new(TeardownResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceHelper/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceHelperClient) Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error) {
	out := new(UpgradeResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceHelper/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceHelperServer is the server API for DeviceHelper service.
// All implementations must embed UnimplementedDeviceHelperServer
// for forward compatibility
type DeviceHelperServer interface {
	// Push and apply new VPN configuration.
	Configure(context.Context, *Configuration) (*ConfigureResponse, error)
	// Delete VPN configuration and shut down connections.
	Teardown(context.Context, *TeardownRequest) (*TeardownResponse, error)
	// Install the newest version of naisdevice.
	Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error)
	mustEmbedUnimplementedDeviceHelperServer()
}

// UnimplementedDeviceHelperServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceHelperServer struct {
}

func (UnimplementedDeviceHelperServer) Configure(context.Context, *Configuration) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedDeviceHelperServer) Teardown(context.Context, *TeardownRequest) (*TeardownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}
func (UnimplementedDeviceHelperServer) Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upgrade not implemented")
}
func (UnimplementedDeviceHelperServer) mustEmbedUnimplementedDeviceHelperServer() {}

// UnsafeDeviceHelperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceHelperServer will
// result in compilation errors.
type UnsafeDeviceHelperServer interface {
	mustEmbedUnimplementedDeviceHelperServer()
}

func RegisterDeviceHelperServer(s grpc.ServiceRegistrar, srv DeviceHelperServer) {
	s.RegisterService(&DeviceHelper_ServiceDesc, srv)
}

func _DeviceHelper_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceHelperServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceHelper/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceHelperServer).Configure(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceHelper_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceHelperServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceHelper/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceHelperServer).Teardown(ctx, req.(*TeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceHelper_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceHelperServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceHelper/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceHelperServer).Upgrade(ctx, req.(*UpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceHelper_ServiceDesc is the grpc.ServiceDesc for DeviceHelper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceHelper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "naisdevice.DeviceHelper",
	HandlerType: (*DeviceHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _DeviceHelper_Configure_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _DeviceHelper_Teardown_Handler,
		},
		{
			MethodName: "Upgrade",
			Handler:    _DeviceHelper_Upgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/protobuf-api.proto",
}

// DeviceAgentClient is the client API for DeviceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceAgentClient interface {
	// DeviceAgent will stream all state changes on this endpoint.
	// Use Status() to continuously monitor the current Agent status.
	Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (DeviceAgent_StatusClient, error)
	// Open the JITA form in a web browser.
	ConfigureJITA(ctx context.Context, in *ConfigureJITARequest, opts ...grpc.CallOption) (*ConfigureJITAResponse, error)
	// Log in to API server, enabling access to protected resources.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Log out of API server, shutting down all VPN connections.
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type deviceAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceAgentClient(cc grpc.ClientConnInterface) DeviceAgentClient {
	return &deviceAgentClient{cc}
}

func (c *deviceAgentClient) Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (DeviceAgent_StatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &DeviceAgent_ServiceDesc.Streams[0], "/naisdevice.DeviceAgent/Status", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceAgentStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeviceAgent_StatusClient interface {
	Recv() (*AgentStatus, error)
	grpc.ClientStream
}

type deviceAgentStatusClient struct {
	grpc.ClientStream
}

func (x *deviceAgentStatusClient) Recv() (*AgentStatus, error) {
	m := new(AgentStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deviceAgentClient) ConfigureJITA(ctx context.Context, in *ConfigureJITARequest, opts ...grpc.CallOption) (*ConfigureJITAResponse, error) {
	out := new(ConfigureJITAResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceAgent/ConfigureJITA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAgentClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceAgent/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAgentClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/naisdevice.DeviceAgent/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceAgentServer is the server API for DeviceAgent service.
// All implementations must embed UnimplementedDeviceAgentServer
// for forward compatibility
type DeviceAgentServer interface {
	// DeviceAgent will stream all state changes on this endpoint.
	// Use Status() to continuously monitor the current Agent status.
	Status(*AgentStatusRequest, DeviceAgent_StatusServer) error
	// Open the JITA form in a web browser.
	ConfigureJITA(context.Context, *ConfigureJITARequest) (*ConfigureJITAResponse, error)
	// Log in to API server, enabling access to protected resources.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Log out of API server, shutting down all VPN connections.
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedDeviceAgentServer()
}

// UnimplementedDeviceAgentServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceAgentServer struct {
}

func (UnimplementedDeviceAgentServer) Status(*AgentStatusRequest, DeviceAgent_StatusServer) error {
	return status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedDeviceAgentServer) ConfigureJITA(context.Context, *ConfigureJITARequest) (*ConfigureJITAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureJITA not implemented")
}
func (UnimplementedDeviceAgentServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDeviceAgentServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedDeviceAgentServer) mustEmbedUnimplementedDeviceAgentServer() {}

// UnsafeDeviceAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceAgentServer will
// result in compilation errors.
type UnsafeDeviceAgentServer interface {
	mustEmbedUnimplementedDeviceAgentServer()
}

func RegisterDeviceAgentServer(s grpc.ServiceRegistrar, srv DeviceAgentServer) {
	s.RegisterService(&DeviceAgent_ServiceDesc, srv)
}

func _DeviceAgent_Status_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeviceAgentServer).Status(m, &deviceAgentStatusServer{stream})
}

type DeviceAgent_StatusServer interface {
	Send(*AgentStatus) error
	grpc.ServerStream
}

type deviceAgentStatusServer struct {
	grpc.ServerStream
}

func (x *deviceAgentStatusServer) Send(m *AgentStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _DeviceAgent_ConfigureJITA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureJITARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAgentServer).ConfigureJITA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceAgent/ConfigureJITA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAgentServer).ConfigureJITA(ctx, req.(*ConfigureJITARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAgent_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAgentServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceAgent/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAgentServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAgent_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAgentServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/naisdevice.DeviceAgent/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAgentServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceAgent_ServiceDesc is the grpc.ServiceDesc for DeviceAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "naisdevice.DeviceAgent",
	HandlerType: (*DeviceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureJITA",
			Handler:    _DeviceAgent_ConfigureJITA_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DeviceAgent_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _DeviceAgent_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Status",
			Handler:       _DeviceAgent_Status_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/pb/protobuf-api.proto",
}
