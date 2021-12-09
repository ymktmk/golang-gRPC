// this is practice gRPC Client(Golang)

package main

import (
    "fmt"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
	"example.com/code"
)

func main() {

    var conn *grpc.ClientConn

    conn, err := grpc.Dial(":9000", grpc.WithInsecure())

    if err != nil {
        fmt.Println("error")
    }
    
    defer conn.Close()

    c := code.NewCodeServiceClient(conn)

    // リクエストを送ってレスポンスが帰ってくる
    response, err := c.GetResult(context.Background(), &code.Request{Code: "print(1)"})
    
    if err != nil {
        fmt.Println("error")
    }

    // サーバーからのレスポンスを表示
    fmt.Println(response.Result)
}