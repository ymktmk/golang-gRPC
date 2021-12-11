package main

import (
	"google.golang.org/grpc"
	"fmt"
	"net"
	"example.com/code"
)

func main() {
	
	// gRPC
	fmt.Println("Hello gRPC Server !")

	// listenするportを設定します
	lis, err := net.Listen("tcp",":9000")
	
	// エラーハンドリング
	if err != nil {
        fmt.Println(err)
	}

	// gRPCサーバーの生成
	grpcServer := grpc.NewServer()
	
	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したハンドラを設定します。
	code.RegisterCodeServiceServer(grpcServer, &code.Server{})
	
	// サーバーを起動
    if err := grpcServer.Serve(lis); err != nil {
        fmt.Println("error")
	}
	
}