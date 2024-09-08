package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"unary-rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type User struct {
	Id   string
	Name string
	Age  int32
}

func Run() {
	listen, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, NewUserService())
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		panic(err)
	}

}

type UserService struct {
	pb.UnimplementedUserServiceServer

	users map[string]*User
	mu    sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

func (us *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := &User{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}

	us.users[req.Id] = user

	return &pb.AddUserResponse{
		Id:   user.Id,
		Age:  user.Age,
		Name: user.Name,
	}, nil
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{
		Id:   user.Id,
		Age:  user.Age,
		Name: user.Name,
	}, nil
}
