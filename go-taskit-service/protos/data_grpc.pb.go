// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: data.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DataLayer_CreateUser_FullMethodName        = "/DataLayer/CreateUser"
	DataLayer_ReadUser_FullMethodName          = "/DataLayer/ReadUser"
	DataLayer_UpdateUser_FullMethodName        = "/DataLayer/UpdateUser"
	DataLayer_DeleteUser_FullMethodName        = "/DataLayer/DeleteUser"
	DataLayer_CreateTask_FullMethodName        = "/DataLayer/CreateTask"
	DataLayer_ReadTask_FullMethodName          = "/DataLayer/ReadTask"
	DataLayer_ReadMultipleTasks_FullMethodName = "/DataLayer/ReadMultipleTasks"
	DataLayer_UpdateTask_FullMethodName        = "/DataLayer/UpdateTask"
	DataLayer_DeleteTask_FullMethodName        = "/DataLayer/DeleteTask"
)

// DataLayerClient is the client API for DataLayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataLayerClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*ReadUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	ReadTask(ctx context.Context, in *ReadTaskRequest, opts ...grpc.CallOption) (*ReadTaskResponse, error)
	ReadMultipleTasks(ctx context.Context, in *ReadMultipleTasksRequest, opts ...grpc.CallOption) (*ReadMultipleTasksResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
}

type dataLayerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataLayerClient(cc grpc.ClientConnInterface) DataLayerClient {
	return &dataLayerClient{cc}
}

func (c *dataLayerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, DataLayer_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*ReadUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadUserResponse)
	err := c.cc.Invoke(ctx, DataLayer_ReadUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, DataLayer_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, DataLayer_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, DataLayer_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) ReadTask(ctx context.Context, in *ReadTaskRequest, opts ...grpc.CallOption) (*ReadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadTaskResponse)
	err := c.cc.Invoke(ctx, DataLayer_ReadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) ReadMultipleTasks(ctx context.Context, in *ReadMultipleTasksRequest, opts ...grpc.CallOption) (*ReadMultipleTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadMultipleTasksResponse)
	err := c.cc.Invoke(ctx, DataLayer_ReadMultipleTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, DataLayer_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, DataLayer_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataLayerServer is the server API for DataLayer service.
// All implementations must embed UnimplementedDataLayerServer
// for forward compatibility.
type DataLayerServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ReadUser(context.Context, *ReadUserRequest) (*ReadUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	ReadTask(context.Context, *ReadTaskRequest) (*ReadTaskResponse, error)
	ReadMultipleTasks(context.Context, *ReadMultipleTasksRequest) (*ReadMultipleTasksResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	mustEmbedUnimplementedDataLayerServer()
}

// UnimplementedDataLayerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataLayerServer struct{}

func (UnimplementedDataLayerServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDataLayerServer) ReadUser(context.Context, *ReadUserRequest) (*ReadUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
func (UnimplementedDataLayerServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedDataLayerServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedDataLayerServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedDataLayerServer) ReadTask(context.Context, *ReadTaskRequest) (*ReadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTask not implemented")
}
func (UnimplementedDataLayerServer) ReadMultipleTasks(context.Context, *ReadMultipleTasksRequest) (*ReadMultipleTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMultipleTasks not implemented")
}
func (UnimplementedDataLayerServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedDataLayerServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedDataLayerServer) mustEmbedUnimplementedDataLayerServer() {}
func (UnimplementedDataLayerServer) testEmbeddedByValue()                   {}

// UnsafeDataLayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataLayerServer will
// result in compilation errors.
type UnsafeDataLayerServer interface {
	mustEmbedUnimplementedDataLayerServer()
}

func RegisterDataLayerServer(s grpc.ServiceRegistrar, srv DataLayerServer) {
	// If the following call pancis, it indicates UnimplementedDataLayerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataLayer_ServiceDesc, srv)
}

func _DataLayer_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_ReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).ReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_ReadUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).ReadUser(ctx, req.(*ReadUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_ReadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).ReadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_ReadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).ReadTask(ctx, req.(*ReadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_ReadMultipleTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMultipleTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).ReadMultipleTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_ReadMultipleTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).ReadMultipleTasks(ctx, req.(*ReadMultipleTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataLayer_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataLayer_ServiceDesc is the grpc.ServiceDesc for DataLayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataLayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataLayer",
	HandlerType: (*DataLayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DataLayer_CreateUser_Handler,
		},
		{
			MethodName: "ReadUser",
			Handler:    _DataLayer_ReadUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _DataLayer_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _DataLayer_DeleteUser_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _DataLayer_CreateTask_Handler,
		},
		{
			MethodName: "ReadTask",
			Handler:    _DataLayer_ReadTask_Handler,
		},
		{
			MethodName: "ReadMultipleTasks",
			Handler:    _DataLayer_ReadMultipleTasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _DataLayer_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _DataLayer_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
