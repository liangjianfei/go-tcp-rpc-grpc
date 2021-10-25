package main

import (
	"fmt"
	"net"
	"net/http"
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

// 定义服务
type Server struct {

}

// 定义服务可以调用的方法
func (s *Server) AddNum(req Req,res *Res) error{
	time.Sleep(5*time.Second)
	res.Num = req.Number1 + req.Number2
	return nil
}

func main() {
	// 注册服务
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	// 服务端监听端口
	l,e := net.Listen("tcp",":8080")
	if e!=nil {
		fmt.Println(e)
	}
	// 服务端启动服务
	http.Serve(l,nil)
}