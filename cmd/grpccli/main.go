package main

import (
	"context"
	"log"

	"github.com/Chroq/sdv-m2-dev-A/pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	resp, err := client.GetUser(context.Background(), &pb.UserRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("fail to get user: %v", err)
	}
	log.Println(resp)

	resp2, err := client.GetUser(context.Background(), &pb.UserRequest{
		Id: "2",
	})
	if err != nil {
		log.Fatalf("fail to get user: %v", err)
	}
	log.Println(resp2)
}
