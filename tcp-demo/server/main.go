package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务端建立TCPsocket
	tcpAddr,err := net.ResolveTCPAddr("tcp",":8088")
	if err!=nil {
		fmt.Println(err)
	}
	// 服务端监听端口
	listener,_ := net.ListenTCP("tcp",tcpAddr)
	// 服务端循环接收请求
	for  {
		conn,e := listener.AcceptTCP()
		if e!=nil {
			fmt.Println(e)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	// 服务端循环读取客户端发送的数据
	for  {
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println(conn.RemoteAddr().String()+string(buf[0:n]))
		// 给客户端发送数据
		str := "收到了"+string(buf[0:n])
		conn.Write([]byte(str))
	}
}
