package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcTest/gen/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "hello everyone!"})
	if err != nil {
		log.Fatalln(err)
	}

	stream, err := client.GetUser(context.Background(), &pb.UserRequest{Uuid: "Kamranno"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Msg)

	go func() {
		for {
			value, err := stream.Recv()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(value.Age)
		}
	}()
	fmt.Scanln()
}
