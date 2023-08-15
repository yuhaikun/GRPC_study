package main

import (
	"add_client/pb"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "127.0.0.1:8888", "the address to connect to")
	x    = flag.Int("x", 0, "Adder arg1")
	y    = flag.Int("y", 0, "Adder arg2")
)

func main() {
	//解析命令行参数
	flag.Parse()

	//连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdderClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Add(ctx, &pb.AddRequest{
		X: int32(*x),
		Y: int32(*y),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("result: %s", resp)

}
