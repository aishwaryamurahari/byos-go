## 🚀 GoServe - Build Your Own Server (BYOS)

GoServe is a **from-scratch HTTP server implementation** in Go that demonstrates low-level network programming without any web frameworks. This educational project provides deep insights into how web servers work under the hood, covering everything from raw TCP socket handling to HTTP protocol implementation.

Whether you're learning how web servers really work, preparing for systems programming interviews, or want to understand what happens behind frameworks like Express.js or Gin, GoServe is your comprehensive learning toolkit.

## 🎯 What You'll Learn

### **Core Concepts**
- **Raw TCP Socket Programming**: Direct manipulation of network connections
- **HTTP Protocol Deep Dive**: Manual parsing and response generation
- **Concurrency Patterns**: Goroutines for handling multiple simultaneous connections
- **Network Security**: Path traversal prevention and input validation
- **Memory Management**: Efficient buffer handling and resource cleanup

### **Skills Developed**
- Understanding the foundation beneath web frameworks
- Low-level networking and protocol implementation
- Concurrent programming with Go's goroutines
- HTTP request/response lifecycle management
- File system operations and MIME type handling

## 🏗️ Architecture Overview

```
Client Request → TCP Socket → HTTP Parser → Router → Handler → Response
     ↓              ↓           ↓          ↓        ↓         ↓
  Browser        net.Conn    Custom      Route    Handler   Manual
   cURL         Accept()     Parser     Matcher   Function  Response
```

## 🔧 Core Features

### **1. Raw TCP Socket Server**
- Direct `net.Listen()` and `net.Accept()` usage
- Manual connection lifecycle management
- Graceful shutdown with signal handling
- No dependency on `net/http` package

### **2. Custom HTTP Request Parser**
- Manual parsing of HTTP request lines
- Header extraction and processing
- Body content handling with Content-Length
- Support for GET, POST, and other HTTP methods

### **3. Flexible Routing System**
- **Static Routes**: Direct path matching (`/api/hello`)
- **Dynamic Routes**: Parameter extraction (`/api/user/:id`)
- **Static File Serving**: Automatic MIME type detection

### **4. Security Features**
- Directory traversal prevention (`../` attacks)
- URL decoding for special characters
- Safe file path handling
- Input validation and sanitization

### **5. Concurrent Request Handling**
- Each connection handled in separate goroutine
- Non-blocking server architecture
- Context-based graceful shutdown
- Memory-efficient buffer management

## 📁 Project Structure

```
byos-go/
├── main.go              # Entry point & TCP server logic
├── routes/
│   └── routes.go        # HTTP routing & handler functions
├── server/
│   └── handler.go       # Static file serving logic
├── utils/
│   └── parser.go        # HTTP request parsing utilities
├── static/
│   └── index.html       # Static web assets
├── go.mod               # Go module definition
└── README.md            # This documentation
```

## 🚀 Quick Start

### **Prerequisites**
- Go 1.19+ installed
- Basic understanding of HTTP protocol
- Terminal/command line access

### **Running the Server**
```bash
# Clone or navigate to project directory
cd byos-go

# Start the server
go run main.go

# Server starts on http://localhost:8080
```

### **Testing Endpoints**
```bash
# Home page
curl http://localhost:8080/

# API endpoint
curl http://localhost:8080/api/hello

# Dynamic route with parameters
curl http://localhost:8080/api/user/john

# POST request with body
curl -X POST http://localhost:8080/api/echo \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello World"}'

# Static file serving
curl http://localhost:8080/static/index.html
```

## 🔍 Implementation Deep Dive

### **1. TCP Server (`main.go`)**
```go
// Raw TCP socket creation
listener, err := net.Listen("tcp", ":8080")

// Manual connection acceptance loop
for {
    conn, err := listener.Accept()
    go handleConnection(conn)  // Concurrent handling
}
```

**Key Learning Points:**
- How web servers accept network connections
- Goroutine-based concurrency model
- Resource management and cleanup
- Signal handling for graceful shutdown

### **2. HTTP Parser (`utils/parser.go`)**
```go
// Manual HTTP request line parsing
parts := strings.Split(line, " ")  // ["GET", "/path", "HTTP/1.1"]

// Header extraction
headerParts := strings.SplitN(line, ":", 2)
```

**Key Learning Points:**
- HTTP protocol structure and formatting
- Header parsing and validation
- Content-Length handling for request bodies
- Error handling for malformed requests

### **3. Routing Engine (`routes/routes.go`)**
```go
// Static route matching
key := method + " " + path
handler, exists := RouteTable[key]

// Dynamic parameter extraction
if strings.HasPrefix(patternParts[i], ":") {
    params[key] = pathParts[i]
}
```

**Key Learning Points:**
- URL pattern matching algorithms
- Parameter extraction techniques
- Handler function composition
- Route priority and matching order

### **4. Static File Server (`server/handler.go`)**
```go
// Security: prevent directory traversal
if strings.Contains(cleanPath, "..") {
    return "text/plain", "Forbidden", 403
}

// MIME type detection
contentType := mime.TypeByExtension(ext)
```

**Key Learning Points:**
- File system security considerations
- MIME type handling and content negotiation
- Error handling for file operations
- Performance considerations for file serving