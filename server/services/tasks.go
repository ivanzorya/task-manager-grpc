package services

import (
	routes "server/routes"
	tasks "server/tasks"

	"golang.org/x/net/context"
)

type TaskService struct{}

func (s *TaskService) QueryGetTasks(ctx context.Context, in *tasks.TasksRequest) (*tasks.TasksResponse, error) {
	claims := getAuthUserClaims(ctx)
	respTasks := routes.GetTasks(claims)
	return &tasks.TasksResponse{Tasks: respTasks}, nil
}

func (s *TaskService) QueryCreateTask(ctx context.Context, in *tasks.CreateTaskRequest) (*tasks.SuccessResponse, error) {
	claims := getAuthUserClaims(ctx)
	success := routes.CreateTask(in, claims)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) QueryUpdateTask(ctx context.Context, in *tasks.UpdateTaskRequest) (*tasks.SuccessResponse, error) {
	claims := getAuthUserClaims(ctx)
	success := routes.UpdateTask(in, claims)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) QueryDeleteTask(ctx context.Context, in *tasks.DeleteTaskRequest) (*tasks.SuccessResponse, error) {
	success := routes.DeleteTask(in)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) mustEmbedUnimplementedTaskServiceServer() {}
