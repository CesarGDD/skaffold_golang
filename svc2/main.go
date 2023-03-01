package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	svcspb "svc2/svcspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	svcspb.PostServiceServer
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Users service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	svcspb.RegisterPostServiceServer(s, &server{})

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

func (*server) CreatePost(ctx context.Context, req *svcspb.CreatePostRequest) (*svcspb.CreatePostResponse, error) {
	fmt.Println("Create Post gRPC")
	return &svcspb.CreatePostResponse{
		Post: &svcspb.Post{
			Title:   req.Post.Title,
			Content: req.Post.Content,
		},
	}, nil
}
