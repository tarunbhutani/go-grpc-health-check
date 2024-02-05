package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/grpc_server/health_info"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnsafeHealthInfoServiceServer
}

func (s *server) CheckHealth(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("Received Check Health API")
	return &pb.HealthCheckResponse{Status: true, Message: "Healthy"}, nil
}

func (s *server) GetServiceInfo(ctx context.Context, in *pb.ServiceInfoRequest) (*pb.ServiceInfoResponse, error) {
	log.Println("Received Get Service Info")
	return &pb.ServiceInfoResponse{ServiceName: "grpc_server", Version: "1.0"}, nil
}

func main() {
	log.Println("Starting GRPC Server...")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHealthInfoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
