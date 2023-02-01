package main

import (
	"context"
	"grpc-golang/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (server *Server) sayHello (context context.Context, inbound *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Name: "Fala aí " + inbound.GetName()}, nil
}

func main() {
	println("gRPC server running")

	listener, err := net.Listen("tcp", "localhost:9099")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}