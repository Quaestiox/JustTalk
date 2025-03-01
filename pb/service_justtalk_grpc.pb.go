// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: service_justtalk.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	JustTalk_SayHello_FullMethodName       = "/pb.JustTalk/SayHello"
	JustTalk_CreateUser_FullMethodName     = "/pb.JustTalk/CreateUser"
	JustTalk_LoginUser_FullMethodName      = "/pb.JustTalk/LoginUser"
	JustTalk_UpdateUser_FullMethodName     = "/pb.JustTalk/UpdateUser"
	JustTalk_DeleteUser_FullMethodName     = "/pb.JustTalk/DeleteUser"
	JustTalk_GetUserInfo_FullMethodName    = "/pb.JustTalk/GetUserInfo"
	JustTalk_SendMessage_FullMethodName    = "/pb.JustTalk/SendMessage"
	JustTalk_GetMessage_FullMethodName     = "/pb.JustTalk/GetMessage"
	JustTalk_ReceiveMessage_FullMethodName = "/pb.JustTalk/ReceiveMessage"
)

// JustTalkClient is the client API for JustTalk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JustTalkClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	SendMessage(ctx context.Context, in *CreateMsgRequest, opts ...grpc.CallOption) (*CreateMsgResponse, error)
	GetMessage(ctx context.Context, in *GetMsgRequest, opts ...grpc.CallOption) (*GetMsgResponse, error)
	ReceiveMessage(ctx context.Context, in *ReceiveMsgRequest, opts ...grpc.CallOption) (*ReceiveMsgResponse, error)
}

type justTalkClient struct {
	cc grpc.ClientConnInterface
}

func NewJustTalkClient(cc grpc.ClientConnInterface) JustTalkClient {
	return &justTalkClient{cc}
}

func (c *justTalkClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, JustTalk_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, JustTalk_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, JustTalk_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, JustTalk_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, JustTalk_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, JustTalk_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) SendMessage(ctx context.Context, in *CreateMsgRequest, opts ...grpc.CallOption) (*CreateMsgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMsgResponse)
	err := c.cc.Invoke(ctx, JustTalk_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) GetMessage(ctx context.Context, in *GetMsgRequest, opts ...grpc.CallOption) (*GetMsgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMsgResponse)
	err := c.cc.Invoke(ctx, JustTalk_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *justTalkClient) ReceiveMessage(ctx context.Context, in *ReceiveMsgRequest, opts ...grpc.CallOption) (*ReceiveMsgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReceiveMsgResponse)
	err := c.cc.Invoke(ctx, JustTalk_ReceiveMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JustTalkServer is the server API for JustTalk service.
// All implementations must embed UnimplementedJustTalkServer
// for forward compatibility.
type JustTalkServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	SendMessage(context.Context, *CreateMsgRequest) (*CreateMsgResponse, error)
	GetMessage(context.Context, *GetMsgRequest) (*GetMsgResponse, error)
	ReceiveMessage(context.Context, *ReceiveMsgRequest) (*ReceiveMsgResponse, error)
	mustEmbedUnimplementedJustTalkServer()
}

// UnimplementedJustTalkServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJustTalkServer struct{}

func (UnimplementedJustTalkServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedJustTalkServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedJustTalkServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedJustTalkServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedJustTalkServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedJustTalkServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedJustTalkServer) SendMessage(context.Context, *CreateMsgRequest) (*CreateMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedJustTalkServer) GetMessage(context.Context, *GetMsgRequest) (*GetMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedJustTalkServer) ReceiveMessage(context.Context, *ReceiveMsgRequest) (*ReceiveMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedJustTalkServer) mustEmbedUnimplementedJustTalkServer() {}
func (UnimplementedJustTalkServer) testEmbeddedByValue()                  {}

// UnsafeJustTalkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JustTalkServer will
// result in compilation errors.
type UnsafeJustTalkServer interface {
	mustEmbedUnimplementedJustTalkServer()
}

func RegisterJustTalkServer(s grpc.ServiceRegistrar, srv JustTalkServer) {
	// If the following call pancis, it indicates UnimplementedJustTalkServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JustTalk_ServiceDesc, srv)
}

func _JustTalk_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).SendMessage(ctx, req.(*CreateMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).GetMessage(ctx, req.(*GetMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JustTalk_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JustTalkServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JustTalk_ReceiveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JustTalkServer).ReceiveMessage(ctx, req.(*ReceiveMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JustTalk_ServiceDesc is the grpc.ServiceDesc for JustTalk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JustTalk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.JustTalk",
	HandlerType: (*JustTalkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _JustTalk_SayHello_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _JustTalk_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _JustTalk_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _JustTalk_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _JustTalk_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _JustTalk_GetUserInfo_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _JustTalk_SendMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _JustTalk_GetMessage_Handler,
		},
		{
			MethodName: "ReceiveMessage",
			Handler:    _JustTalk_ReceiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_justtalk.proto",
}
