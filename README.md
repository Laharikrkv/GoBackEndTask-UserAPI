# GoBackEndTask-UserAPI

A simple REST API built using Go to manage users with name and date of birth (DOB). This project is structured to follow real-world backend practices such as clean layering, SQLC-based DB access, validation, and structured logging.

---

## ✨ Features

- ✅ **Create, update, delete, and fetch users** - Full CRUD operations
- 📅 **Store only name and DOB in the database** - Minimal data storage
- 🎂 **Calculate age dynamically** - Uses Go's time package for real-time calculation
- 🏗️ **Clean project structure** - Handler → Service → Repository pattern
- 🔒 **SQL-safe database access using SQLC** - Type-safe queries
- ✔️ **Input validation** - Using go-playground/validator
- 📝 **Structured logging with Uber Zap** - Professional logging
- 🚀 **Proper HTTP status codes and error handling** - RESTful best practices

---

## 🔧 Tech Stack

- **Go** + **Fiber** (HTTP Framework)
- **PostgreSQL** + **SQLC** (Query Functions Generator)
- **Uber Zap** for logging
- **go-playground/validator** for input validation

---

## 📂 Project Structure

```
/cmd/server/main.go
/config/
/db/
  ├── migrations/
  ├── queries/  
  └── sqlc/               # SQLC generated code
/internal/
  ├── handler/            # HTTP handlers
  ├── service/            # Business logic
  ├── repository/         # DB access abstraction
  ├── routes/             # Route registration
  ├── middleware/
  ├── models/             # Request/response models
  └── logger/             # Zap logger setup
/sqlc.yaml
```

---

## 📊 Database Schema

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);
```

---

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.21 or higher
- **Docker** (for PostgreSQL)
- **Migrations** tool (optional)
- **SQLC** (for regenerating queries if needed)

**Check your versions:**

```bash
go version
docker --version
```

---

## 🚀 Setup and Run

### Step 1: Clone the Repository

```bash
git clone https://github.com/Laharikrkv/GoBackEndTask-UserAPI.git
cd GoBackEndTask-UserAPI
```

### Step 2: Install Dependencies

```bash
go mod download
```

### Step 3: Start PostgreSQL with Docker

```bash
docker run --name postgres-userdb \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=usersdb \
  -p 5432:5432 \
  -d postgres:15
```

**Verify the database is running:**

```bash
docker ps
```

You should see a container named `postgres-userdb` with status `Up`.

### Step 4: Run Database Migrations

If you have migration files in `db/migrations/`, run:

```bash
migrate -path db/migrations \
  -database "postgresql://postgres:postgres@localhost:5432/usersdb?sslmode=disable" \
  up
```

**Or**, if migrations are handled in the application code, skip this step.

### Step 5: Run the Application

```bash
go run cmd/server/main.go
```

**Expected output:**

```
2024/12/18 10:30:45 INFO Application started
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.52.0                   │ 
 │               http://127.0.0.1:8080               │ 
 └───────────────────────────────────────────────────┘ 
```

The API is now running at **`http://localhost:8080`** 🎉

---

## 🔌 API Endpoints

### 1. Create User

**POST** `/users`

**Request:**

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

**Response:** `201 Created`

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### 2. Get User by ID

**GET** `/users/:id`

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 34
}
```

---

### 3. List All Users

**GET** `/users`

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]
```

---

### 4. Update User

**PUT** `/users/:id`

**Request:**

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### 5. Delete User

**DELETE** `/users/:id`

**Response:** `204 No Content`

---

## 🧪 Testing the API

### Using cURL

```bash
# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","dob":"1990-05-10"}'

# Get all users
curl http://localhost:8080/users

# Get user by ID
curl http://localhost:8080/users/1

# Update user
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","dob":"1991-03-15"}'

# Delete user
curl -X DELETE http://localhost:8080/users/1
```

---

## 🛠️ Development

### Regenerate SQLC Queries

If you modify SQL queries:

```bash
sqlc generate
```

### Create New Migration

```bash
migrate create -ext sql -dir db/migrations -seq migration_name
```



---

## 🐛 Troubleshooting

### Database Connection Failed

**Error:** `connection refused`

**Solution:**

```bash
# Check if container is running
docker ps

# Start the container
docker start postgres-userdb

# Check logs
docker logs postgres-userdb
```

### Port Already in Use

**Error:** `bind: address already in use`

**Solution:**

```bash
# Find process on port 8080
lsof -ti:8080

# Kill the process
kill -9 <PID>
```

---

## 📦 Dependencies

```bash
go get github.com/gofiber/fiber/v2
go get github.com/go-playground/validator/v10
go get go.uber.org/zap
```

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


## 👤 Author

**Lahari**

- GitHub: [@Laharikrkv](https://github.com/Laharikrkv)
- Project: [GoBackEndTask-UserAPI](https://github.com/Laharikrkv/GoBackEndTask-UserAPI)

---

⭐ **If you found this project helpful, please give it a star!**
