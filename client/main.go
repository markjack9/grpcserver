package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	trippb "server/proto/gen/go"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	//grpc建立一个客户端连接，使用不安全的模式
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip234",
	})
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	fmt.Println(r)
}
