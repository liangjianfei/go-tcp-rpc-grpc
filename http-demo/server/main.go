package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request)  {
	switch r.Method {
	case "GET":
		w.Write([]byte("我收到了GET请求"))
	case "POST":
		b,_ := io.ReadAll(r.Body)
		fmt.Println(r.Header["Test"])
		w.Write(b)
		break
	}

}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8980",nil)
}