// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// BlogAdminServiceClient is the client API for BlogAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogAdminServiceClient interface {
	//  删除用户 - id
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserReply, error)
	// 根据id查询用户
	SearchUserByID(ctx context.Context, in *SearchUserByIDReq, opts ...grpc.CallOption) (*SearchUserByIDReply, error)
	// 添加用户
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error)
	// 更新用户
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserReply, error)
	// 分页查询用户列表
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListReply, error)
	// 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	// 用户退出
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error)
	// 注册用户
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error)
}

type blogAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogAdminServiceClient(cc grpc.ClientConnInterface) BlogAdminServiceClient {
	return &blogAdminServiceClient{cc}
}

func (c *blogAdminServiceClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) SearchUserByID(ctx context.Context, in *SearchUserByIDReq, opts ...grpc.CallOption) (*SearchUserByIDReply, error) {
	out := new(SearchUserByIDReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/SearchUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListReply, error) {
	out := new(GetUserListReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogAdminServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/blog.admin.v1.BlogAdminService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogAdminServiceServer is the server API for BlogAdminService service.
// All implementations must embed UnimplementedBlogAdminServiceServer
// for forward compatibility
type BlogAdminServiceServer interface {
	//  删除用户 - id
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserReply, error)
	// 根据id查询用户
	SearchUserByID(context.Context, *SearchUserByIDReq) (*SearchUserByIDReply, error)
	// 添加用户
	CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error)
	// 更新用户
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserReply, error)
	// 分页查询用户列表
	GetUserList(context.Context, *GetUserListReq) (*GetUserListReply, error)
	// 登录
	Login(context.Context, *LoginReq) (*LoginReply, error)
	// 用户退出
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	// 注册用户
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	mustEmbedUnimplementedBlogAdminServiceServer()
}

// UnimplementedBlogAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogAdminServiceServer struct {
}

func (UnimplementedBlogAdminServiceServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedBlogAdminServiceServer) SearchUserByID(context.Context, *SearchUserByIDReq) (*SearchUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserByID not implemented")
}
func (UnimplementedBlogAdminServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBlogAdminServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedBlogAdminServiceServer) GetUserList(context.Context, *GetUserListReq) (*GetUserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedBlogAdminServiceServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBlogAdminServiceServer) Logout(context.Context, *LogoutReq) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedBlogAdminServiceServer) Register(context.Context, *RegisterReq) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBlogAdminServiceServer) mustEmbedUnimplementedBlogAdminServiceServer() {}

// UnsafeBlogAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogAdminServiceServer will
// result in compilation errors.
type UnsafeBlogAdminServiceServer interface {
	mustEmbedUnimplementedBlogAdminServiceServer()
}

func RegisterBlogAdminServiceServer(s grpc.ServiceRegistrar, srv BlogAdminServiceServer) {
	s.RegisterService(&BlogAdminService_ServiceDesc, srv)
}

func _BlogAdminService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_SearchUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).SearchUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/SearchUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).SearchUserByID(ctx, req.(*SearchUserByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogAdminService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogAdminServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.admin.v1.BlogAdminService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogAdminServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogAdminService_ServiceDesc is the grpc.ServiceDesc for BlogAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.admin.v1.BlogAdminService",
	HandlerType: (*BlogAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUser",
			Handler:    _BlogAdminService_DeleteUser_Handler,
		},
		{
			MethodName: "SearchUserByID",
			Handler:    _BlogAdminService_SearchUserByID_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _BlogAdminService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _BlogAdminService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _BlogAdminService_GetUserList_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BlogAdminService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _BlogAdminService_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _BlogAdminService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog_admin.proto",
}
