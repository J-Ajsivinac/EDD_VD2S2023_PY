package main

import (
	"backend/pkg/arbol"
	"backend/pkg/arbolM"
	"backend/pkg/grafoA"
	"backend/pkg/tabla"
	"backend/schemas"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaHash = tabla.TablaHash{Tabla: make(map[int]tabla.NodoHash), Capacidad: 7, Utilizacion: 0}
var arbolB *arbol.ArbolB = &arbol.ArbolB{Raiz: nil, Orden: 3}
var grafo *grafoA.Grafo = &grafoA.Grafo{Principal: nil}
var arbolMerkle *arbolM.ArbolMerkle = &arbolM.ArbolMerkle{RaizMerkle: nil, BloqueDeDatos: nil, CantidadBloques: 0}

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
	carnet, _ := strconv.Atoi(login.Carnet)
	if !login.Estutor {
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
		resp := arbolB.Buscar(carnet)
		if resp != nil {
			if resp.Password == encriptarPassword(login.Contrasena) {
				return c.JSON(fiber.Map{
					"message": "Login success",
					"carnet":  resp.Carnet,
					"nombre":  resp.Nombre,
					"mode":    "tutor",
				})
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Credenciales Incorectas",
				})
			}
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Credenciales Incorectas",
	})
}

func encriptarPassword(password string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(password))
	hexaString = hex.EncodeToString(h.Sum(nil))
	return hexaString
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

	arbolB.LeerCSV(fileReader)
	return c.JSON(fiber.Map{
		"message": "Archivo cargado exitosamente",
	})
}

func cargarCursos(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	fileReader, err := file.Open()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}
	defer fileReader.Close()

	value, resp := grafo.Lectura(fileReader)
	if !value {
		return c.Status(400).JSON(fiber.Map{
			"message": resp,
		})
	}
	grafo.Reporte()
	return c.JSON(fiber.Map{
		"message": "Archivo cargado exitosamente",
	})
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
			userD := schemas.UserData{Indice: nodo.Llave, Carnet: nodo.Persona.Carnet, Nombre: nodo.Persona.Nombre, Password: nodo.Persona.Password, Cursos: nodo.Persona.Cursos[:]}
			estudiantes = append(estudiantes, userD)
		}

	}
	response := fiber.Map{
		"data":    estudiantes,
		"message": "Estudiantes obtenidos exitosamente",
	}

	return c.JSON(response)
}

func GraficarB(c *fiber.Ctx) error {
	arbolB.Graficar()
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

func AceptarL(c *fiber.Ctx) error {
	var nuevo schemas.AceptarLibros
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolMerkle.AgregarBloque(nuevo.Estado, nuevo.Nombre, nuevo.Carnet)
	return c.JSON(fiber.Map{
		"message": "Bloque agregado exitosamente",
	})
}

func GraficarM(c *fiber.Ctx) error {
	arbolMerkle.GenerarArbol()
	arbolMerkle.Graficar()
	return c.JSON(&fiber.Map{
		"status":  200,
		"message": "Grafica Generada",
	})
}

func AgregarL(c *fiber.Ctx) error {
	var nuevo schemas.AgregarLibro
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.GuardarLibro(arbolB.Raiz.Primero, nuevo.Nombre, nuevo.Contenido, nuevo.Carnet)
	return c.JSON(fiber.Map{
		"message": "Libro agregado exitosamente",
	})
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", Login)
	//Administrador
	admin := app.Group("/admin")
	admin.Post("/cargar-e", cargarEstudiantes)
	admin.Get("/obtener-e", obtenerEstudiantes)
	admin.Post("/cargar-t", cargarTutores)
	admin.Post("/cargar-c", cargarCursos)
	admin.Get("/graficar-arbolB", GraficarB)
	admin.Post("/buscar-arbolB", BuscarB)
	admin.Post("/aceptar-arbolM", AceptarL)
	admin.Get("/graficar-arbolM", GraficarM)
	//Tutor
	tutor := app.Group("/tutor")
	tutor.Post("/agregar-arbolB", AgregarL)
	//Estudiante
	app.Listen(":3000")
}
