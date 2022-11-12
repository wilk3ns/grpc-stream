package main

import (
	"context"
	pb "grpcTest/gen/proto"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(req *pb.UserRequest, stream pb.TestApi_GetUserServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream cancelled")
		default:
			time.Sleep(1 * time.Second)
			name := "Kamran"
			age := 0 + rand.Int31n(40)
			email := "kamran@gmail.com"

			stream.SendMsg(&pb.UserResponse{Name: name, Age: age, Email: email})
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
