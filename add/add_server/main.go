package main

import (
	"add_server/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Res: in.X + in.Y}, nil
}

func main() {
	// 监听本地8888端口
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
		return
	}
	// 创建grpc服务器
	s := grpc.NewServer()
	// 在grpc服务端注册服务
	pb.RegisterAdderServer(s, &server{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}

}
