# Go HTTP Server - Day 1 Internship Task

## Overview

This project demonstrates the fundamentals of Go's `net/http` package and the HTTP request lifecycle. It implements a basic HTTP server with multiple endpoints, request logging middleware, and proper HTTP response handling.

The project is designed as a Day 1 backend internship task to understand how web servers work in Go.

---

## Features

- Basic HTTP Server using Go
- Multiple Route Handlers
- Request Logging Middleware
- Health Check Endpoint
- About Endpoint
- HTTP Status Code Handling
- Request Lifecycle Monitoring

---

## Tech Stack

- Go (Golang)
- net/http Package
- Standard Library Middleware

---

## Project Structure

```bash
day1-http-server/
│
├── main.go
├── go.mod
└── README.md
```

---

## HTTP Request Lifecycle

1. Client sends an HTTP request.
2. Server receives the request.
3. Middleware logs request details.
4. Request is routed to the appropriate handler.
5. Handler processes the request.
6. Response is sent back to the client.
7. Middleware logs completion time.

---

## Available Endpoints

### Home Endpoint

```http
GET /
```

Response:

```text
Welcome to Internship HTTP Server
```

---

### Health Check Endpoint

```http
GET /health
```

Response:

```text
Server is Healthy
```

Status Code:

```http
200 OK
```

---

### About Endpoint

```http
GET /about
```

Response:

```text
This server is built using Go net/http package
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/your-username/day1-http-server.git
```

### Navigate to Project

```bash
cd day1-http-server
```

### Initialize Go Modules

```bash
go mod init day1-http-server
```

### Run the Server

```bash
go run main.go
```

---

## Server Output

```bash
Server running at http://localhost:8080
```

---

## Testing with Browser

Open:

```text
http://localhost:8080
```

```text
http://localhost:8080/health
```

```text
http://localhost:8080/about
```

---

## Testing with Curl

### Home Route

```bash
curl http://localhost:8080/
```

### Health Route

```bash
curl http://localhost:8080/health
```

### About Route

```bash
curl http://localhost:8080/about
```

---

## Sample Logs

```bash
2025/06/24 10:30:11 Method=GET Path=/ RemoteAddr=[::1]:54321
2025/06/24 10:30:11 Completed in 240.1µs

2025/06/24 10:31:05 Method=GET Path=/health RemoteAddr=[::1]:54322
2025/06/24 10:31:05 Completed in 180.4µs
```

---

## Concepts Learned

- Go HTTP Server Fundamentals
- net/http Package
- HTTP Request Lifecycle
- Route Handling
- Middleware Implementation
- Logging and Monitoring
- Status Codes
- Server Initialization

---

## Future Improvements

- REST API Development
- JSON Responses
- CRUD Operations
- Database Integration
- Authentication & Authorization
- Docker Deployment
- Unit Testing

---
