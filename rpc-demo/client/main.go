package main

import (
	"fmt"
	"net/rpc"
	"time"
)
// 请求结构体定义请求参数
type Req struct {
	Number1 int
	Number2 int
}
// 响应结构体定义响应参数
type Res struct {
	Num int
}

func main() {
	// 初始化请求参数
	req := Req{Number1: 1,Number2: 2}
	// 定义响应参数
	var res Res
	// 客户端与服务端建立连接
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	// 客户端调用服务端的方法
	call := client.Go("Server.AddNum",req,&res,nil)
	// 客户端接收服务端的响应
	for  {
		select {
		case <-call.Done:
			fmt.Println("我收到了")
			fmt.Println(res)
			return
		default:
			time.Sleep(1*time.Second)
			fmt.Println("我等着")
		}
	}
}