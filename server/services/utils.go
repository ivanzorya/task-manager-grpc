package services

import (
	"os"
	"time"

	models "server/models"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var secretKey = os.Getenv("SECRET")
var tokenDuration = 1440 * time.Minute

func getAuthUserClaims(ctx context.Context) *models.UserClaims {
	md, _ := metadata.FromIncomingContext(ctx)
	jwtManager := models.NewJWTManager(secretKey, tokenDuration)
	claims, _ := jwtManager.Verify(md["authorization"][0])
	return claims
}

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	authMethods := map[string]bool{
		"/main.TaskService/QueryGetTasks":   true,
		"/main.TaskService/QueryCreateTask": true,
		"/main.TaskService/QueryUpdateTask": true,
		"/main.TaskService/QueryDeleteTask": true,
		"/main.UserService/QueryGetUsers":   true,
		"/main.UserService/QueryDeleteUser": true,
	}

	if _, auth := authMethods[info.FullMethod]; auth {
		md, _ := metadata.FromIncomingContext(ctx)
		jwtManager := models.NewJWTManager(secretKey, tokenDuration)
		if len(md["authorization"]) > 0 {
			_, err := jwtManager.Verify(md["authorization"][0])
			if err != nil {
				return nil, status.Errorf(codes.PermissionDenied, "does not have permission")
			}
		}
	}

	return handler(ctx, req)
}
