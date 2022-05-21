// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package durudexv1

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// User Sign Up.
	UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error)
	// User Sign In.
	UserSignIn(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error)
	// User Sign Out.
	UserSignOut(ctx context.Context, in *UserSignOutRequest, opts ...grpc.CallOption) (*UserSignOutResponse, error)
	// Refresh user authentication token.
	RefreshUserToken(ctx context.Context, in *RefreshUserTokenRequest, opts ...grpc.CallOption) (*RefreshUserTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error) {
	out := new(UserSignUpResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.AuthService/UserSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserSignIn(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error) {
	out := new(UserSignInResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.AuthService/UserSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserSignOut(ctx context.Context, in *UserSignOutRequest, opts ...grpc.CallOption) (*UserSignOutResponse, error) {
	out := new(UserSignOutResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.AuthService/UserSignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshUserToken(ctx context.Context, in *RefreshUserTokenRequest, opts ...grpc.CallOption) (*RefreshUserTokenResponse, error) {
	out := new(RefreshUserTokenResponse)
	err := c.cc.Invoke(ctx, "/durudex.v1.AuthService/RefreshUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// User Sign Up.
	UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error)
	// User Sign In.
	UserSignIn(context.Context, *UserSignInRequest) (*UserSignInResponse, error)
	// User Sign Out.
	UserSignOut(context.Context, *UserSignOutRequest) (*UserSignOutResponse, error)
	// Refresh user authentication token.
	RefreshUserToken(context.Context, *RefreshUserTokenRequest) (*RefreshUserTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedAuthServiceServer) UserSignIn(context.Context, *UserSignInRequest) (*UserSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignIn not implemented")
}
func (UnimplementedAuthServiceServer) UserSignOut(context.Context, *UserSignOutRequest) (*UserSignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignOut not implemented")
}
func (UnimplementedAuthServiceServer) RefreshUserToken(context.Context, *RefreshUserTokenRequest) (*RefreshUserTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshUserToken not implemented")
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

func _AuthService_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.AuthService/UserSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignUp(ctx, req.(*UserSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.AuthService/UserSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignIn(ctx, req.(*UserSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserSignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.AuthService/UserSignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignOut(ctx, req.(*UserSignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshUserTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.v1.AuthService/RefreshUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshUserToken(ctx, req.(*RefreshUserTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "durudex.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _AuthService_UserSignUp_Handler,
		},
		{
			MethodName: "UserSignIn",
			Handler:    _AuthService_UserSignIn_Handler,
		},
		{
			MethodName: "UserSignOut",
			Handler:    _AuthService_UserSignOut_Handler,
		},
		{
			MethodName: "RefreshUserToken",
			Handler:    _AuthService_RefreshUserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "durudex/v1/auth.proto",
}