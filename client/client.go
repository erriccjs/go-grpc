// This code connects to the gRPC server, creates a `Greetings` client,
// and calls the `Greet` method with a `GreetRequest` containing
// the name "Eric". The response is then logged to the console.

package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/erriccjs/go-grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Greet(ctx, &pb.GreetRequest{Name: "Eric"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", response.Message)
}
