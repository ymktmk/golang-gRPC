// this is practice gRPC Client(Golang)

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"example.com/code"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
    http.HandleFunc("/api/v1/python", handleExec)
	http.ListenAndServe(":8000", nil)
}

type Editor struct {
    Code string `json:code`
    Result string `json:result`
}

func handleExec(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
    w.Header().Set("Content-Type", "application/json")

    // recieve Request
	body := make([]byte, r.ContentLength)
    r.Body.Read(body)
    var editor Editor
	json.Unmarshal(body, &editor)

    text := editor.Code

    // gRPC
    var conn *grpc.ClientConn

    conn, err := grpc.Dial(":9000", grpc.WithInsecure())

    if err != nil {
        fmt.Println(err)
    }
    
    defer conn.Close()

    c := code.NewCodeServiceClient(conn)

    // send request
    response, err := c.GetResult(context.Background(), &code.Request{Code: text})
    
    if err != nil {
        fmt.Println(err)
    }

    editor.Result = response.Result
    json.NewEncoder(w).Encode(editor)
}
