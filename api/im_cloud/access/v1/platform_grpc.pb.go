// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: im_cloud/access/v1/platform.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Platform_DistributeMasterNodeConfig_FullMethodName = "/api.im_cloud.access.v2.Platform/DistributeMasterNodeConfig"
	Platform_DistributeProxyNodeConfig_FullMethodName  = "/api.im_cloud.access.v2.Platform/DistributeProxyNodeConfig"
	Platform_GetNodeKey_FullMethodName                 = "/api.im_cloud.access.v2.Platform/GetNodeKey"
	Platform_GetNodeStatus_FullMethodName              = "/api.im_cloud.access.v2.Platform/GetNodeStatus"
	Platform_GetNodeFunction_FullMethodName            = "/api.im_cloud.access.v2.Platform/GetNodeFunction"
	Platform_DistributeUser_FullMethodName             = "/api.im_cloud.access.v2.Platform/DistributeUser"
	Platform_EscalationUser_FullMethodName             = "/api.im_cloud.access.v2.Platform/EscalationUser"
	Platform_UpdateCenterUserPassword_FullMethodName   = "/api.im_cloud.access.v2.Platform/UpdateCenterUserPassword"
	Platform_CancelAccount_FullMethodName              = "/api.im_cloud.access.v2.Platform/CancelAccount"
)

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlatformClient interface {
	// 平台下发节点配置
	DistributeMasterNodeConfig(ctx context.Context, in *DistributeMasterNodeConfigReq, opts ...grpc.CallOption) (*DistributeMasterNodeReply, error)
	// 平台下发代理节点配置
	DistributeProxyNodeConfig(ctx context.Context, in *DistributeProxyNodeConfigReq, opts ...grpc.CallOption) (*DistributeProxyNodeConfigReply, error)
	GetNodeKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNodeKeyReply, error)
	GetNodeStatus(ctx context.Context, in *GetNodeStatusReq, opts ...grpc.CallOption) (*GetNodeStatusReply, error)
	GetNodeFunction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNodeFunctionReply, error)
	// 平台下发用户信息
	DistributeUser(ctx context.Context, in *DistributeUserReq, opts ...grpc.CallOption) (*DistributeUserReply, error)
	// 上报用户信息
	EscalationUser(ctx context.Context, in *EscalationUserReq, opts ...grpc.CallOption) (*EscalationUserReply, error)
	// 更新用户中心的账号密码
	UpdateCenterUserPassword(ctx context.Context, in *UpdateCenterUserPasswordReq, opts ...grpc.CallOption) (*UpdateCenterUserPasswordReply, error)
	// 注销账号
	CancelAccount(ctx context.Context, in *CancelAccountReq, opts ...grpc.CallOption) (*CancelAccountReply, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) DistributeMasterNodeConfig(ctx context.Context, in *DistributeMasterNodeConfigReq, opts ...grpc.CallOption) (*DistributeMasterNodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DistributeMasterNodeReply)
	err := c.cc.Invoke(ctx, Platform_DistributeMasterNodeConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DistributeProxyNodeConfig(ctx context.Context, in *DistributeProxyNodeConfigReq, opts ...grpc.CallOption) (*DistributeProxyNodeConfigReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DistributeProxyNodeConfigReply)
	err := c.cc.Invoke(ctx, Platform_DistributeProxyNodeConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetNodeKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNodeKeyReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodeKeyReply)
	err := c.cc.Invoke(ctx, Platform_GetNodeKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetNodeStatus(ctx context.Context, in *GetNodeStatusReq, opts ...grpc.CallOption) (*GetNodeStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodeStatusReply)
	err := c.cc.Invoke(ctx, Platform_GetNodeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetNodeFunction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNodeFunctionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodeFunctionReply)
	err := c.cc.Invoke(ctx, Platform_GetNodeFunction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DistributeUser(ctx context.Context, in *DistributeUserReq, opts ...grpc.CallOption) (*DistributeUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DistributeUserReply)
	err := c.cc.Invoke(ctx, Platform_DistributeUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) EscalationUser(ctx context.Context, in *EscalationUserReq, opts ...grpc.CallOption) (*EscalationUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EscalationUserReply)
	err := c.cc.Invoke(ctx, Platform_EscalationUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) UpdateCenterUserPassword(ctx context.Context, in *UpdateCenterUserPasswordReq, opts ...grpc.CallOption) (*UpdateCenterUserPasswordReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCenterUserPasswordReply)
	err := c.cc.Invoke(ctx, Platform_UpdateCenterUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) CancelAccount(ctx context.Context, in *CancelAccountReq, opts ...grpc.CallOption) (*CancelAccountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelAccountReply)
	err := c.cc.Invoke(ctx, Platform_CancelAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
// All implementations must embed UnimplementedPlatformServer
// for forward compatibility
type PlatformServer interface {
	// 平台下发节点配置
	DistributeMasterNodeConfig(context.Context, *DistributeMasterNodeConfigReq) (*DistributeMasterNodeReply, error)
	// 平台下发代理节点配置
	DistributeProxyNodeConfig(context.Context, *DistributeProxyNodeConfigReq) (*DistributeProxyNodeConfigReply, error)
	GetNodeKey(context.Context, *emptypb.Empty) (*GetNodeKeyReply, error)
	GetNodeStatus(context.Context, *GetNodeStatusReq) (*GetNodeStatusReply, error)
	GetNodeFunction(context.Context, *emptypb.Empty) (*GetNodeFunctionReply, error)
	// 平台下发用户信息
	DistributeUser(context.Context, *DistributeUserReq) (*DistributeUserReply, error)
	// 上报用户信息
	EscalationUser(context.Context, *EscalationUserReq) (*EscalationUserReply, error)
	// 更新用户中心的账号密码
	UpdateCenterUserPassword(context.Context, *UpdateCenterUserPasswordReq) (*UpdateCenterUserPasswordReply, error)
	// 注销账号
	CancelAccount(context.Context, *CancelAccountReq) (*CancelAccountReply, error)
	mustEmbedUnimplementedPlatformServer()
}

// UnimplementedPlatformServer must be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (UnimplementedPlatformServer) DistributeMasterNodeConfig(context.Context, *DistributeMasterNodeConfigReq) (*DistributeMasterNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeMasterNodeConfig not implemented")
}
func (UnimplementedPlatformServer) DistributeProxyNodeConfig(context.Context, *DistributeProxyNodeConfigReq) (*DistributeProxyNodeConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeProxyNodeConfig not implemented")
}
func (UnimplementedPlatformServer) GetNodeKey(context.Context, *emptypb.Empty) (*GetNodeKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeKey not implemented")
}
func (UnimplementedPlatformServer) GetNodeStatus(context.Context, *GetNodeStatusReq) (*GetNodeStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeStatus not implemented")
}
func (UnimplementedPlatformServer) GetNodeFunction(context.Context, *emptypb.Empty) (*GetNodeFunctionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeFunction not implemented")
}
func (UnimplementedPlatformServer) DistributeUser(context.Context, *DistributeUserReq) (*DistributeUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeUser not implemented")
}
func (UnimplementedPlatformServer) EscalationUser(context.Context, *EscalationUserReq) (*EscalationUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscalationUser not implemented")
}
func (UnimplementedPlatformServer) UpdateCenterUserPassword(context.Context, *UpdateCenterUserPasswordReq) (*UpdateCenterUserPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCenterUserPassword not implemented")
}
func (UnimplementedPlatformServer) CancelAccount(context.Context, *CancelAccountReq) (*CancelAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAccount not implemented")
}
func (UnimplementedPlatformServer) mustEmbedUnimplementedPlatformServer() {}

// UnsafePlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformServer will
// result in compilation errors.
type UnsafePlatformServer interface {
	mustEmbedUnimplementedPlatformServer()
}

func RegisterPlatformServer(s grpc.ServiceRegistrar, srv PlatformServer) {
	s.RegisterService(&Platform_ServiceDesc, srv)
}

func _Platform_DistributeMasterNodeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributeMasterNodeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DistributeMasterNodeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_DistributeMasterNodeConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DistributeMasterNodeConfig(ctx, req.(*DistributeMasterNodeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DistributeProxyNodeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributeProxyNodeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DistributeProxyNodeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_DistributeProxyNodeConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DistributeProxyNodeConfig(ctx, req.(*DistributeProxyNodeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetNodeKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetNodeKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_GetNodeKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetNodeKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_GetNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetNodeStatus(ctx, req.(*GetNodeStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetNodeFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetNodeFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_GetNodeFunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetNodeFunction(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DistributeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributeUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DistributeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_DistributeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DistributeUser(ctx, req.(*DistributeUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_EscalationUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EscalationUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).EscalationUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_EscalationUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).EscalationUser(ctx, req.(*EscalationUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_UpdateCenterUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCenterUserPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).UpdateCenterUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_UpdateCenterUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).UpdateCenterUserPassword(ctx, req.(*UpdateCenterUserPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_CancelAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).CancelAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_CancelAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).CancelAccount(ctx, req.(*CancelAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Platform_ServiceDesc is the grpc.ServiceDesc for Platform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Platform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.im_cloud.access.v2.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeMasterNodeConfig",
			Handler:    _Platform_DistributeMasterNodeConfig_Handler,
		},
		{
			MethodName: "DistributeProxyNodeConfig",
			Handler:    _Platform_DistributeProxyNodeConfig_Handler,
		},
		{
			MethodName: "GetNodeKey",
			Handler:    _Platform_GetNodeKey_Handler,
		},
		{
			MethodName: "GetNodeStatus",
			Handler:    _Platform_GetNodeStatus_Handler,
		},
		{
			MethodName: "GetNodeFunction",
			Handler:    _Platform_GetNodeFunction_Handler,
		},
		{
			MethodName: "DistributeUser",
			Handler:    _Platform_DistributeUser_Handler,
		},
		{
			MethodName: "EscalationUser",
			Handler:    _Platform_EscalationUser_Handler,
		},
		{
			MethodName: "UpdateCenterUserPassword",
			Handler:    _Platform_UpdateCenterUserPassword_Handler,
		},
		{
			MethodName: "CancelAccount",
			Handler:    _Platform_CancelAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im_cloud/access/v1/platform.proto",
}
