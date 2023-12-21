package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Listen(":3000")
}
