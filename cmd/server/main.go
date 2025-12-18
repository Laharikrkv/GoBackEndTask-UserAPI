package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"go-api-task/db/sqlc"
	"go-api-task/internal/handler"
	"go-api-task/internal/repository"
	"go-api-task/internal/routes"
	"go-api-task/internal/service"
	"go-api-task/internal/logger"
)

func main() {

	dbConn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/usersdb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(dbConn)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// logger initialization
	logger.Init()
	defer logger.Log.Sync()

	logger.Log.Info("Application started")




	app := fiber.New()

	routes.RegisterUserRoutes(app, userHandler)

	log.Fatal(app.Listen(":8080"))
}
