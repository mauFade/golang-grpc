package main

import (
	"context"
	"grpc-golang/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, error := grpc.Dial("localhost:9099", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if error != nil {
		log.Fatal(error)
	}

	client := pb.NewHelloClient(connection)

	request := &pb.HelloRequest{
		Name: "John Doe",
	}

	response, err := client.SayHello(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}