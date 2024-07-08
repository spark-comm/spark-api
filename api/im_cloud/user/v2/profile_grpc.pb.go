// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: im_cloud/user/v2/profile.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Profile_GetProfile_FullMethodName                 = "/api.im_cloud.user.v2.Profile/GetProfile"
	Profile_UpdateProfile_FullMethodName              = "/api.im_cloud.user.v2.Profile/UpdateProfile"
	Profile_SetOption_FullMethodName                  = "/api.im_cloud.user.v2.Profile/SetOption"
	Profile_GetOptionVal_FullMethodName               = "/api.im_cloud.user.v2.Profile/GetOptionVal"
	Profile_FindProfileByUserId_FullMethodName        = "/api.im_cloud.user.v2.Profile/FindProfileByUserId"
	Profile_FindFullProfileByUserId_FullMethodName    = "/api.im_cloud.user.v2.Profile/FindFullProfileByUserId"
	Profile_GetGlobalReceiveMessageOpt_FullMethodName = "/api.im_cloud.user.v2.Profile/GetGlobalReceiveMessageOpt"
	Profile_SearchProfile_FullMethodName              = "/api.im_cloud.user.v2.Profile/SearchProfile"
	Profile_GetAllUserID_FullMethodName               = "/api.im_cloud.user.v2.Profile/GetAllUserID"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	// 获取用户信息 存在
	GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileReply, error)
	// 更新用户新 存在
	UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*UpdateProfileReply, error)
	// 用户配置项 存在
	SetOption(ctx context.Context, in *SetOptionReq, opts ...grpc.CallOption) (*SetOptionReply, error)
	// 获取配置项的值 存在
	GetOptionVal(ctx context.Context, in *GetOptionValReq, opts ...grpc.CallOption) (*GetOptionValReply, error)
	// 根据ids获取用户信息 存在
	FindProfileByUserId(ctx context.Context, in *FindProfileByUserReq, opts ...grpc.CallOption) (*FindProfileByUserReply, error)
	// 根据ids获取完整的用户信息不过过滤删除和和租户 存在
	FindFullProfileByUserId(ctx context.Context, in *FindFullProfileByUserIdReq, opts ...grpc.CallOption) (*FindFullProfileByUserIdReply, error)
	// 获取用户全局消息配置项 存在
	GetGlobalReceiveMessageOpt(ctx context.Context, in *GetGlobalReceiveMessageOptReq, opts ...grpc.CallOption) (*GetGlobalReceiveMessageOptReply, error)
	// 搜索用户 存在
	SearchProfile(ctx context.Context, in *SearchProfileReq, opts ...grpc.CallOption) (*SearchProfileReply, error)
	// 获取所有的用户id 存在
	GetAllUserID(ctx context.Context, in *GetAllUserIDReq, opts ...grpc.CallOption) (*GetAllUserIDReply, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileReply)
	err := c.cc.Invoke(ctx, Profile_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*UpdateProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileReply)
	err := c.cc.Invoke(ctx, Profile_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) SetOption(ctx context.Context, in *SetOptionReq, opts ...grpc.CallOption) (*SetOptionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetOptionReply)
	err := c.cc.Invoke(ctx, Profile_SetOption_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetOptionVal(ctx context.Context, in *GetOptionValReq, opts ...grpc.CallOption) (*GetOptionValReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOptionValReply)
	err := c.cc.Invoke(ctx, Profile_GetOptionVal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FindProfileByUserId(ctx context.Context, in *FindProfileByUserReq, opts ...grpc.CallOption) (*FindProfileByUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindProfileByUserReply)
	err := c.cc.Invoke(ctx, Profile_FindProfileByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FindFullProfileByUserId(ctx context.Context, in *FindFullProfileByUserIdReq, opts ...grpc.CallOption) (*FindFullProfileByUserIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindFullProfileByUserIdReply)
	err := c.cc.Invoke(ctx, Profile_FindFullProfileByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetGlobalReceiveMessageOpt(ctx context.Context, in *GetGlobalReceiveMessageOptReq, opts ...grpc.CallOption) (*GetGlobalReceiveMessageOptReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGlobalReceiveMessageOptReply)
	err := c.cc.Invoke(ctx, Profile_GetGlobalReceiveMessageOpt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) SearchProfile(ctx context.Context, in *SearchProfileReq, opts ...grpc.CallOption) (*SearchProfileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProfileReply)
	err := c.cc.Invoke(ctx, Profile_SearchProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetAllUserID(ctx context.Context, in *GetAllUserIDReq, opts ...grpc.CallOption) (*GetAllUserIDReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUserIDReply)
	err := c.cc.Invoke(ctx, Profile_GetAllUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	// 获取用户信息 存在
	GetProfile(context.Context, *GetProfileReq) (*GetProfileReply, error)
	// 更新用户新 存在
	UpdateProfile(context.Context, *UpdateProfileReq) (*UpdateProfileReply, error)
	// 用户配置项 存在
	SetOption(context.Context, *SetOptionReq) (*SetOptionReply, error)
	// 获取配置项的值 存在
	GetOptionVal(context.Context, *GetOptionValReq) (*GetOptionValReply, error)
	// 根据ids获取用户信息 存在
	FindProfileByUserId(context.Context, *FindProfileByUserReq) (*FindProfileByUserReply, error)
	// 根据ids获取完整的用户信息不过过滤删除和和租户 存在
	FindFullProfileByUserId(context.Context, *FindFullProfileByUserIdReq) (*FindFullProfileByUserIdReply, error)
	// 获取用户全局消息配置项 存在
	GetGlobalReceiveMessageOpt(context.Context, *GetGlobalReceiveMessageOptReq) (*GetGlobalReceiveMessageOptReply, error)
	// 搜索用户 存在
	SearchProfile(context.Context, *SearchProfileReq) (*SearchProfileReply, error)
	// 获取所有的用户id 存在
	GetAllUserID(context.Context, *GetAllUserIDReq) (*GetAllUserIDReply, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) GetProfile(context.Context, *GetProfileReq) (*GetProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServer) UpdateProfile(context.Context, *UpdateProfileReq) (*UpdateProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServer) SetOption(context.Context, *SetOptionReq) (*SetOptionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOption not implemented")
}
func (UnimplementedProfileServer) GetOptionVal(context.Context, *GetOptionValReq) (*GetOptionValReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionVal not implemented")
}
func (UnimplementedProfileServer) FindProfileByUserId(context.Context, *FindProfileByUserReq) (*FindProfileByUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProfileByUserId not implemented")
}
func (UnimplementedProfileServer) FindFullProfileByUserId(context.Context, *FindFullProfileByUserIdReq) (*FindFullProfileByUserIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFullProfileByUserId not implemented")
}
func (UnimplementedProfileServer) GetGlobalReceiveMessageOpt(context.Context, *GetGlobalReceiveMessageOptReq) (*GetGlobalReceiveMessageOptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalReceiveMessageOpt not implemented")
}
func (UnimplementedProfileServer) SearchProfile(context.Context, *SearchProfileReq) (*SearchProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProfile not implemented")
}
func (UnimplementedProfileServer) GetAllUserID(context.Context, *GetAllUserIDReq) (*GetAllUserIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserID not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateProfile(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_SetOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).SetOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_SetOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).SetOption(ctx, req.(*SetOptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetOptionVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOptionValReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetOptionVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetOptionVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetOptionVal(ctx, req.(*GetOptionValReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FindProfileByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProfileByUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FindProfileByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_FindProfileByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FindProfileByUserId(ctx, req.(*FindProfileByUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FindFullProfileByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFullProfileByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FindFullProfileByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_FindFullProfileByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FindFullProfileByUserId(ctx, req.(*FindFullProfileByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetGlobalReceiveMessageOpt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalReceiveMessageOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetGlobalReceiveMessageOpt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetGlobalReceiveMessageOpt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetGlobalReceiveMessageOpt(ctx, req.(*GetGlobalReceiveMessageOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_SearchProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).SearchProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_SearchProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).SearchProfile(ctx, req.(*SearchProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetAllUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetAllUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetAllUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetAllUserID(ctx, req.(*GetAllUserIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.im_cloud.user.v2.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _Profile_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Profile_UpdateProfile_Handler,
		},
		{
			MethodName: "SetOption",
			Handler:    _Profile_SetOption_Handler,
		},
		{
			MethodName: "GetOptionVal",
			Handler:    _Profile_GetOptionVal_Handler,
		},
		{
			MethodName: "FindProfileByUserId",
			Handler:    _Profile_FindProfileByUserId_Handler,
		},
		{
			MethodName: "FindFullProfileByUserId",
			Handler:    _Profile_FindFullProfileByUserId_Handler,
		},
		{
			MethodName: "GetGlobalReceiveMessageOpt",
			Handler:    _Profile_GetGlobalReceiveMessageOpt_Handler,
		},
		{
			MethodName: "SearchProfile",
			Handler:    _Profile_SearchProfile_Handler,
		},
		{
			MethodName: "GetAllUserID",
			Handler:    _Profile_GetAllUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im_cloud/user/v2/profile.proto",
}
