package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

    "github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
	tasks "server/tasks"
	routes "server/routes"
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
	flag.Parse()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	grpcServer := grpc.NewServer()
	tasks.RegisterTaskServiceServer(grpcServer, &TaskService{})
	grpclog.SetLogger(log.New(os.Stdout, "Server: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: cors(wrappedServer),
	}

	grpclog.Printf("Starting server. http port: %s", port)

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}

type TaskService struct{}


func (s *TaskService) QueryGetTasks(ctx context.Context, in *tasks.TasksRequest) (*tasks.TasksResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
	respTasks := routes.GetTasks()
	return &tasks.TasksResponse{Tasks: respTasks}, nil
}

func (s *TaskService) QueryCreateTask(ctx context.Context, in *tasks.CreateTaskRequest) (*tasks.SuccessResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
	success := routes.CreateTask(in)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) QueryUpdateTask(ctx context.Context, in *tasks.UpdateTaskRequest) (*tasks.SuccessResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
	success := routes.UpdateTask(in)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) QueryDeleteTask(ctx context.Context, in *tasks.DeleteTaskRequest) (*tasks.SuccessResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))
	success := routes.DeleteTask(in)
	return &tasks.SuccessResponse{Success: success}, nil
}

func (s *TaskService) mustEmbedUnimplementedTaskServiceServer() {}
