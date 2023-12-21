package main

import (
	"backend/schemas"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Login(c *fiber.Ctx) error {
	var login schemas.Login
	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	if login.Usuario == "admin" && login.Contrasena == "admin" {
		return c.JSON(fiber.Map{
			"message": "Login success",
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Post("/login", Login)
	app.Listen(":3000")
}
