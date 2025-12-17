# Task API

A RESTful task management API built with Go, PostgreSQL, and JWT authentication.

## Features

- User registration and authentication
- JWT-based authorization
- Full CRUD operations for tasks
- Tasks scoped to authenticated users

## Tech Stack

- **Language:** Go 1.22
- **Router:** chi
- **Database:** PostgreSQL 16
- **Authentication:** JWT (HS256)
- **Password Hashing:** bcrypt

## Quick Start

### Prerequisites

- Go 1.22+
- Docker and Docker Compose
- migrate CLI

### Setup

### Setup

1. Clone the repository
```bash
   git clone https://github.com/yourusername/task-api.git
   cd task-api
```

2. Create environment file
```bash
   cp .env.example .env
```

3. (Optional) Edit `.env` and change `JWT_SECRET` to your own value

4. Start the database
```bash
   make up
```

5. Run migrations
```bash
   make migrate-up
```

6. Start the server
```bash
   make run
```

Server runs at `http://localhost:8080`

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /health | Health check |
| POST | /register | Create account |
| POST | /login | Get JWT token |

### Protected Endpoints (require JWT)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /tasks | List all tasks |
| POST | /tasks | Create task |
| GET | /tasks/:id | Get single task |
| PUT | /tasks/:id | Update task |
| DELETE | /tasks/:id | Delete task |

## Usage Examples

### Register a User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Response:
```json
{
  "id": 1,
  "email": "user@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### Create a Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Build a REST API"
  }'
```

Response:
```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Build a REST API",
  "completed": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### List Tasks
```bash
curl http://localhost:8080/tasks \
  -H "Authorization: Bearer <your-token>"
```

### Get Single Task
```bash
curl http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-token>"
```

### Update Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Build a REST API",
    "completed": true
  }'
```

### Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-token>"
```

Response: `204 No Content`

## Project Structure
```
task-api/
├── main.go                 # Entry point
├── internal/
│   ├── api/
│   │   ├── server.go       # Server setup
│   │   └── middleware.go   # Auth middleware
│   ├── handler/
│   │   ├── auth.go         # Register, login handlers
│   │   └── task.go         # Task CRUD handlers
│   ├── storage/
│   │   ├── storage.go      # Database connection
│   │   ├── user.go         # User queries
│   │   └── task.go         # Task queries
│   ├── auth/
│   │   ├── jwt.go          # Token generation/validation
│   │   └── password.go     # Password hashing
│   └── config/
│       └── config.go       # Environment config
├── migrations/
│   ├── 000001_create_users.up.sql
│   ├── 000002_create_tasks.up.sql
├── docker-compose.yml
├── Dockerfile
└── Makefile
```

## Error Responses
```json
{
  "error": "error message here"
}
```

| Status | Meaning |
|--------|---------|
| 400 | Bad request / validation error |
| 401 | Unauthorized / invalid token |
| 404 | Resource not found |
| 409 | Conflict (duplicate email) |
| 500 | Server error |

## License

MIT