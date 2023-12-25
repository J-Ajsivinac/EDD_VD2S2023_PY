package main

import (
	"backend/pkg/arbol"
	"backend/pkg/tabla"
	"backend/schemas"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaHash = tabla.TablaHash{Tabla: make(map[int]tabla.NodoHash), Capacidad: 7, Utilizacion: 0}
var arbolB *arbol.ArbolB = &arbol.ArbolB{Raiz: nil, Orden: 3}

func Login(c *fiber.Ctx) error {
	var login schemas.Login
	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	if login.Carnet == "admin" && login.Contrasena == "admin" {
		return c.JSON(fiber.Map{
			"message": "Login success",
			"carnet":  "ADMIN_202200135",
			"nombre":  "admin",
			"mode":    "admin",
		})
	}
	if !login.Estutor {
		carnet, _ := strconv.Atoi(login.Carnet)
		user, resp := tablaHash.BuscarUsuario(carnet, login.Contrasena)
		if resp {
			return c.JSON(fiber.Map{
				"message": "Login success",
				"carnet":  user.Carnet,
				"nombre":  user.Nombre,
				"mode":    "user",
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Credenciales Incorectas",
			})
		}
	} else {
		return c.JSON(fiber.Map{
			"message": "Login success",
			"carnet":  "---",
			"nombre":  "Tutor",
			"mode":    "tutor",
		})
	}
}

func cargarEstudiantes(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	tablaHash.LeerCSVFromReader(fileReader)

	return c.JSON(fiber.Map{
		"message": "Archivo cargado exitosamente",
	})
}

func cargarTutores(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	// tablaHash.LeerCSVFromReader(fileReader)
	arbolB.LeerCSV(fileReader)
	return c.JSON(fiber.Map{
		"message": "Archivo cargado exitosamente",
	})
}

func imprimir(c *fiber.Ctx) error {
	for i := 0; i < tablaHash.Capacidad; i++ {
		if usuario, existe := tablaHash.Tabla[i]; existe {
			fmt.Println("Posicion: ", i, " Carnet: ", usuario.Persona.Carnet, "Password: ", usuario.Persona.Password)
		}
	}
	return nil
}

func obtenerEstudiantes(c *fiber.Ctx) error {
	var estudiantes []schemas.UserData

	if tablaHash.Utilizacion == 0 {
		return c.Status(200).JSON(fiber.Map{
			"data":    nil,
			"message": "No hay estudiantes registrados",
		})
	}

	for i := 0; i < tablaHash.Capacidad; i++ {
		if nodo, existe := tablaHash.Tabla[i]; existe {
			cursosStr := strings.Join(nodo.Persona.Cursos[:], " - ")
			userD := schemas.UserData{Indice: nodo.Llave, Carnet: nodo.Persona.Carnet, Nombre: nodo.Persona.Nombre, Password: nodo.Persona.Password, Cursos: cursosStr}
			estudiantes = append(estudiantes, userD)
		}

	}
	response := fiber.Map{
		"data":    estudiantes,
		"message": "Estudiantes obtenidos exitosamente",
	}

	return c.JSON(response)
}

func AgregarB(c *fiber.Ctx) error {
	var nuevo schemas.AgregarArbol
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.Insertar(nuevo.Carnet, nuevo.Nombre, nuevo.Curso, nuevo.Password)
	return c.JSON(fiber.Map{
		"message": "Estudiante agregado exitosamente",
	})
}

func GraficarB(c *fiber.Ctx) error {
	arbolB.Graficar("arbolB")
	return c.JSON(&fiber.Map{
		"status":  200,
		"message": "Grafica Generada",
	})
}

func BuscarB(c *fiber.Ctx) error {
	type buscar struct {
		Carnet int `json:"carnet"`
	}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.Buscar(nuevo.Carnet)
	return c.JSON(fiber.Map{
		"message": "Estudiante encontrado",
	})
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", Login)
	app.Post("/upload", cargarEstudiantes)
	app.Get("/imprimir", imprimir)
	app.Get("/estudiantes", obtenerEstudiantes)
	app.Get("/graficarB", GraficarB)
	app.Post("/buscarB", BuscarB)
	app.Post("/agregarB", cargarTutores)
	app.Listen(":3000")
}
