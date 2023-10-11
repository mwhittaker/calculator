// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: multiplier.proto

package calc

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
	Multiplier_Multiply_FullMethodName = "/Multiplier/Multiply"
)

// MultiplierClient is the client API for Multiplier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplierClient interface {
	Multiply(ctx context.Context, in *Multiplicands, opts ...grpc.CallOption) (*Product, error)
}

type multiplierClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplierClient(cc grpc.ClientConnInterface) MultiplierClient {
	return &multiplierClient{cc}
}

func (c *multiplierClient) Multiply(ctx context.Context, in *Multiplicands, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, Multiplier_Multiply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiplierServer is the server API for Multiplier service.
// All implementations must embed UnimplementedMultiplierServer
// for forward compatibility
type MultiplierServer interface {
	Multiply(context.Context, *Multiplicands) (*Product, error)
	mustEmbedUnimplementedMultiplierServer()
}

// UnimplementedMultiplierServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplierServer struct {
}

func (UnimplementedMultiplierServer) Multiply(context.Context, *Multiplicands) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedMultiplierServer) mustEmbedUnimplementedMultiplierServer() {}

// UnsafeMultiplierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplierServer will
// result in compilation errors.
type UnsafeMultiplierServer interface {
	mustEmbedUnimplementedMultiplierServer()
}

func RegisterMultiplierServer(s grpc.ServiceRegistrar, srv MultiplierServer) {
	s.RegisterService(&Multiplier_ServiceDesc, srv)
}

func _Multiplier_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Multiplicands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplierServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Multiplier_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplierServer).Multiply(ctx, req.(*Multiplicands))
	}
	return interceptor(ctx, in, info, handler)
}

// Multiplier_ServiceDesc is the grpc.ServiceDesc for Multiplier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Multiplier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Multiplier",
	HandlerType: (*MultiplierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _Multiplier_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiplier.proto",
}
