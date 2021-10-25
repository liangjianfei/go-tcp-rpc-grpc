# go
go的http、tcp、rpc、grpc代码示例

go 使用grpc

1、protobuf协议编译器是用c++编写的，根据自己的操作系统下载对应版本的protoc编译器：https://github.com/protocolbuffers/protobuf/releases，解压后将protoc.exe拷贝到GOPATH/bin目录下。
2、分别执行 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
3、执行 go get google.golang.org/grpc
4、配置.proto 文件
5、执行 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./helloworld.proto  //在.proto文件目录执行，helloworld为.proto文件名

//下面为.proto文件内容
syntax = "proto3";

option go_package ="./;hello_grpc";

package hello_grpc;

message Req {
    string message = 1;
}

message Res {
    string message = 1;
}

service HelloGRPC{
    rpc SayHi (Req) returns (Res);
}
