# Go Server Examples

This project contains both gRPC and HTTP REST API server examples built with Go. It includes basic CRUD operations for user management.

## Project Structure

```
├── go.mod                    # Go module definition
├── proto/                    # Protocol Buffer definitions (gRPC)
│   ├── user.proto           # Proto definitions
│   ├── user.pb.go           # Generated Go structs
│   └── user_grpc.pb.go      # Generated gRPC code
├── server/                   # gRPC Server
│   └── main.go              # gRPC Server implementation
├── client/                   # gRPC Client
│   └── main.go              # gRPC Test Client
├── simple-server/            # HTTP REST API Server
│   └── main.go              # HTTP Server implementation
├── simple-client/            # HTTP REST API Client
│   └── main.go              # HTTP Test Client
├── test-commands.sh          # Automated test script
└── README.md
```

## Requirements

- Go 1.21+
- Protocol Buffers compiler (protoc) - optional for gRPC

## Installation

```bash
# Download dependencies
go mod tidy
```

## Running the HTTP REST API Server (Recommended)

### 1. Start the HTTP Server
```bash
go run simple-server/main.go
```

The server will start on port 8080 with these endpoints:
- `POST /users` - Create user
- `GET /users/{id}` - Get user by ID  
- `GET /users` - List all users

### 2. Test the HTTP Client
In another terminal:
```bash
go run simple-client/main.go
```

## How to View and Test the Server

### 1. Web Browser
Once the server is running, open your browser and visit:
- **List all users:** http://localhost:8080/users
- **Get specific user:** http://localhost:8080/users/1

### 2. Command Line with curl

#### Manual Testing:
```bash
# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com"}'

# Get a user by ID
curl http://localhost:8080/users/1

# List all users
curl http://localhost:8080/users

# Test non-existent user (should return error)
curl http://localhost:8080/users/999
```

#### Automated Testing:
Run the included test script:
```bash
chmod +x test-commands.sh
./test-commands.sh
```

### 3. Go Client
```bash
go run simple-client/main.go
```

### Expected Responses:

**Create User:**
```json
{"id":1,"message":"User successfully created: Alice Johnson"}
```

**Get User:**
```json
{"id":1,"name":"Alice Johnson","email":"alice@example.com"}
```

**List Users:**
```json
[{"id":1,"name":"Alice Johnson","email":"alice@example.com"}]
```

**Non-existent User:**
```
User not found: 999
```

## Running the gRPC Server (Advanced)

### 1. Start the gRPC Server
```bash
go run server/main.go
```

The server will start on port 50051.

### 2. Test the gRPC Client
In another terminal:
```bash
go run client/main.go
```

## Features

- **GetUser**: Retrieve user by ID
- **CreateUser**: Create new user
- **ListUsers**: List all users (HTTP only)
- Thread-safe in-memory storage
- Simple error handling
- Context timeout support (gRPC)

## Sample Output (HTTP)

```
=== Creating New Users ===
Result: User successfully created: John Doe (ID: 1)
Result: User successfully created: Jane Smith (ID: 2)

=== Retrieving Users ===
User 1: John Doe (john@example.com)
User 2: Jane Smith (jane@example.com)

=== Listing All Users ===
User 1: John Doe (john@example.com)
User 2: Jane Smith (jane@example.com)

=== Non-existent User Test ===
Expected error: user not found: User not found: 999
```

## Performance Comparison

**HTTP REST API:**
- Simple JSON over HTTP/1.1
- Easy to debug and test
- Wide compatibility
- Human-readable

**gRPC:**
- Binary protocol over HTTP/2
- Type-safe with Protocol Buffers
- Better performance for high-throughput
- Streaming support
- Multi-language client generation

## Advanced Features

You can extend these examples with:

- Database integration (PostgreSQL, MongoDB)
- Authentication/Authorization (JWT)
- Middleware (logging, metrics, tracing)
- Docker containerization
- Kubernetes deployment
- API documentation (OpenAPI/Swagger for REST)