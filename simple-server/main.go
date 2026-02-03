package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
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

type UserServer struct {
	users  map[int]*User
	mutex  sync.RWMutex
	nextID int
}

func NewUserServer() *UserServer {
	return &UserServer{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (s *UserServer) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	s.mutex.RLock()
	user, exists := s.users[id]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, fmt.Sprintf("User not found: %d", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	log.Printf("User retrieved: ID=%d, Name=%s", user.ID, user.Name)
}

func (s *UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.mutex.Lock()
	user := &User{
		ID:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[s.nextID] = user
	currentID := s.nextID
	s.nextID++
	s.mutex.Unlock()

	response := CreateUserResponse{
		ID:      currentID,
		Message: fmt.Sprintf("User successfully created: %s", req.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("New user created: ID=%d, Name=%s, Email=%s", currentID, req.Name, req.Email)
}

func (s *UserServer) ListUsers(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	s.mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	log.Printf("Listed %d users", len(users))
}

func main() {
	server := NewUserServer()
	
	r := mux.NewRouter()
	r.HandleFunc("/users", server.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	r.HandleFunc("/users", server.ListUsers).Methods("GET")

	log.Println("Starting HTTP Server on port 8080...")
	log.Println("Endpoints:")
	log.Println("  POST /users - Create user")
	log.Println("  GET /users/{id} - Get user by ID")
	log.Println("  GET /users - List all users")
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}