package main

import (
	"context"
	"flag"
	"hello_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "7k7k", "通过-name告诉server你是谁")

// grpc客户端
// 调用server端的SayHello方法
func main() {
	flag.Parse() //解析命令行参数
	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err: %v", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	c := pb.NewGreeterClient(conn)
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("c.SayHello failed,err: %v", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp: %v", resp.GetReply())
}
