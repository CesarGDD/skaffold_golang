package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	svcspb "svc1/svcspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	svcspb.UserServiceServer
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Users service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	svcspb.RegisterUserServiceServer(s, &server{})

	//Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	//Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//BLock until a signal is received
	<-ch

	fmt.Println("Closing Conection Connection")
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of program")
}

func (*server) CreateUser(ctx context.Context, req *svcspb.CreateUserRequest) (*svcspb.CreateUserResponse, error) {
	fmt.Println("Create User gRPC")
	return &svcspb.CreateUserResponse{
		User: &svcspb.User{
			Username: req.Username,
		},
	}, nil
}
