package internal

import (
	"context"

	pb "gn222gq.2dv013.a2/protos"
)

// Represents the user service with functionality tied to user actions
type UserService interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	ReadUser(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}

// Represents the task service with functionality tied to tasks
type TaskService interface {
	CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error)
	ReadTask(ctx context.Context, req *pb.ReadTaskRequest) (*pb.ReadTaskResponse, error)
	ReadMultipleTasks(ctx context.Context, req *pb.ReadMultipleTasksRequest) (*pb.ReadMultipleTasksResponse, error)
	UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error)
}

// Represents the data layer server
type DataLayerServer struct {
	pb.UnimplementedDataLayerServer
	UserService UserService
	TaskService TaskService
}

// Constructor function
func NewDataLayerServer(userService UserService, taskService TaskService) *DataLayerServer {
	return &DataLayerServer{UserService: userService, TaskService: taskService}
}

// Function for creating a user
func (d *DataLayerServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return d.UserService.CreateUser(ctx, req)
}

// Function for getting a user
func (d *DataLayerServer) ReadUser(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	return d.UserService.ReadUser(ctx, req)
}

// Function for updating a user
func (d *DataLayerServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return d.UserService.UpdateUser(ctx, req)
}

// Function for deleting a user
func (d *DataLayerServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return d.UserService.DeleteUser(ctx, req)
}

// Function for creating a task
func (d *DataLayerServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return d.TaskService.CreateTask(ctx, req)
}

// Function for getting a task
func (d *DataLayerServer) ReadTask(ctx context.Context, req *pb.ReadTaskRequest) (*pb.ReadTaskResponse, error) {
	return d.TaskService.ReadTask(ctx, req)
}

// Function for getting multiple tasks of a user
func (d *DataLayerServer) ReadMultipleTasks(ctx context.Context, req *pb.ReadMultipleTasksRequest) (*pb.ReadMultipleTasksResponse, error) {
	return d.TaskService.ReadMultipleTasks(ctx, req)
}

// Function for updating a task
func (d *DataLayerServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	return d.TaskService.UpdateTask(ctx, req)
}

// Function for deleting a task
func (d *DataLayerServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return d.TaskService.DeleteTask(ctx, req)
}
