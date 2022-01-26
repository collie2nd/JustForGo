package main

import (
	hello_grpc2 "JustForGo/grpc_study/init_demo/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc2.UnimplementedHELLOGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc2.Req) (res *hello_grpc2.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc2.Res{Message: "我是从服务端返回的grpc内容"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello_grpc2.RegisterHELLOGRPCServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		fmt.Println(err)
	}
}
