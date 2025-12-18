package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api-task/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	user := app.Group("/users")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUser)
	user.Get("/:id", userHandler.GetUserById)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
