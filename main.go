package main

import (
	"log"
	"testeGO/config"
	"testeGO/routes"
	"github.com/gofiber/fiber/v2"
);

func main() {
	config.ConnectDb();
	defer config.CloseDb();

	app := fiber.New();
	routes.SetupUserRoutes(app);
	
	log.Println("Servidor rodando em https://localhost:3000");
	log.Fatal(app.Listen(":3000"));
}