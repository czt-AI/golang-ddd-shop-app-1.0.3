package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"yourdomain/your_service" // Replace with your actual package path
)

type server struct {
	YourServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	yourservice.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
