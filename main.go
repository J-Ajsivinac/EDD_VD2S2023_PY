package main

import (
	"Proyecto/pkg/arbolAVL"
	"Proyecto/pkg/cola"
	"Proyecto/pkg/listaD"
	"Proyecto/pkg/listaDCircular"
	"Proyecto/pkg/matrizE"
	"Proyecto/pkg/utilities"
	"fmt"
	"os"
	"strconv"
)

var listaE *listaD.ListaDobleE = &listaD.ListaDobleE{Primero: nil, Ultimo: nil, Longitud: 0}
var ColaPrioridad *cola.Cola = &cola.Cola{Inicio: nil, Longitud: 0}
var listaTutores *listaDCircular.ListaCircularDoble = &listaDCircular.ListaCircularDoble{Inicio: nil, Longitud: 0}
var matriz *matrizE.Matriz = &matrizE.Matriz{Raiz: &matrizE.NodoMatriz{PosX: -1, PosY: -1, Dato: &matrizE.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var cookies int = 0
var arbolCursos *arbolAVL.Arbol = &arbolAVL.Arbol{Raiz: nil}

func titulos(titulo string) {
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════════╗")
	fmt.Printf(" ║ 🔹 %-35s ║\n", titulo)
	fmt.Println(" ╚════════════════════════════════════════╝")
	fmt.Println("")
}

func menu() {
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════╗")
	fmt.Println(" ║            Menu Principal          ║")
	fmt.Println(" ╠════════════════════════════════════╣")
	fmt.Println(" ║                                    ║")
	fmt.Println(" ║         1. Login                   ║")
	fmt.Println(" ║         2. Salir                   ║")
	fmt.Println(" ║                                    ║")
	fmt.Println(" ╚════════════════════════════════════╝")
	fmt.Println("")
	fmt.Print("Ingrese una opcion: ")
}

func menuReportes() {
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════╗")
	fmt.Println(" ║              Reportes              ║")
	fmt.Println(" ╠════════════════════════════════════╣")
	fmt.Println(" ║                                    ║")
	fmt.Println(" ║         1. Alumnos                 ║")
	fmt.Println(" ║         2. Tutores                 ║")
	fmt.Println(" ║         3. Asignaciones            ║")
	fmt.Println(" ║         4. Cursos                  ║")
	fmt.Println(" ║         5. Salir                   ║")
	fmt.Println(" ║                                    ║")
	fmt.Println(" ╚════════════════════════════════════╝")
	fmt.Println("")
	fmt.Print("Ingrese una opcion: ")
}

func menuAdmin() {
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════════╗")
	fmt.Println(" ║             Administrador              ║")
	fmt.Println(" ╠════════════════════════════════════════╣")
	fmt.Println(" ║                                        ║")
	fmt.Println(" ║     1. Carga de Estudiantes Tutores    ║")
	fmt.Println(" ║     2. Carga de Estudiantes            ║")
	fmt.Println(" ║     3. Carga Cursos al sistema         ║")
	fmt.Println(" ║     4. Control de Estudiantes tutores  ║")
	fmt.Println(" ║     5. Reportes Estructuras            ║")
	fmt.Println(" ║     6. Salir                           ║")
	fmt.Println(" ║                                        ║")
	fmt.Println(" ╚════════════════════════════════════════╝")
	fmt.Println("")
	fmt.Print("Ingrese una opcion: ")
}

func menuUsuario() {
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════════╗")
	fmt.Println(" ║                 Usuario                ║")
	fmt.Println(" ╠════════════════════════════════════════╣")
	fmt.Println(" ║                                        ║")
	fmt.Println(" ║     1. Ver Tutores Disponibles         ║")
	fmt.Println(" ║     2. Asignarse a Tutores             ║")
	fmt.Println(" ║     3. Salir                           ║")
	fmt.Println(" ║                                        ║")
	fmt.Println(" ╚════════════════════════════════════════╝")
	fmt.Println("")

	fmt.Print("Ingrese una opcion: ")
}

func MenuAceptar() {
	titulos("Control de Estudiantes tutores")
	opcion := 0
	salir := false

	for !salir {
		if ColaPrioridad.Longitud == 0 {
			fmt.Print("\033[H\033[2J")
			break
		}
		ColaPrioridad.Primero()
		fmt.Scanln(&opcion)
		if opcion == 1 {
			listaTutores.InsertarOrdenado(ColaPrioridad.Inicio.Tutor.Carnet, ColaPrioridad.Inicio.Tutor.Nombre, ColaPrioridad.Inicio.Tutor.Curso, ColaPrioridad.Inicio.Tutor.Nota)
			ColaPrioridad.Descolar()
		} else if opcion == 2 {
			ColaPrioridad.Descolar()
		} else {
			salir = true
		}
	}

}

func login(usuario string, contra string) {
	if usuario == "admin" && contra == "admin" {
		for {
			menuAdmin()
			var opcionL int
			fmt.Scanln(&opcionL)
			if opcionL == 6 {
				fmt.Println("\n Saliendo del administrador \n")
				break
			}
			opcionesAdmin(opcionL)
		}
	} else {
		user, _ := strconv.Atoi(usuario)
		valor, _ := strconv.Atoi(contra)
		if listaE.LoginUser(user, valor) {
			cookies = user
			fmt.Println("Login exitoso")
			for {
				menuUsuario()
				var opcionL int
				fmt.Scanln(&opcionL)
				if opcionL == 3 {
					fmt.Println("Saliendo del Usuario")
					cookies = 0
					break
				}
				opcionesUsuario(opcionL)
			}
		} else {
			utilities.MensajeConsola("Credenciales Incorrectas ", "rojo")
		}
	}
}

func opcionesLogin(opcion int) {
	switch opcion {
	case 1:
		fmt.Print("\033[H\033[2J")
		titulos("Login")
		fmt.Print(" 🙍‍♂️ Usuario: ")
		var usuario string
		fmt.Scanln(&usuario)
		fmt.Print(" 🔒 Password: ")
		var contra string
		fmt.Scanln(&contra)
		login(usuario, contra)
	case 2:
		fmt.Println("\n Saliendo del sistema \n")
		os.Exit(0)
	default:
		utilities.MensajeConsola("Opcion no valida", "rojo")
	}
}

func opcionesAdmin(opcion int) {
	switch opcion {
	case 1:
		fmt.Print("\033[H\033[2J")
		titulos("Carga de Estudiantes Tutores")
		ruta := ""
		fmt.Print("Ingrese la ruta del archivo: ")
		fmt.Scanln(&ruta)
		ColaPrioridad.LeerArchivoTutores(ruta)
	case 2:
		fmt.Print("\033[H\033[2J")
		titulos("Carga de Estudiantes")
		ruta := ""
		fmt.Print("Ingrese la ruta del archivo: ")
		fmt.Scanln(&ruta)
		listaE.LeerArchivo(ruta)
	case 3:
		cargarCursos()
	case 4:
		MenuAceptar()
	case 5:
		generarGraficas()
	default:
		utilities.MensajeConsola("Opcion no valida", "rojo")
	}
}

func generarGraficas() {
	fmt.Print("\033[H\033[2J")
	titulos("Reportes Estructuras")
	opcion := 0
	salir := false

	for !salir {
		menuReportes()
		fmt.Scanln(&opcion)
		if opcion == 1 {
			listaE.Reporte()
		} else if opcion == 2 {
			listaTutores.Reporte()
		} else if opcion == 3 {
			matriz.Reporte()
		} else if opcion == 4 {
			arbolCursos.Reporte()
		} else if opcion == 5 {
			salir = true
		} else {
			utilities.MensajeConsola("Opcion no valida", "rojo")
		}
	}
}

func cargarCursos() {
	fmt.Print("\033[H\033[2J")
	titulos("Carga Cursos al sistema")
	fmt.Print("Ingrese la ruta del archivo: ")
	ruta := ""
	fmt.Scanln(&ruta)
	arbolCursos.LeerJson(ruta)

}

func asignarCursos() {
	fmt.Print("\033[H\033[2J")
	titulos("Asignarse a Tutores")
	var curso string
	salirCursos := false
	for !salirCursos {
		fmt.Print("Ingrese el codigo del curso [0 para salirCursos]: ")
		fmt.Scanln(&curso)
		if curso == "0" {
			salirCursos = true
			break
		}
		if !arbolCursos.Busqueda(curso) {
			utilities.MensajeConsola("El curso no existe", "rojo")
			continue
		}
		respuesta := listaTutores.Buscar(curso)
		if respuesta == nil {
			utilities.MensajeConsola("No hay tutores para el curso "+curso, "rojo")
			continue
		}
		fmt.Println("Usuario: ", cookies)
		matriz.Insertar_Elemento(cookies, respuesta.Estudiante.Carnet, curso)

	}
}

func opcionesUsuario(opcion int) {
	switch opcion {
	case 1:
		fmt.Print("\033[H\033[2J")
		titulos("Ver Tutores Disponibles")
		listaTutores.Recorrer()
	case 2:
		asignarCursos()
	default:
		utilities.MensajeConsola("Opcion no valida", "rojo")
	}
}

func main() {
	for {
		menu()
		var opcion int
		fmt.Scanln(&opcion)
		opcionesLogin(opcion)
	}
}
