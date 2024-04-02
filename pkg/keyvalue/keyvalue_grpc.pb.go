// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pkg/keyvalue/keyvalue.proto

package keyvalue

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

// KeyvalueClient is the client API for Keyvalue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyvalueClient interface {
	Get(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
	Set(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
	Del(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
}

type keyvalueClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyvalueClient(cc grpc.ClientConnInterface) KeyvalueClient {
	return &keyvalueClient{cc}
}

func (c *keyvalueClient) Get(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, "/keyvalue.Keyvalue/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyvalueClient) Set(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, "/keyvalue.Keyvalue/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyvalueClient) Del(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, "/keyvalue.Keyvalue/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyvalueServer is the server API for Keyvalue service.
// All implementations must embed UnimplementedKeyvalueServer
// for forward compatibility
type KeyvalueServer interface {
	Get(context.Context, *KeyValue) (*KeyValue, error)
	Set(context.Context, *KeyValue) (*KeyValue, error)
	Del(context.Context, *KeyValue) (*KeyValue, error)
	mustEmbedUnimplementedKeyvalueServer()
}

// UnimplementedKeyvalueServer must be embedded to have forward compatible implementations.
type UnimplementedKeyvalueServer struct {
}

func (UnimplementedKeyvalueServer) Get(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeyvalueServer) Set(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKeyvalueServer) Del(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedKeyvalueServer) mustEmbedUnimplementedKeyvalueServer() {}

// UnsafeKeyvalueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyvalueServer will
// result in compilation errors.
type UnsafeKeyvalueServer interface {
	mustEmbedUnimplementedKeyvalueServer()
}

func RegisterKeyvalueServer(s grpc.ServiceRegistrar, srv KeyvalueServer) {
	s.RegisterService(&Keyvalue_ServiceDesc, srv)
}

func _Keyvalue_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyvalueServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvalue.Keyvalue/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyvalueServer).Get(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keyvalue_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyvalueServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvalue.Keyvalue/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyvalueServer).Set(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keyvalue_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyvalueServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvalue.Keyvalue/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyvalueServer).Del(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

// Keyvalue_ServiceDesc is the grpc.ServiceDesc for Keyvalue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keyvalue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyvalue.Keyvalue",
	HandlerType: (*KeyvalueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Keyvalue_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Keyvalue_Set_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Keyvalue_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/keyvalue/keyvalue.proto",
}
