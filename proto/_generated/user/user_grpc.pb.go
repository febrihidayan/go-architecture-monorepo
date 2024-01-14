// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user/user.proto

package user

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

const (
	UserServices_CreateUser_FullMethodName              = "/user.UserServices/CreateUser"
	UserServices_FindUser_FullMethodName                = "/user.UserServices/FindUser"
	UserServices_UpdateEmailVerifiedUser_FullMethodName = "/user.UserServices/UpdateEmailVerifiedUser"
)

// UserServicesClient is the client API for UserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicesClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error)
	UpdateEmailVerifiedUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicesClient(cc grpc.ClientConnInterface) UserServicesClient {
	return &userServicesClient{cc}
}

func (c *userServicesClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserServices_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error) {
	out := new(FindUserResponse)
	err := c.cc.Invoke(ctx, UserServices_FindUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UpdateEmailVerifiedUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserServices_UpdateEmailVerifiedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicesServer is the server API for UserServices service.
// All implementations should embed UnimplementedUserServicesServer
// for forward compatibility
type UserServicesServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error)
	UpdateEmailVerifiedUser(context.Context, *User) (*emptypb.Empty, error)
}

// UnimplementedUserServicesServer should be embedded to have forward compatible implementations.
type UnimplementedUserServicesServer struct {
}

func (UnimplementedUserServicesServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServicesServer) FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedUserServicesServer) UpdateEmailVerifiedUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailVerifiedUser not implemented")
}

// UnsafeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicesServer will
// result in compilation errors.
type UnsafeUserServicesServer interface {
	mustEmbedUnimplementedUserServicesServer()
}

func RegisterUserServicesServer(s grpc.ServiceRegistrar, srv UserServicesServer) {
	s.RegisterService(&UserServices_ServiceDesc, srv)
}

func _UserServices_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_FindUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).FindUser(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UpdateEmailVerifiedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UpdateEmailVerifiedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UpdateEmailVerifiedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UpdateEmailVerifiedUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServices_ServiceDesc is the grpc.ServiceDesc for UserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserServices",
	HandlerType: (*UserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserServices_CreateUser_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _UserServices_FindUser_Handler,
		},
		{
			MethodName: "UpdateEmailVerifiedUser",
			Handler:    _UserServices_UpdateEmailVerifiedUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
