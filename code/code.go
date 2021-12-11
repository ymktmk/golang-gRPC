package code

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
    "fmt"
    "golang.org/x/net/context"
)

type Server struct {
	Code string `json:code`
	Result string `json:result`
}

func (s *Server) GetResult(ctx context.Context, in *Request) (*Response, error) {
	// クライアントからリクエストを受け取る
	fmt.Println(in.Code)
	// 内部処理---------------------------------------
	file_name := writeFile(in.Code)
	
	result := dockerRun(file_name)

	exec.Command(
		"rm","/go/src/work/" + file_name,
	).Run()
	// ---------------------------------------
	// ここでサーバーがクライアントにレスポンスを返す
    return &Response{Result: result}, nil
}

// dockerで実行して結果を返す
func dockerRun(file_name string) string {

	// コンテナ間マウント
	cmd := exec.Command(
		"docker","run","-i","--rm",
		"--volumes-from","grpc",
		"-w","/go/src/work",
		"python:latest",
		"python", file_name,
	)

	stdin, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(stdin, "hoge foo bar")
	stdin.Close()
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}

// ファイルにコードを書き込む
func writeFile(code string) string {

	t := time.Now()
	file_name := t.Format(time.RFC3339) + ".py"

	// コンテナ側に作ってホストと共有
	file_path := filepath.Join("go/src/work", file_name)

	f, err := os.Create(file_path)

	if err != nil {
		fmt.Println(err)
	}

	data := []byte(code)
	
	f.Write(data)

	return file_name
}