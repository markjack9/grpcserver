package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	trippb "server/proto/gen/go"
	trip "server/tripservice"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	//建立一个tcp监听地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	//使用grpc架构启动服务
	log.Fatal(s.Serve(lis))
}
