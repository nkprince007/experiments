package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/nkprince007/echo/server/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Say(context context.Context, req *proto.Request) (*proto.Response, error) {
	fmt.Println("Got a new Echo request")
	result := &proto.Response{Msg: req.Msg}
	return result, nil
}

func main() {
	fmt.Println("Starting Echo server")
	lis, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
