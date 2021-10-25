package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 客户端与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	// 循环读数据并且发送给服务端
	for  {
		// 给服务端发送数据
		bytes,_,_ := reader.ReadLine()
		conn.Write(bytes)
		// 从服务端读取数据
		rb := make([]byte,1024)
		rn,_ := conn.Read(rb)
		fmt.Println(string(rb[0:rn]))
	}
}
