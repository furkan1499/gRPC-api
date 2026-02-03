package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"grpc-example/proto"
	"google.golang.org/grpc"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	users map[int32]*proto.GetUserResponse
	mutex sync.RWMutex
	nextID int32
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: make(map[int32]*proto.GetUserResponse),
		nextID: 1,
	}
}

func (s *UserServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found: %d", req.Id)
	}
	
	log.Printf("User retrieved: ID=%d, Name=%s", user.Id, user.Name)
	return user, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user := &proto.GetUserResponse{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	
	s.users[s.nextID] = user
	currentID := s.nextID
	s.nextID++
	
	log.Printf("New user created: ID=%d, Name=%s, Email=%s", currentID, req.Name, req.Email)
	
	return &proto.CreateUserResponse{
		Id:      currentID,
		Message: fmt.Sprintf("User successfully created: %s", req.Name),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	userServer := NewUserServer()
	proto.RegisterUserServiceServer(grpcServer, userServer)
	
	log.Println("Starting gRPC Server... Port: 50051")
	log.Println("You can run the client to test it")
	
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}