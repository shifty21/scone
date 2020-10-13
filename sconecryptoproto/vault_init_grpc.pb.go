// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sconecryptoproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	Encrypt(ctx context.Context, in *BytePacket, opts ...grpc.CallOption) (*BytePacket, error)
	Decrypt(ctx context.Context, in *BytePacket, opts ...grpc.CallOption) (*BytePacket, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) Encrypt(ctx context.Context, in *BytePacket, opts ...grpc.CallOption) (*BytePacket, error) {
	out := new(BytePacket)
	err := c.cc.Invoke(ctx, "/sconecryptogrpc.CryptoService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) Decrypt(ctx context.Context, in *BytePacket, opts ...grpc.CallOption) (*BytePacket, error) {
	out := new(BytePacket)
	err := c.cc.Invoke(ctx, "/sconecryptogrpc.CryptoService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	Encrypt(context.Context, *BytePacket) (*BytePacket, error)
	Decrypt(context.Context, *BytePacket) (*BytePacket, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) Encrypt(context.Context, *BytePacket) (*BytePacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedCryptoServiceServer) Decrypt(context.Context, *BytePacket) (*BytePacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytePacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sconecryptogrpc.CryptoService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).Encrypt(ctx, req.(*BytePacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytePacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sconecryptogrpc.CryptoService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).Decrypt(ctx, req.(*BytePacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sconecryptogrpc.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _CryptoService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _CryptoService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vault_init.proto",
}