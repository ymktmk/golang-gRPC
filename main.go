package main

import (
	"google.golang.org/grpc"
	"fmt"
	"net"
	"example.com/code"
)

func main() {
	
	// gRPC
	fmt.Println("Hello gRPC API Server !")

	// listenするportを設定します
	lis, err := net.Listen("tcp",":9000")
	
	// エラーハンドリング
	if err != nil {
        fmt.Println("error")
	}

	s := code.Server{}
	
	grpcServer := grpc.NewServer()
	
	code.RegisterCodeServiceServer(grpcServer, &s)
	
	// エラーハンドリング
    if err := grpcServer.Serve(lis); err != nil {
        fmt.Println("error")
    }
}