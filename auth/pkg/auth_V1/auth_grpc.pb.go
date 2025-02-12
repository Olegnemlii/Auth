// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: auth.proto

package note_v1

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

// AuthV1Client is the client API for AuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthV1Client interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdatUser(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type authV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthV1Client(cc grpc.ClientConnInterface) AuthV1Client {
	return &authV1Client{cc}
}

func (c *authV1Client) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/auth_V1.AuthV1/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/auth_V1.AuthV1/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) UpdatUser(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/auth_V1.AuthV1/UpdatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/auth_V1.AuthV1/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthV1Server is the server API for AuthV1 service.
// All implementations must embed UnimplementedAuthV1Server
// for forward compatibility
type AuthV1Server interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdatUser(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedAuthV1Server()
}

// UnimplementedAuthV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthV1Server struct {
}

func (UnimplementedAuthV1Server) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthV1Server) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthV1Server) UpdatUser(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatUser not implemented")
}
func (UnimplementedAuthV1Server) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthV1Server) mustEmbedUnimplementedAuthV1Server() {}

// UnsafeAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthV1Server will
// result in compilation errors.
type UnsafeAuthV1Server interface {
	mustEmbedUnimplementedAuthV1Server()
}

func RegisterAuthV1Server(s grpc.ServiceRegistrar, srv AuthV1Server) {
	s.RegisterService(&AuthV1_ServiceDesc, srv)
}

func _AuthV1_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_V1.AuthV1/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_V1.AuthV1/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_UpdatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).UpdatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_V1.AuthV1/UpdatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).UpdatUser(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_V1.AuthV1/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthV1_ServiceDesc is the grpc.ServiceDesc for AuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_V1.AuthV1",
	HandlerType: (*AuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AuthV1_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthV1_GetUser_Handler,
		},
		{
			MethodName: "UpdatUser",
			Handler:    _AuthV1_UpdatUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthV1_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
