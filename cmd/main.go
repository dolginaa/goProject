package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/dolginaa/goProject/internal/database"
	"github.com/dolginaa/goProject/internal/database/repository"
	"github.com/dolginaa/goProject/internal/services"
	pb "github.com/dolginaa/goProject/pkg/schedule"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		err := pb.RegisterScheduleHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
		if err != nil {
			panic(err)
		}

		log.Printf("server listening at 8081")
		if err := http.ListenAndServe(":8081", mux); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool, err := database.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	activityRepo := repository.NewActivities(pool)

	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScheduleServer(s, &services.Server{Repo: activityRepo})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
