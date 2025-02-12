// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: platform/tenant/v2/tenant.proto

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
	Tenant_Application_FullMethodName             = "/api.platform.tenant.v2.tenant/Application"
	Tenant_Create_FullMethodName                  = "/api.platform.tenant.v2.tenant/Create"
	Tenant_Update_FullMethodName                  = "/api.platform.tenant.v2.tenant/Update"
	Tenant_BindingPhone_FullMethodName            = "/api.platform.tenant.v2.tenant/BindingPhone"
	Tenant_Delete_FullMethodName                  = "/api.platform.tenant.v2.tenant/Delete"
	Tenant_FindByID_FullMethodName                = "/api.platform.tenant.v2.tenant/FindByID"
	Tenant_WebsitePath_FullMethodName             = "/api.platform.tenant.v2.tenant/WebsitePath"
	Tenant_GetList_FullMethodName                 = "/api.platform.tenant.v2.tenant/GetList"
	Tenant_FindByAssociatedUsersID_FullMethodName = "/api.platform.tenant.v2.tenant/FindByAssociatedUsersID"
	Tenant_WebsiteApplication_FullMethodName      = "/api.platform.tenant.v2.tenant/WebsiteApplication"
)

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	// 套餐申请
	Application(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationReply, error)
	// 创建租户
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	// 更新租户
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	// 绑定手机号
	BindingPhone(ctx context.Context, in *BindingPhoneRequest, opts ...grpc.CallOption) (*BindingPhoneReply, error)
	// 删除租户
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
	// 查询租户
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDReply, error)
	// 获取网站地址
	WebsitePath(ctx context.Context, in *WebsitePathRequest, opts ...grpc.CallOption) (*WebsitePathReply, error)
	// 获取租户列表
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListReply, error)
	// 根据创建租户id获取租户
	FindByAssociatedUsersID(ctx context.Context, in *FindByAssociatedUsersIDRequest, opts ...grpc.CallOption) (*FindByAssociatedUsersIDReply, error)
	// 网站申请
	WebsiteApplication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tenantClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantClient(cc grpc.ClientConnInterface) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) Application(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApplicationReply)
	err := c.cc.Invoke(ctx, Tenant_Application_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, Tenant_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, Tenant_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) BindingPhone(ctx context.Context, in *BindingPhoneRequest, opts ...grpc.CallOption) (*BindingPhoneReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BindingPhoneReply)
	err := c.cc.Invoke(ctx, Tenant_BindingPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReply)
	err := c.cc.Invoke(ctx, Tenant_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindByIDReply)
	err := c.cc.Invoke(ctx, Tenant_FindByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) WebsitePath(ctx context.Context, in *WebsitePathRequest, opts ...grpc.CallOption) (*WebsitePathReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebsitePathReply)
	err := c.cc.Invoke(ctx, Tenant_WebsitePath_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListReply)
	err := c.cc.Invoke(ctx, Tenant_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) FindByAssociatedUsersID(ctx context.Context, in *FindByAssociatedUsersIDRequest, opts ...grpc.CallOption) (*FindByAssociatedUsersIDReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindByAssociatedUsersIDReply)
	err := c.cc.Invoke(ctx, Tenant_FindByAssociatedUsersID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) WebsiteApplication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Tenant_WebsiteApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	// 套餐申请
	Application(context.Context, *ApplicationRequest) (*ApplicationReply, error)
	// 创建租户
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	// 更新租户
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	// 绑定手机号
	BindingPhone(context.Context, *BindingPhoneRequest) (*BindingPhoneReply, error)
	// 删除租户
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
	// 查询租户
	FindByID(context.Context, *FindByIDRequest) (*FindByIDReply, error)
	// 获取网站地址
	WebsitePath(context.Context, *WebsitePathRequest) (*WebsitePathReply, error)
	// 获取租户列表
	GetList(context.Context, *GetListRequest) (*GetListReply, error)
	// 根据创建租户id获取租户
	FindByAssociatedUsersID(context.Context, *FindByAssociatedUsersIDRequest) (*FindByAssociatedUsersIDReply, error)
	// 网站申请
	WebsiteApplication(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
}

func (UnimplementedTenantServer) Application(context.Context, *ApplicationRequest) (*ApplicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Application not implemented")
}
func (UnimplementedTenantServer) Create(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTenantServer) Update(context.Context, *UpdateRequest) (*UpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTenantServer) BindingPhone(context.Context, *BindingPhoneRequest) (*BindingPhoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindingPhone not implemented")
}
func (UnimplementedTenantServer) Delete(context.Context, *DeleteRequest) (*DeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTenantServer) FindByID(context.Context, *FindByIDRequest) (*FindByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedTenantServer) WebsitePath(context.Context, *WebsitePathRequest) (*WebsitePathReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebsitePath not implemented")
}
func (UnimplementedTenantServer) GetList(context.Context, *GetListRequest) (*GetListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTenantServer) FindByAssociatedUsersID(context.Context, *FindByAssociatedUsersIDRequest) (*FindByAssociatedUsersIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByAssociatedUsersID not implemented")
}
func (UnimplementedTenantServer) WebsiteApplication(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebsiteApplication not implemented")
}
func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_Application_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Application(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_Application_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Application(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_BindingPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindingPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).BindingPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_BindingPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).BindingPhone(ctx, req.(*BindingPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_FindByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_WebsitePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsitePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).WebsitePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_WebsitePath_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).WebsitePath(ctx, req.(*WebsitePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_FindByAssociatedUsersID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByAssociatedUsersIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).FindByAssociatedUsersID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_FindByAssociatedUsersID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).FindByAssociatedUsersID(ctx, req.(*FindByAssociatedUsersIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_WebsiteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).WebsiteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_WebsiteApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).WebsiteApplication(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.platform.tenant.v2.tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Application",
			Handler:    _Tenant_Application_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Tenant_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tenant_Update_Handler,
		},
		{
			MethodName: "BindingPhone",
			Handler:    _Tenant_BindingPhone_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Tenant_Delete_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _Tenant_FindByID_Handler,
		},
		{
			MethodName: "WebsitePath",
			Handler:    _Tenant_WebsitePath_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Tenant_GetList_Handler,
		},
		{
			MethodName: "FindByAssociatedUsersID",
			Handler:    _Tenant_FindByAssociatedUsersID_Handler,
		},
		{
			MethodName: "WebsiteApplication",
			Handler:    _Tenant_WebsiteApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform/tenant/v2/tenant.proto",
}
