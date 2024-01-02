package main

import (
	"backend/pkg/arbol"
	"backend/pkg/arbolM"
	"backend/pkg/grafoA"
	"backend/pkg/tabla"
	"backend/schemas"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaHash = tabla.TablaHash{Tabla: make(map[int]tabla.NodoHash), Capacidad: 7, Utilizacion: 0}
var arbolB *arbol.ArbolB = &arbol.ArbolB{Raiz: nil, Orden: 3}
var grafo *grafoA.Grafo = &grafoA.Grafo{Principal: nil}
var arbolMerkle *arbolM.ArbolMerkle = &arbolM.ArbolMerkle{RaizMerkle: nil, BloqueDeDatos: nil, CantidadBloques: 0}
var listaSimple *arbol.ListaSimple

func Login(c *fiber.Ctx) error {
	var login schemas.Login
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
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
		user, resp := tablaHash.BuscarUsuario(carnet)
		fmt.Println(user, resp)
		if resp && user.Password == encriptarPassword(login.Contrasena) {
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
		arbolB.Buscar(carnet, listaSimple)
		if listaSimple.Longitud > 0 {
			if listaSimple.Inicio.Tutor.Usuario.Password == encriptarPassword(login.Contrasena) {
				return c.JSON(fiber.Map{
					"message": "Login success",
					"carnet":  listaSimple.Inicio.Tutor.Usuario.Carnet,
					"nombre":  listaSimple.Inicio.Tutor.Usuario.Nombre,
					"mode":    "tutor",
				})
			} else {
				fmt.Println(listaSimple.Inicio.Tutor.Usuario.Password, "encontrador")
				fmt.Println(encriptarPassword(login.Contrasena), "encriptado")
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

func GenerarGraficas(c *fiber.Ctx) error {
	var grafica schemas.GraphR
	err := c.BodyParser(&grafica)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Se esperaba un JSON",
		})
	}
	if grafica.Grafica == "ArbolB" {
		img := arbolB.Graficar()
		if len(img) == 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "No se pudo generar la grafica",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Grafica Generada",
			"graph":   img,
		})
	} else if grafica.Grafica == "Merkle" {
		arbolMerkle.GenerarArbol()
		img := arbolMerkle.Graficar()
		if len(img) == 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "No se pudo generar la grafica",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Grafica Generada",
			"graph":   img,
		})
	} else if grafica.Grafica == "Grafo" {
		img := grafo.Graficar()
		if len(img) == 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "No se pudo generar la grafica",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Grafica Generada",
			"graph":   img,
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro la grafica",
		})
	}
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

func obtenerLibrosTutor(c *fiber.Ctx) error {
	type buscar struct {
		Carnet int `json:"carnet"`
	}
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.Buscar(nuevo.Carnet, listaSimple)
	if listaSimple.Longitud == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro el tutor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Libros obtenidos exitosamente",
		"libros":  listaSimple.Inicio.Tutor.Usuario.Libros,
	})
}

func aceptarLibroAdmin(c *fiber.Ctx) error {
	var nuevo schemas.AceptarLibros
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.CambiarEstadoLibro(arbolB.Raiz.Primero, nuevo.Carnet, nuevo.Nombre, nuevo.Estado)
	arbolMerkle.AgregarBloque(nuevo.Estado, nuevo.Nombre, nuevo.Carnet)
	return c.JSON(fiber.Map{
		"message": "Libro" + nuevo.Estado + " exitosamente",
	})

}

func agregarContenido(c *fiber.Ctx) error {
	var nuevo schemas.ContenidoPub
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.GuardarPublicacion(arbolB.Raiz.Primero, nuevo.Contenido, nuevo.Carnet)
	return c.JSON(fiber.Map{
		"message": "Publicacion agregada exitosamente",
	})
}

func obtenerPublicaciones(c *fiber.Ctx) error {
	type buscar struct {
		Carnet int `json:"carnet"`
	}
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	arbolB.Buscar(nuevo.Carnet, listaSimple)
	if listaSimple.Longitud == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro el tutor",
		})
	}
	return c.JSON(fiber.Map{
		"message":       "Libros obtenidos exitosamente",
		"publicaciones": listaSimple.Inicio.Tutor.Usuario.Publicaciones,
	})
}

func obtenerTLibros(c *fiber.Ctx) error {
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	arbolB.ObtenerLibros(arbolB.Raiz.Primero, listaSimple)
	if listaSimple.Longitud == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro el tutor",
		})
	}
	usuarios := listaSimple.ConverirUsuarios()
	return c.JSON(fiber.Map{
		"message": "Libros obtenidos exitosamente",
		"data":    usuarios,
	})
}

func validarCurso(c *fiber.Ctx) error {
	var nuevo schemas.BuscarCurso
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	notFound := false
	for i := 0; i < len(nuevo.Codigo); i++ {
		if !grafo.Buscar(nuevo.Codigo[i]) {
			notFound = true
		}
	}
	if !notFound {
		return c.JSON(fiber.Map{
			"message": "Cursos validados exitosamente",
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontraron los cursos",
		})
	}
}

func buscarLibrosCurso(c *fiber.Ctx) error {
	type buscar struct {
		Codigo []string `json:"codigo"`
	}
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	for i := 0; i < len(nuevo.Codigo); i++ {
		arbolB.LlamarBuscarLibros(nuevo.Codigo[i], listaSimple)
	}

	libros := listaSimple.ConverirUsuarios()

	if listaSimple.Longitud == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro el tutor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Libros obtenidos exitosamente",
		"libros":  libros,
	})
}

func buscarLibrosAprobados(c *fiber.Ctx) error {
	type buscar struct {
		Codigo []string `json:"codigo"`
	}
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	for i := 0; i < len(nuevo.Codigo); i++ {
		arbolB.LlamarBuscarLibros(nuevo.Codigo[i], listaSimple)
	}

	libros := listaSimple.ConverirLibros()

	if listaSimple.Longitud == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No se encontro el tutor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Libros obtenidos exitosamente",
		"libros":  libros,
	})
}

func buscarEstudiante(c *fiber.Ctx) error {
	type buscar struct {
		Carnet int `json:"carnet"`
	}
	listaSimple = &arbol.ListaSimple{Inicio: nil, Longitud: 0}
	var nuevo buscar
	err := c.BodyParser(&nuevo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	user, resp := tablaHash.BuscarUsuario(nuevo.Carnet)
	if resp {
		return c.JSON(fiber.Map{
			"message": "Login success",
			"cursos":  user.Cursos,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No se encontro el usuario",
		})
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/reportes", "./reportes")
	app.Post("/login", Login)
	app.Get("/obtener-tlibros", obtenerTLibros)
	app.Post("/validar-cursos", validarCurso)
	//Administrador
	admin := app.Group("/admin")
	admin.Post("/cargar-e", cargarEstudiantes)
	admin.Get("/obtener-e", obtenerEstudiantes)
	admin.Post("/cargar-t", cargarTutores)
	admin.Post("/cargar-c", cargarCursos)
	admin.Post("/aceptar-arbolM", AceptarL)
	admin.Post("/graficar", GenerarGraficas)
	admin.Post("/aceptar-libro", aceptarLibroAdmin)
	//Tutor
	tutor := app.Group("/tutor")
	tutor.Post("/agregar-arbolB", AgregarL)
	tutor.Post("/obtener-libros", obtenerLibrosTutor)
	tutor.Post("/agregar-contenido", agregarContenido)
	tutor.Post("/obtener-publicaciones", obtenerPublicaciones)

	//Estudiante
	estudiante := app.Group("/estudiante")
	estudiante.Post("/buscar-libros", buscarLibrosCurso)
	estudiante.Post("/buscar-estudiante", buscarEstudiante)
	estudiante.Post("/buscar-libros-aprobados", buscarLibrosAprobados)
	app.Listen(":3000")
}
