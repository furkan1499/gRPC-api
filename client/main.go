package main

import (
	"context"
	"log"
	"time"

	"grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := proto.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Create new user
	log.Println("=== Creating New User ===")
	createResp, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Result: %s (ID: %d)", createResp.Message, createResp.Id)

	// Create second user
	createResp2, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		Name:  "Jane Smith",
		Email: "jane@example.com",
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Result: %s (ID: %d)", createResp2.Message, createResp2.Id)

	// Retrieve users
	log.Println("\n=== Retrieving Users ===")
	for i := int32(1); i <= 2; i++ {
		getUserResp, err := client.GetUser(ctx, &proto.GetUserRequest{Id: i})
		if err != nil {
			log.Printf("Failed to get user (ID: %d): %v", i, err)
			continue
		}
		log.Printf("User %d: %s (%s)", getUserResp.Id, getUserResp.Name, getUserResp.Email)
	}

	// Test non-existent user
	log.Println("\n=== Non-existent User Test ===")
	_, err = client.GetUser(ctx, &proto.GetUserRequest{Id: 999})
	if err != nil {
		log.Printf("Expected error: %v", err)
	}
}