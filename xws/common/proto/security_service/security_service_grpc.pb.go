// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package security

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	TwoFactorAuthentication(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SetupOTP(ctx context.Context, in *SetupOTPRequest, opts ...grpc.CallOption) (*SetupOTPResponse, error)
	PasswordlessLogin(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	RecoverPassword(ctx context.Context, in *RecoverPasswordRequest, opts ...grpc.CallOption) (*RecoverPasswordResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
}

type securityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityServiceClient(cc grpc.ClientConnInterface) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) TwoFactorAuthentication(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/TwoFactorAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) SetupOTP(ctx context.Context, in *SetupOTPRequest, opts ...grpc.CallOption) (*SetupOTPResponse, error) {
	out := new(SetupOTPResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/SetupOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) PasswordlessLogin(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/PasswordlessLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/security.SecurityService/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) RecoverPassword(ctx context.Context, in *RecoverPasswordRequest, opts ...grpc.CallOption) (*RecoverPasswordResponse, error) {
	out := new(RecoverPasswordResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/RecoverPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/security.SecurityService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
// All implementations must embed UnimplementedSecurityServiceServer
// for forward compatibility
type SecurityServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	TwoFactorAuthentication(context.Context, *PasswordlessLoginRequest) (*LoginResponse, error)
	SetupOTP(context.Context, *SetupOTPRequest) (*SetupOTPResponse, error)
	PasswordlessLogin(context.Context, *PasswordlessLoginRequest) (*LoginResponse, error)
	VerifyUser(context.Context, *VerifyUserRequest) (*httpbody.HttpBody, error)
	RecoverPassword(context.Context, *RecoverPasswordRequest) (*RecoverPasswordResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	mustEmbedUnimplementedSecurityServiceServer()
}

// UnimplementedSecurityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (*UnimplementedSecurityServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSecurityServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedSecurityServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedSecurityServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedSecurityServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedSecurityServiceServer) TwoFactorAuthentication(context.Context, *PasswordlessLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwoFactorAuthentication not implemented")
}
func (*UnimplementedSecurityServiceServer) SetupOTP(context.Context, *SetupOTPRequest) (*SetupOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupOTP not implemented")
}
func (*UnimplementedSecurityServiceServer) PasswordlessLogin(context.Context, *PasswordlessLoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordlessLogin not implemented")
}
func (*UnimplementedSecurityServiceServer) VerifyUser(context.Context, *VerifyUserRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (*UnimplementedSecurityServiceServer) RecoverPassword(context.Context, *RecoverPasswordRequest) (*RecoverPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverPassword not implemented")
}
func (*UnimplementedSecurityServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (*UnimplementedSecurityServiceServer) mustEmbedUnimplementedSecurityServiceServer() {}

func RegisterSecurityServiceServer(s *grpc.Server, srv SecurityServiceServer) {
	s.RegisterService(&_SecurityService_serviceDesc, srv)
}

func _SecurityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_TwoFactorAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordlessLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).TwoFactorAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/TwoFactorAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).TwoFactorAuthentication(ctx, req.(*PasswordlessLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_SetupOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).SetupOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/SetupOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).SetupOTP(ctx, req.(*SetupOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_PasswordlessLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordlessLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).PasswordlessLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/PasswordlessLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).PasswordlessLogin(ctx, req.(*PasswordlessLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).VerifyUser(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_RecoverPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).RecoverPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/RecoverPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).RecoverPassword(ctx, req.(*RecoverPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "security.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SecurityService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _SecurityService_GetAll_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _SecurityService_Register_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SecurityService_Update_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SecurityService_Login_Handler,
		},
		{
			MethodName: "TwoFactorAuthentication",
			Handler:    _SecurityService_TwoFactorAuthentication_Handler,
		},
		{
			MethodName: "SetupOTP",
			Handler:    _SecurityService_SetupOTP_Handler,
		},
		{
			MethodName: "PasswordlessLogin",
			Handler:    _SecurityService_PasswordlessLogin_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _SecurityService_VerifyUser_Handler,
		},
		{
			MethodName: "RecoverPassword",
			Handler:    _SecurityService_RecoverPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _SecurityService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/security_service/security_service.proto",
}
