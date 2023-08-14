package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello_server/pb"
	"net"
)

// grpc server

type Server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 方法里面是我们对外提供
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello" + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen,err: %v\n", err)
		return
	}
	// 创建grpc服务
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterGreeterServer(s, &Server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
