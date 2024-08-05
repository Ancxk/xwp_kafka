package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	logic "xwp_kafka/grpc_/hello_/logic"
)

type serv struct {
	logic.UnimplementedUserServerServer
}

func (*serv) GetUserInfo(context.Context, *logic.GetUserInfoReq) (*logic.GetUserInfoRes, error) {
	res := &logic.GetUserInfoRes{
		Uid:  111,
		Name: "hello!",
	}
	return res, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	logic.RegisterUserServerServer(s, &serv{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
