package main

import (
	pb "../pb"
	"context"
	"log"
	"time"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.Msg)
	return &pb.EchoReply{
		Msg:      req.Msg,
		Unixtime: time.Now().Unix(),
	}, nil

}

func main() {
	apiListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return
	}

	// 註冊 grpc
	es := &EchoServer{}

	grpc := grpc.NewServer()

	pb.RegisterEchoServer(grpc, es)

	reflection.Register(grpc)
	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}
}
