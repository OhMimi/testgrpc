package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/OhMimi/testgrpc/pb"

	"google.golang.org/grpc"
)

type Message struct {
	Bet      float64 `json:"bet"`
	GameType string  `json:"game_type"`
}

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := pb.NewEchoServerClient(conn)

	var testMessage = Message{
		Bet:      100,
		GameType: "243way",
	}

	data, err := json.Marshal(testMessage)
	if err != nil {
		return
	}

	r, err := c.Echo(context.Background(), &pb.EchoRequest{Message: string(data)})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}

	log.Printf("回傳結果：%s , 時間:%d", r.Message, r.Unixtime)

}
