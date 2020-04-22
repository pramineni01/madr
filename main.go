package main

import (
	"log"
	"net"

	mock "github.com/pramineni01/madr/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error opening port to listen. Err: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterLogServiceServer(grpcServer, &mock.Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error starting the server. Err: %v", err)
	}
}
