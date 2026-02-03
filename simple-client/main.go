package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func createUser(name, email string) (*CreateUserResponse, error) {
	req := CreateUserRequest{Name: name, Email: email}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createResp CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

func getUser(id int) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user not found: %s", string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func listUsers() ([]*User, error) {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []*User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	log.Println("=== Creating New Users ===")
	
	// Create first user
	createResp1, err := createUser("John Doe", "john@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Result: %s (ID: %d)", createResp1.Message, createResp1.ID)

	// Create second user
	createResp2, err := createUser("Jane Smith", "jane@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("Result: %s (ID: %d)", createResp2.Message, createResp2.ID)

	log.Println("\n=== Retrieving Users ===")
	
	// Get users by ID
	for i := 1; i <= 2; i++ {
		user, err := getUser(i)
		if err != nil {
			log.Printf("Failed to get user (ID: %d): %v", i, err)
			continue
		}
		log.Printf("User %d: %s (%s)", user.ID, user.Name, user.Email)
	}

	log.Println("\n=== Listing All Users ===")
	
	// List all users
	users, err := listUsers()
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	
	for _, user := range users {
		log.Printf("User %d: %s (%s)", user.ID, user.Name, user.Email)
	}

	log.Println("\n=== Non-existent User Test ===")
	
	// Test non-existent user
	_, err = getUser(999)
	if err != nil {
		log.Printf("Expected error: %v", err)
	}
}