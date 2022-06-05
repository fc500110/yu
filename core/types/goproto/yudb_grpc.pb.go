// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package goproto

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

// YuDBClient is the client API for YuDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YuDBClient interface {
	GetTxn(ctx context.Context, in *TxnHash, opts ...grpc.CallOption) (*TxnResponse, error)
	SetTxn(ctx context.Context, in *SignedTxn, opts ...grpc.CallOption) (*Err, error)
	GetTxns(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*TxnsResponse, error)
	SetTxns(ctx context.Context, in *TxnsRequest, opts ...grpc.CallOption) (*Err, error)
	GetEvents(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*EventsResponse, error)
	SetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*Err, error)
	GetErrors(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*ErrorsResponse, error)
	SetError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Err, error)
}

type yuDBClient struct {
	cc grpc.ClientConnInterface
}

func NewYuDBClient(cc grpc.ClientConnInterface) YuDBClient {
	return &yuDBClient{cc}
}

func (c *yuDBClient) GetTxn(ctx context.Context, in *TxnHash, opts ...grpc.CallOption) (*TxnResponse, error) {
	out := new(TxnResponse)
	err := c.cc.Invoke(ctx, "/YuDB/GetTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) SetTxn(ctx context.Context, in *SignedTxn, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/YuDB/SetTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) GetTxns(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*TxnsResponse, error) {
	out := new(TxnsResponse)
	err := c.cc.Invoke(ctx, "/YuDB/GetTxns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) SetTxns(ctx context.Context, in *TxnsRequest, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/YuDB/SetTxns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) GetEvents(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/YuDB/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) SetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/YuDB/SetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) GetErrors(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*ErrorsResponse, error) {
	out := new(ErrorsResponse)
	err := c.cc.Invoke(ctx, "/YuDB/GetErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yuDBClient) SetError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/YuDB/SetError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YuDBServer is the server API for YuDB service.
// All implementations must embed UnimplementedYuDBServer
// for forward compatibility
type YuDBServer interface {
	GetTxn(context.Context, *TxnHash) (*TxnResponse, error)
	SetTxn(context.Context, *SignedTxn) (*Err, error)
	GetTxns(context.Context, *BlockHash) (*TxnsResponse, error)
	SetTxns(context.Context, *TxnsRequest) (*Err, error)
	GetEvents(context.Context, *BlockHash) (*EventsResponse, error)
	SetEvents(context.Context, *EventsRequest) (*Err, error)
	GetErrors(context.Context, *BlockHash) (*ErrorsResponse, error)
	SetError(context.Context, *Error) (*Err, error)
	mustEmbedUnimplementedYuDBServer()
}

// UnimplementedYuDBServer must be embedded to have forward compatible implementations.
type UnimplementedYuDBServer struct {
}

func (UnimplementedYuDBServer) GetTxn(context.Context, *TxnHash) (*TxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxn not implemented")
}
func (UnimplementedYuDBServer) SetTxn(context.Context, *SignedTxn) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTxn not implemented")
}
func (UnimplementedYuDBServer) GetTxns(context.Context, *BlockHash) (*TxnsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxns not implemented")
}
func (UnimplementedYuDBServer) SetTxns(context.Context, *TxnsRequest) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTxns not implemented")
}
func (UnimplementedYuDBServer) GetEvents(context.Context, *BlockHash) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedYuDBServer) SetEvents(context.Context, *EventsRequest) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEvents not implemented")
}
func (UnimplementedYuDBServer) GetErrors(context.Context, *BlockHash) (*ErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetErrors not implemented")
}
func (UnimplementedYuDBServer) SetError(context.Context, *Error) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetError not implemented")
}
func (UnimplementedYuDBServer) mustEmbedUnimplementedYuDBServer() {}

// UnsafeYuDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YuDBServer will
// result in compilation errors.
type UnsafeYuDBServer interface {
	mustEmbedUnimplementedYuDBServer()
}

func RegisterYuDBServer(s grpc.ServiceRegistrar, srv YuDBServer) {
	s.RegisterService(&YuDB_ServiceDesc, srv)
}

func _YuDB_GetTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).GetTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/GetTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).GetTxn(ctx, req.(*TxnHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_SetTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedTxn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).SetTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/SetTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).SetTxn(ctx, req.(*SignedTxn))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_GetTxns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).GetTxns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/GetTxns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).GetTxns(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_SetTxns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).SetTxns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/SetTxns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).SetTxns(ctx, req.(*TxnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).GetEvents(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_SetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).SetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/SetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).SetEvents(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_GetErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).GetErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/GetErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).GetErrors(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _YuDB_SetError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Error)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YuDBServer).SetError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/YuDB/SetError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YuDBServer).SetError(ctx, req.(*Error))
	}
	return interceptor(ctx, in, info, handler)
}

// YuDB_ServiceDesc is the grpc.ServiceDesc for YuDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YuDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YuDB",
	HandlerType: (*YuDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxn",
			Handler:    _YuDB_GetTxn_Handler,
		},
		{
			MethodName: "SetTxn",
			Handler:    _YuDB_SetTxn_Handler,
		},
		{
			MethodName: "GetTxns",
			Handler:    _YuDB_GetTxns_Handler,
		},
		{
			MethodName: "SetTxns",
			Handler:    _YuDB_SetTxns_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _YuDB_GetEvents_Handler,
		},
		{
			MethodName: "SetEvents",
			Handler:    _YuDB_SetEvents_Handler,
		},
		{
			MethodName: "GetErrors",
			Handler:    _YuDB_GetErrors_Handler,
		},
		{
			MethodName: "SetError",
			Handler:    _YuDB_SetError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yudb.proto",
}
