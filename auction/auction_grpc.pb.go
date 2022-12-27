// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: auction/auction.proto

package auction

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

// ConnectServiceClient is the client API for ConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectServiceClient interface {
	FinishConnect(ctx context.Context, in *PrimaryNode, opts ...grpc.CallOption) (*BackupDetails, error)
	AddBackup(ctx context.Context, in *BackupJoin, opts ...grpc.CallOption) (*Void, error)
	RemoveBackup(ctx context.Context, in *BackupLeave, opts ...grpc.CallOption) (*Void, error)
}

type connectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectServiceClient(cc grpc.ClientConnInterface) ConnectServiceClient {
	return &connectServiceClient{cc}
}

func (c *connectServiceClient) FinishConnect(ctx context.Context, in *PrimaryNode, opts ...grpc.CallOption) (*BackupDetails, error) {
	out := new(BackupDetails)
	err := c.cc.Invoke(ctx, "/auctionSystem.ConnectService/FinishConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) AddBackup(ctx context.Context, in *BackupJoin, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auctionSystem.ConnectService/AddBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) RemoveBackup(ctx context.Context, in *BackupLeave, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auctionSystem.ConnectService/RemoveBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServiceServer is the server API for ConnectService service.
// All implementations must embed UnimplementedConnectServiceServer
// for forward compatibility
type ConnectServiceServer interface {
	FinishConnect(context.Context, *PrimaryNode) (*BackupDetails, error)
	AddBackup(context.Context, *BackupJoin) (*Void, error)
	RemoveBackup(context.Context, *BackupLeave) (*Void, error)
	mustEmbedUnimplementedConnectServiceServer()
}

// UnimplementedConnectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectServiceServer struct {
}

func (UnimplementedConnectServiceServer) FinishConnect(context.Context, *PrimaryNode) (*BackupDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishConnect not implemented")
}
func (UnimplementedConnectServiceServer) AddBackup(context.Context, *BackupJoin) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBackup not implemented")
}
func (UnimplementedConnectServiceServer) RemoveBackup(context.Context, *BackupLeave) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBackup not implemented")
}
func (UnimplementedConnectServiceServer) mustEmbedUnimplementedConnectServiceServer() {}

// UnsafeConnectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectServiceServer will
// result in compilation errors.
type UnsafeConnectServiceServer interface {
	mustEmbedUnimplementedConnectServiceServer()
}

func RegisterConnectServiceServer(s grpc.ServiceRegistrar, srv ConnectServiceServer) {
	s.RegisterService(&ConnectService_ServiceDesc, srv)
}

func _ConnectService_FinishConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).FinishConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.ConnectService/FinishConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).FinishConnect(ctx, req.(*PrimaryNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_AddBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupJoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).AddBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.ConnectService/AddBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).AddBackup(ctx, req.(*BackupJoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_RemoveBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupLeave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).RemoveBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.ConnectService/RemoveBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).RemoveBackup(ctx, req.(*BackupLeave))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectService_ServiceDesc is the grpc.ServiceDesc for ConnectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auctionSystem.ConnectService",
	HandlerType: (*ConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FinishConnect",
			Handler:    _ConnectService_FinishConnect_Handler,
		},
		{
			MethodName: "AddBackup",
			Handler:    _ConnectService_AddBackup_Handler,
		},
		{
			MethodName: "RemoveBackup",
			Handler:    _ConnectService_RemoveBackup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/auction.proto",
}

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionServiceClient interface {
	AuctionStarted(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error)
	GetResult(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error)
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) AuctionStarted(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auctionSystem.AuctionService/AuctionStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error) {
	out := new(BidAck)
	err := c.cc.Invoke(ctx, "/auctionSystem.AuctionService/SendBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) GetResult(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error) {
	out := new(Outcome)
	err := c.cc.Invoke(ctx, "/auctionSystem.AuctionService/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auctionSystem.AuctionService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility
type AuctionServiceServer interface {
	AuctionStarted(context.Context, *Void) (*Void, error)
	SendBid(context.Context, *Bid) (*BidAck, error)
	GetResult(context.Context, *Void) (*Outcome, error)
	Ping(context.Context, *Void) (*Void, error)
	mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServiceServer struct {
}

func (UnimplementedAuctionServiceServer) AuctionStarted(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuctionStarted not implemented")
}
func (UnimplementedAuctionServiceServer) SendBid(context.Context, *Bid) (*BidAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBid not implemented")
}
func (UnimplementedAuctionServiceServer) GetResult(context.Context, *Void) (*Outcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedAuctionServiceServer) Ping(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_AuctionStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).AuctionStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.AuctionService/AuctionStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).AuctionStarted(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_SendBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).SendBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.AuctionService/SendBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).SendBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.AuctionService/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).GetResult(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.AuctionService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auctionSystem.AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuctionStarted",
			Handler:    _AuctionService_AuctionStarted_Handler,
		},
		{
			MethodName: "SendBid",
			Handler:    _AuctionService_SendBid_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _AuctionService_GetResult_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _AuctionService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/auction.proto",
}

// FrontendServiceClient is the client API for FrontendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrontendServiceClient interface {
	SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error)
	GetResult(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error)
	SetPrimaryNode(ctx context.Context, in *PrimaryNode, opts ...grpc.CallOption) (*Void, error)
}

type frontendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontendServiceClient(cc grpc.ClientConnInterface) FrontendServiceClient {
	return &frontendServiceClient{cc}
}

func (c *frontendServiceClient) SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error) {
	out := new(BidAck)
	err := c.cc.Invoke(ctx, "/auctionSystem.FrontendService/SendBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendServiceClient) GetResult(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error) {
	out := new(Outcome)
	err := c.cc.Invoke(ctx, "/auctionSystem.FrontendService/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendServiceClient) SetPrimaryNode(ctx context.Context, in *PrimaryNode, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auctionSystem.FrontendService/SetPrimaryNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServiceServer is the server API for FrontendService service.
// All implementations must embed UnimplementedFrontendServiceServer
// for forward compatibility
type FrontendServiceServer interface {
	SendBid(context.Context, *Bid) (*BidAck, error)
	GetResult(context.Context, *Void) (*Outcome, error)
	SetPrimaryNode(context.Context, *PrimaryNode) (*Void, error)
	mustEmbedUnimplementedFrontendServiceServer()
}

// UnimplementedFrontendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFrontendServiceServer struct {
}

func (UnimplementedFrontendServiceServer) SendBid(context.Context, *Bid) (*BidAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBid not implemented")
}
func (UnimplementedFrontendServiceServer) GetResult(context.Context, *Void) (*Outcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedFrontendServiceServer) SetPrimaryNode(context.Context, *PrimaryNode) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryNode not implemented")
}
func (UnimplementedFrontendServiceServer) mustEmbedUnimplementedFrontendServiceServer() {}

// UnsafeFrontendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrontendServiceServer will
// result in compilation errors.
type UnsafeFrontendServiceServer interface {
	mustEmbedUnimplementedFrontendServiceServer()
}

func RegisterFrontendServiceServer(s grpc.ServiceRegistrar, srv FrontendServiceServer) {
	s.RegisterService(&FrontendService_ServiceDesc, srv)
}

func _FrontendService_SendBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).SendBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.FrontendService/SendBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).SendBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontendService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.FrontendService/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).GetResult(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontendService_SetPrimaryNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).SetPrimaryNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auctionSystem.FrontendService/SetPrimaryNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).SetPrimaryNode(ctx, req.(*PrimaryNode))
	}
	return interceptor(ctx, in, info, handler)
}

// FrontendService_ServiceDesc is the grpc.ServiceDesc for FrontendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrontendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auctionSystem.FrontendService",
	HandlerType: (*FrontendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBid",
			Handler:    _FrontendService_SendBid_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _FrontendService_GetResult_Handler,
		},
		{
			MethodName: "SetPrimaryNode",
			Handler:    _FrontendService_SetPrimaryNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/auction.proto",
}
