package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建客户端
	client := new(http.Client)
	// 构建请求
	req,_ := http.NewRequest("POST","http://127.0.0.1:8980",bytes.NewBuffer([]byte("{\"test\":\"我是客户端\"}")))
	req.Header["test"] = []string{"test1","test2"}
	// 发送请求
	res,_ := client.Do(req)
	// 接收响应
	body := res.Body
	// 读取并且打印响应数据
	b,_ := io.ReadAll(body)
	fmt.Println(string(b))
}
