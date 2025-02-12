package routes

import (
	"testeGO/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users");

	userRoutes.Post("/", controllers.AddUser);
	userRoutes.Get("/", controllers.GetUsers);
}