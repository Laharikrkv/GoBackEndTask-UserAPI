# GoBackEndTask-UserAPI
A simple REST API built using Go to manage users with name and date of birth (DOB).
This project is structured to follow real-world backend practices such as clean layering, SQLC-based DB access, validation, and structured logging.

##Features
Create, update, delete, and fetch users
Store only name and dob in the database
Calculate age dynamically using Go’s time package
Clean project structure (handler → service → repository)
SQL-safe database access using SQLC
Input validation using go-playground/validator
Structured logging with Uber Zap
Proper HTTP status codes and error handling

##Tech Stack
Go + Fiber(HTTP Framework)
PostgreSQL + SQLC(Query Functions Generator)
Uber Zap for logging
go-playground/validator for input validation

##Project Structure
/cmd/server/main.go
/config/
 /db/
   ├── migrations/
   ├── queries/
   └── sqlc/        # SQLC generated code
/internal/
 ├── handler/       # HTTP handlers
 ├── service/       # Business logic
 ├── repository/    # DB access abstraction
 ├── routes/        # Route registration
 ├── middleware/  
 ├── models/        # Request/response models
 └── logger/        # Zap logger setup
/sqlc.yaml

##Database Schema 
CREATE TABLE users (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob  DATE NOT NULL
);

##Prerequisites
Before you begin, ensure you have the following installed:
  Go 1.21 or higher
  Docker for Postgresql
  Migrations
  SQLC (for regenerating queries if needed)
  
##Setup and Run




