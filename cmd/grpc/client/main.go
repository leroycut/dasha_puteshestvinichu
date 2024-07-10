package main

import (
	"context"
	"dasha_puteshestvinichu/proto"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	name := flag.String("name", "pockemon", "user")

	flag.Parse()

	conn, err := grpc.NewClient("0.0.0.0:1354", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Greetings: %s", res.Message)
}
