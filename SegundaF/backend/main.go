package main

import (
	"backend/pkg/tabla"
	"backend/schemas"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var tablaHash = tabla.TablaHash{Tabla: make(map[int]tabla.NodoHash), Capacidad: 7, Utilizacion: 0}

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
		carnet, _ := strconv.Atoi(login.Usuario)
		user, resp := tablaHash.BuscarUsuario(carnet, login.Contrasena)
		if resp {
			return c.JSON(fiber.Map{
				"message": "Login success",
				"carnet":  user.Carnet,
				"nombre":  user.Nombre,
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": "No ser registro",
			})
		}

	}
}

func cargarEstudiantes(c *fiber.Ctx) error {
	// Obtén el archivo del formulario enviado en la solicitud
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Crea un lector para el archivo
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	// Llama a la función LeerCSV con el lector del archivo
	tablaHash.LeerCSVFromReader(fileReader)

	return nil
}

func imprimir(c *fiber.Ctx) error {
	for i := 0; i < tablaHash.Capacidad; i++ {
		if usuario, existe := tablaHash.Tabla[i]; existe {
			fmt.Println("Posicion: ", i, " Carnet: ", usuario.Persona.Carnet, "Password: ", usuario.Persona.Password)
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()
	app.Post("/login", Login)
	app.Post("/upload", cargarEstudiantes)
	app.Get("/imprimir", imprimir)

	app.Listen(":3000")
}
