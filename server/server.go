package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/OhMimi/testgrpc/pb"
	"github.com/tidwall/gjson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.Message)
	log.Printf("receive bet:%d, gametype:%s\n", gjson.Get(req.Message, "bet").Int(), gjson.Get(req.Message, "game_type").String())

	return &pb.EchoReply{
		Message:  req.Message,
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

	pb.RegisterEchoServerServer(grpc, es)

	reflection.Register(grpc)
	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}
}
