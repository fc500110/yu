// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.13.0
// source: subscription.proto

package goproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x24, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x45, 0x6d, 0x69, 0x74, 0x12, 0x06,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_subscription_proto_goTypes = []interface{}{
	(*Bytes)(nil), // 0: Bytes
	(*Err)(nil),   // 1: Err
}
var file_subscription_proto_depIdxs = []int32{
	0, // 0: Subscription.Emit:input_type -> Bytes
	1, // 1: Subscription.Emit:output_type -> Err
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	file_base_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubscriptionClient is the client API for Subscription service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionClient interface {
	Emit(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*Err, error)
}

type subscriptionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionClient(cc grpc.ClientConnInterface) SubscriptionClient {
	return &subscriptionClient{cc}
}

func (c *subscriptionClient) Emit(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/Subscription/Emit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServer is the server API for Subscription service.
type SubscriptionServer interface {
	Emit(context.Context, *Bytes) (*Err, error)
}

// UnimplementedSubscriptionServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServer struct {
}

func (*UnimplementedSubscriptionServer) Emit(context.Context, *Bytes) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Emit not implemented")
}

func RegisterSubscriptionServer(s *grpc.Server, srv SubscriptionServer) {
	s.RegisterService(&_Subscription_serviceDesc, srv)
}

func _Subscription_Emit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).Emit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subscription/Emit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).Emit(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscription_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Subscription",
	HandlerType: (*SubscriptionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emit",
			Handler:    _Subscription_Emit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}