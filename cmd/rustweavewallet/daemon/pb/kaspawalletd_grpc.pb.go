// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: rustweavewalletd.proto

package pb

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

// RustweavewalletdClient is the client API for Rustweavewalletd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RustweavewalletdClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type rustweavewalletdClient struct {
	cc grpc.ClientConnInterface
}

func NewRustweavewalletdClient(cc grpc.ClientConnInterface) RustweavewalletdClient {
	return &rustweavewalletdClient{cc}
}

func (c *rustweavewalletdClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error) {
	out := new(GetExternalSpendableUTXOsResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/GetExternalSpendableUTXOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error) {
	out := new(CreateUnsignedTransactionsResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/CreateUnsignedTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error) {
	out := new(ShowAddressesResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/ShowAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/NewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rustweavewalletdClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/rustweavewalletd.rustweavewalletd/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RustweavewalletdServer is the server API for Rustweavewalletd service.
// All implementations must embed UnimplementedRustweavewalletdServer
// for forward compatibility
type RustweavewalletdServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedRustweavewalletdServer()
}

// UnimplementedRustweavewalletdServer must be embedded to have forward compatible implementations.
type UnimplementedRustweavewalletdServer struct {
}

func (UnimplementedRustweavewalletdServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedRustweavewalletdServer) GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalSpendableUTXOs not implemented")
}
func (UnimplementedRustweavewalletdServer) CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnsignedTransactions not implemented")
}
func (UnimplementedRustweavewalletdServer) ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAddresses not implemented")
}
func (UnimplementedRustweavewalletdServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedRustweavewalletdServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedRustweavewalletdServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedRustweavewalletdServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedRustweavewalletdServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedRustweavewalletdServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedRustweavewalletdServer) mustEmbedUnimplementedRustweavewalletdServer() {}

// UnsafeRustweavewalletdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RustweavewalletdServer will
// result in compilation errors.
type UnsafeRustweavewalletdServer interface {
	mustEmbedUnimplementedRustweavewalletdServer()
}

func RegisterRustweavewalletdServer(s grpc.ServiceRegistrar, srv RustweavewalletdServer) {
	s.RegisterService(&Rustweavewalletd_ServiceDesc, srv)
}

func _Rustweavewalletd_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_GetExternalSpendableUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalSpendableUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).GetExternalSpendableUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/GetExternalSpendableUTXOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).GetExternalSpendableUTXOs(ctx, req.(*GetExternalSpendableUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_CreateUnsignedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnsignedTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).CreateUnsignedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/CreateUnsignedTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).CreateUnsignedTransactions(ctx, req.(*CreateUnsignedTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_ShowAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).ShowAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/ShowAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).ShowAddresses(ctx, req.(*ShowAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rustweavewalletd_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RustweavewalletdServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rustweavewalletd.rustweavewalletd/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RustweavewalletdServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rustweavewalletd_ServiceDesc is the grpc.ServiceDesc for Rustweavewalletd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rustweavewalletd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rustweavewalletd.rustweavewalletd",
	HandlerType: (*RustweavewalletdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Rustweavewalletd_GetBalance_Handler,
		},
		{
			MethodName: "GetExternalSpendableUTXOs",
			Handler:    _Rustweavewalletd_GetExternalSpendableUTXOs_Handler,
		},
		{
			MethodName: "CreateUnsignedTransactions",
			Handler:    _Rustweavewalletd_CreateUnsignedTransactions_Handler,
		},
		{
			MethodName: "ShowAddresses",
			Handler:    _Rustweavewalletd_ShowAddresses_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Rustweavewalletd_NewAddress_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Rustweavewalletd_Shutdown_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Rustweavewalletd_Broadcast_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Rustweavewalletd_Send_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Rustweavewalletd_Sign_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Rustweavewalletd_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rustweavewalletd.proto",
}