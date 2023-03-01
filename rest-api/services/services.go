package services

import (
	"log"
	"os"
	svcspb "restAPI/svcspb"

	"google.golang.org/grpc"
)

func PostSrv() svcspb.PostServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(os.Getenv("SVC2_URL")+":50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	blogsvc := svcspb.NewPostServiceClient(cc)

	return blogsvc
}
func UserSrv() svcspb.UserServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(os.Getenv("SVC1_URL")+":50050", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	usersvc := svcspb.NewUserServiceClient(cc)

	return usersvc
}
