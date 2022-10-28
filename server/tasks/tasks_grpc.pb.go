// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tasks.proto

package __

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	QueryGetTasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error)
	QueryCreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	QueryUpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	QueryDeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) QueryGetTasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error) {
	out := new(TasksResponse)
	err := c.cc.Invoke(ctx, "/main.TaskService/QueryGetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) QueryCreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/main.TaskService/QueryCreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) QueryUpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/main.TaskService/QueryUpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) QueryDeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/main.TaskService/QueryDeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations should embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	QueryGetTasks(context.Context, *TasksRequest) (*TasksResponse, error)
	QueryCreateTask(context.Context, *CreateTaskRequest) (*SuccessResponse, error)
	QueryUpdateTask(context.Context, *UpdateTaskRequest) (*SuccessResponse, error)
	QueryDeleteTask(context.Context, *DeleteTaskRequest) (*SuccessResponse, error)
}

// UnimplementedTaskServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) QueryGetTasks(context.Context, *TasksRequest) (*TasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGetTasks not implemented")
}
func (UnimplementedTaskServiceServer) QueryCreateTask(context.Context, *CreateTaskRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCreateTask not implemented")
}
func (UnimplementedTaskServiceServer) QueryUpdateTask(context.Context, *UpdateTaskRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUpdateTask not implemented")
}
func (UnimplementedTaskServiceServer) QueryDeleteTask(context.Context, *DeleteTaskRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeleteTask not implemented")
}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_QueryGetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).QueryGetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TaskService/QueryGetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).QueryGetTasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_QueryCreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).QueryCreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TaskService/QueryCreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).QueryCreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_QueryUpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).QueryUpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TaskService/QueryUpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).QueryUpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_QueryDeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).QueryDeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TaskService/QueryDeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).QueryDeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryGetTasks",
			Handler:    _TaskService_QueryGetTasks_Handler,
		},
		{
			MethodName: "QueryCreateTask",
			Handler:    _TaskService_QueryCreateTask_Handler,
		},
		{
			MethodName: "QueryUpdateTask",
			Handler:    _TaskService_QueryUpdateTask_Handler,
		},
		{
			MethodName: "QueryDeleteTask",
			Handler:    _TaskService_QueryDeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks.proto",
}
