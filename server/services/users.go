package services

import (
	routes "server/routes"
	users "server/users"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserService struct{}

func (s *UserService) QueryGetUsers(ctx context.Context, in *users.UsersRequest) (*users.UsersResponse, error) {
	respUsers := routes.GetUsers()
	return &users.UsersResponse{Users: respUsers}, nil
}

func (s *UserService) QueryCreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.SuccessUserServiceResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	success := routes.CreateUser(in)
	if !success {
		return nil, status.Errorf(codes.NotFound, "user creation error")
	}
	return &users.SuccessUserServiceResponse{Success: success}, nil
}

func (s *UserService) QueryDeleteUser(ctx context.Context, in *users.DeleteUserRequest) (*users.SuccessUserServiceResponse, error) {
	success := routes.DeleteUser(in)
	return &users.SuccessUserServiceResponse{Success: success}, nil
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {}
