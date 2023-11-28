package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	auth "server/auth"
	models "server/models"
	services "server/services"
	tasks "server/tasks"
	users "server/users"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	secretKey := os.Getenv("SECRET")
	tokenDuration := 1440 * time.Minute

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(services.UnaryInterceptor),
	)
	jwtManager := models.NewJWTManager(secretKey, tokenDuration)
	authServer := services.NewAuthServer(jwtManager)
	auth.RegisterAuthServiceServer(grpcServer, authServer)
	tasks.RegisterTaskServiceServer(grpcServer, &services.TaskService{})
	users.RegisterUserServiceServer(grpcServer, &services.UserService{})
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))

	wrappedServer := grpcweb.WrapServer(grpcServer)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: cors(wrappedServer),
	}

	grpclog.Infof("Starting server. http port: %s", port)

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
