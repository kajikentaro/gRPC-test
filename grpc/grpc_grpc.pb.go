// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: grpc.proto

package grpc

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

// DBWriterClient is the client API for DBWriter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBWriterClient interface {
	CreateNewUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	WriteDB(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	ReadDB(ctx context.Context, in *User, opts ...grpc.CallOption) (*DataList, error)
}

type dBWriterClient struct {
	cc grpc.ClientConnInterface
}

func NewDBWriterClient(cc grpc.ClientConnInterface) DBWriterClient {
	return &dBWriterClient{cc}
}

func (c *dBWriterClient) CreateNewUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gRPC_test.DBWriter/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBWriterClient) WriteDB(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/gRPC_test.DBWriter/WriteDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBWriterClient) ReadDB(ctx context.Context, in *User, opts ...grpc.CallOption) (*DataList, error) {
	out := new(DataList)
	err := c.cc.Invoke(ctx, "/gRPC_test.DBWriter/ReadDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBWriterServer is the server API for DBWriter service.
// All implementations must embed UnimplementedDBWriterServer
// for forward compatibility
type DBWriterServer interface {
	CreateNewUser(context.Context, *User) (*User, error)
	WriteDB(context.Context, *Data) (*Data, error)
	ReadDB(context.Context, *User) (*DataList, error)
	mustEmbedUnimplementedDBWriterServer()
}

// UnimplementedDBWriterServer must be embedded to have forward compatible implementations.
type UnimplementedDBWriterServer struct {
}

func (UnimplementedDBWriterServer) CreateNewUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedDBWriterServer) WriteDB(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteDB not implemented")
}
func (UnimplementedDBWriterServer) ReadDB(context.Context, *User) (*DataList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDB not implemented")
}
func (UnimplementedDBWriterServer) mustEmbedUnimplementedDBWriterServer() {}

// UnsafeDBWriterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBWriterServer will
// result in compilation errors.
type UnsafeDBWriterServer interface {
	mustEmbedUnimplementedDBWriterServer()
}

func RegisterDBWriterServer(s grpc.ServiceRegistrar, srv DBWriterServer) {
	s.RegisterService(&DBWriter_ServiceDesc, srv)
}

func _DBWriter_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBWriterServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC_test.DBWriter/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBWriterServer).CreateNewUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBWriter_WriteDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBWriterServer).WriteDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC_test.DBWriter/WriteDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBWriterServer).WriteDB(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBWriter_ReadDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBWriterServer).ReadDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC_test.DBWriter/ReadDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBWriterServer).ReadDB(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// DBWriter_ServiceDesc is the grpc.ServiceDesc for DBWriter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBWriter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC_test.DBWriter",
	HandlerType: (*DBWriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewUser",
			Handler:    _DBWriter_CreateNewUser_Handler,
		},
		{
			MethodName: "WriteDB",
			Handler:    _DBWriter_WriteDB_Handler,
		},
		{
			MethodName: "ReadDB",
			Handler:    _DBWriter_ReadDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
