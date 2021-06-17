package main

import (
	"fmt"
	"grpc-go"
	"log"
	"net"
)

type Server struct{}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
