// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: auth_v1/auth_service.proto

package auth_v1

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
	AuthService_Registration_FullMethodName              = "/auth_v1.AuthService/registration"
	AuthService_Login_FullMethodName                     = "/auth_v1.AuthService/login"
	AuthService_Refresh_FullMethodName                   = "/auth_v1.AuthService/refresh"
	AuthService_Logout_FullMethodName                    = "/auth_v1.AuthService/logout"
	AuthService_LinkActivate_FullMethodName              = "/auth_v1.AuthService/linkActivate"
	AuthService_ChangePasswordWithAuth_FullMethodName    = "/auth_v1.AuthService/changePasswordWithAuth"
	AuthService_ChangePasswordWithoutAuth_FullMethodName = "/auth_v1.AuthService/changePasswordWithoutAuth"
	AuthService_CheckAuth_FullMethodName                 = "/auth_v1.AuthService/checkAuth"
	AuthService_ConfirmChangePassword_FullMethodName     = "/auth_v1.AuthService/confirmChangePassword"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Refresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*AuthResponse, error)
	Logout(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Empty, error)
	LinkActivate(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error)
	ChangePasswordWithAuth(ctx context.Context, in *ChangePasswordRequestWithAuth, opts ...grpc.CallOption) (*Empty, error)
	ChangePasswordWithoutAuth(ctx context.Context, in *ChangePasswordRequestWithoutAuth, opts ...grpc.CallOption) (*Empty, error)
	CheckAuth(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*JwtUser, error)
	ConfirmChangePassword(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Registration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LinkActivate(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_LinkActivate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePasswordWithAuth(ctx context.Context, in *ChangePasswordRequestWithAuth, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_ChangePasswordWithAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePasswordWithoutAuth(ctx context.Context, in *ChangePasswordRequestWithoutAuth, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_ChangePasswordWithoutAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckAuth(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*JwtUser, error) {
	out := new(JwtUser)
	err := c.cc.Invoke(ctx, AuthService_CheckAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmChangePassword(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_ConfirmChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Registration(context.Context, *RegistrationRequest) (*AuthResponse, error)
	Login(context.Context, *LoginRequest) (*AuthResponse, error)
	Refresh(context.Context, *RefreshToken) (*AuthResponse, error)
	Logout(context.Context, *RefreshToken) (*Empty, error)
	LinkActivate(context.Context, *Link) (*Empty, error)
	ChangePasswordWithAuth(context.Context, *ChangePasswordRequestWithAuth) (*Empty, error)
	ChangePasswordWithoutAuth(context.Context, *ChangePasswordRequestWithoutAuth) (*Empty, error)
	CheckAuth(context.Context, *AccessToken) (*JwtUser, error)
	ConfirmChangePassword(context.Context, *Link) (*Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Registration(context.Context, *RegistrationRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Refresh(context.Context, *RefreshToken) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *RefreshToken) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) LinkActivate(context.Context, *Link) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkActivate not implemented")
}
func (UnimplementedAuthServiceServer) ChangePasswordWithAuth(context.Context, *ChangePasswordRequestWithAuth) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordWithAuth not implemented")
}
func (UnimplementedAuthServiceServer) ChangePasswordWithoutAuth(context.Context, *ChangePasswordRequestWithoutAuth) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordWithoutAuth not implemented")
}
func (UnimplementedAuthServiceServer) CheckAuth(context.Context, *AccessToken) (*JwtUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedAuthServiceServer) ConfirmChangePassword(context.Context, *Link) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmChangePassword not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Registration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LinkActivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LinkActivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_LinkActivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LinkActivate(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePasswordWithAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequestWithAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePasswordWithAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePasswordWithAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePasswordWithAuth(ctx, req.(*ChangePasswordRequestWithAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePasswordWithoutAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequestWithoutAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePasswordWithoutAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePasswordWithoutAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePasswordWithoutAuth(ctx, req.(*ChangePasswordRequestWithoutAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckAuth(ctx, req.(*AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ConfirmChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmChangePassword(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registration",
			Handler:    _AuthService_Registration_Handler,
		},
		{
			MethodName: "login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "linkActivate",
			Handler:    _AuthService_LinkActivate_Handler,
		},
		{
			MethodName: "changePasswordWithAuth",
			Handler:    _AuthService_ChangePasswordWithAuth_Handler,
		},
		{
			MethodName: "changePasswordWithoutAuth",
			Handler:    _AuthService_ChangePasswordWithoutAuth_Handler,
		},
		{
			MethodName: "checkAuth",
			Handler:    _AuthService_CheckAuth_Handler,
		},
		{
			MethodName: "confirmChangePassword",
			Handler:    _AuthService_ConfirmChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_v1/auth_service.proto",
}
