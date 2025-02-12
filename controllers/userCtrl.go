package controllers

import (
	"testeGO/repository"
	"testeGO/models"
	"github.com/gofiber/fiber/v2"
);

func AddUser(c *fiber.Ctx) error {

	var user models.User;

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos!"});
	}

	err := repository.AddUser(user);
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar usuário!"});
	}

	return c.Status(200).JSON(fiber.Map{"message": "Usuário criado com sucesso!"});
}

func GetUsers(c *fiber.Ctx) error {
	users, err := repository.GetUsers();
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao recuperar usuarios!"});
	}
	return c.Status(200).JSON(fiber.Map{"usuarios": users});
}