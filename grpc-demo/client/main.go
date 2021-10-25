package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc-class/grpc-demo/po"
)

func main()  {
	// 客户端连接服务端
	conn,err := grpc.Dial("localhost:888",grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	client := hello_grpc.NewHelloGRPCClient(conn)
	// 客户端调用服务端方法并且接收响应
	req,_ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我从客户端来"})
	fmt.Println(req)

}