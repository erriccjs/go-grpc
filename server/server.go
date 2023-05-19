// This code creates a gRPC server, registers our `Greetings` service
// implementation, and starts listening for incoming connections on
// port `50051`.

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	greetings "github.com/erriccjs/go-grpc/proto"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetings.GreetRequest) (*greetings.GreetResponse, error) {
	return &greetings.GreetResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greetings.RegisterGreetingsServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
