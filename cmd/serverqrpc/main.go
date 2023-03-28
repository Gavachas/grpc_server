package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"

	sample_grpc "github.com/Gavachas/grpc_sample/grpc_s"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort = "4041"

type server struct {
	sample_grpc.ItilServiceServer
}

func GetRegion() string {
	regions := []string{
		"Albania",
		"Russia",
		"Spain",
		"China",
	}
	return regions[rand.Intn(len(regions))]
}
func (*server) GetUserRegion(ctx context.Context, req *sample_grpc.GetUserRequest) (*sample_grpc.GetUserRegionResponse, error) {
	fmt.Println("Get user region")
	return &sample_grpc.GetUserRegionResponse{Name: GetRegion()}, nil
}
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Itil grpc Service Started")

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	sample_grpc.RegisterItilServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
