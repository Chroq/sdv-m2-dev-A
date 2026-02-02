package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Chroq/sdv-m2-dev-A/pb/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if req.Id != "1" {
		return &pb.UserResponse{}, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.UserResponse{
		Name: "John Doe",
		Age:  30,
	}, nil
}
