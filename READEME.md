# Go Week 1 API 🚀

A simple REST API built with Go to practice backend engineering fundamentals such as HTTP handlers, project structure, and testing.
This project follows a clean, production-style layout rather than a beginner single-file setup.

---

## 📌 Features

✅ Health check endpoint  
✅ Current time endpoint  
✅ Clean Go project structure  
✅ HTTP handler testing  
✅ Lightweight and fast using Go’s standard library  

---

## 🏗️ Project Structure

go-week1-api
│
├── cmd/api # Application entry point
├── internal
│ ├── handlers # HTTP handlers
│ └── service # Business logic (future expansion)
│
├── go.mod
└── README.md

This structure mirrors how production Go services are commonly organized and helps keep the codebase maintainable and scalable.

---

## 🚀 How To Run

### 1️⃣ Clone the repository
git clone https://github.com/YOUR_USERNAME/go-week1-api.git


### 2️⃣ Navigate into the project
cd go-week1-api


### 3️⃣ Run the server
go run cmd/api/main.go


Server starts at:
http://localhost:8080

---

## 📡 API Endpoints

### ✅ Health Check
GET /health

**Response:**
ok

---

### ✅ Current Time
GET /time

**Response:**
```json
{
  "time": "RFC3339 timestamp",
  "status": "running"
}

Run all tests with:
go test ./...

🛠️ Tech Stack

Go
net/http
httptest

No external frameworks were used — the goal is to build a strong foundation in core Go before introducing additional tooling.

🎯 Learning Goals

This project was created to:

Understand Go syntax and idioms
Build REST endpoints
Practice handler-based architecture
Write fast, reliable HTTP tests
Follow production-style project organization

🔥 Next Improvements
Planned enhancements include:
Adding a service layer
Implementing logging middleware
Dockerizing the application
Introducing configuration management
Expanding test coverage
Adding graceful shutdown