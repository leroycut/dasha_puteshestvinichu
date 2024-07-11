package main

import (
	"context"
	"dasha_puteshestvinichu/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println("Recived %s\n", req.Name)

	return &proto.HelloReply{
		Message: "zdarova " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1354))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
