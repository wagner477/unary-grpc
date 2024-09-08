package client

import (
	"context"
	"fmt"
	"unary-rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	dial, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	userClient := pb.NewUserServiceClient(dial)

	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Age:  40,
		Name: "John Doe",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"User added: ",
		user,
	)

	user2, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "2",
		Age:  30,
		Name: "Jane Doe",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"User added: ",
		user2,
	)

	getUserResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "2",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"User found: ",
		getUserResponse,
	)
}
