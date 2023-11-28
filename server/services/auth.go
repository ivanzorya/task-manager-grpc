package services

import (
	auth "server/auth"
	models "server/models"
	routes "server/routes"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthServer struct {
	jwtManager *models.JWTManager
}

func NewAuthServer(jwtManager *models.JWTManager) *AuthServer {
	return &AuthServer{jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := routes.GetUser(in)

	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &auth.LoginResponse{AccessToken: token}
	return res, nil
}

func (s *AuthServer) mustEmbedUnimplementedAuthServiceServer() {}
