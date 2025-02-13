package routes

import (
	"testeGO/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users");

	userRoutes.Post("/", controllers.AddUser);
	userRoutes.Get("/", controllers.GetUsers);
	userRoutes.Put("/:cpf", controllers.AlterUser);
	userRoutes.Delete("/:cpf", controllers.DeleteUser);
}