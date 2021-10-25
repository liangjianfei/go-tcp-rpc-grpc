package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc-class/grpc-demo/po"
	"net"
)

// 定义未注册的服务结构体
type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// 定义服务端可调用方法
func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())

	return &hello_grpc.Res{Message: "我是从服务端返回的grpc"}, nil
}

func main() {
	// 服务端监听端口并且启动
	l,err := net.Listen("tcp", ":888")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s,&server{})
	s.Serve(l)
}
