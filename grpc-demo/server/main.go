package main

import (
	"context"
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("server received: ", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("server listening at ", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterCreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
