package main

import (
	hello_grpc2 "JustForGo/grpc_study/init_demo/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	client := hello_grpc2.NewHELLOGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc2.Req{Message: "我从客户端来"})
	fmt.Println(req.GetMessage())
}
