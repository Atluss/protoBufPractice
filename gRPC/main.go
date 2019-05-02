package gRPC

import (
	"context"
	greeting "github.com/Atluss/protoBufPractice/proto/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// generate proto file
//go:generate protoc --go_out=plugins=grpc:proto proto/greeting.proto

// server is used to implement greeting.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *greeting.HelloRequest) (*greeting.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &greeting.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greeting.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}