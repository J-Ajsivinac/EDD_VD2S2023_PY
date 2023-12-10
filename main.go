package main

import (
	"Proyecto/pkg/cola"
	"Proyecto/pkg/listaD"
	"Proyecto/pkg/listaDCircular"
	"Proyecto/pkg/utilities"
	"fmt"
	"os"
	"strconv"
)

var listaE *listaD.ListaDobleE = &listaD.ListaDobleE{Primero: nil, Ultimo: nil, Longitud: 0}
var ColaPrioridad *cola.Cola = &cola.Cola{Inicio: nil, Longitud: 0}
var listaTutores *listaDCircular.ListaCircularDoble = &listaDCircular.ListaCircularDoble{Inicio: nil, Longitud: 0}

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
	opcion := 0
	salir := false

	for !salir {
		if ColaPrioridad.Longitud == 0 {
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
			fmt.Println("Login exitoso")
			for {
				menuUsuario()
				var opcionL int
				fmt.Scanln(&opcionL)
				if opcionL == 3 {
					fmt.Println("Saliendo del Usuario")
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
		titulos("Carga de Estudiantes Tutores")
		ruta := ""
		fmt.Print("Ingrese la ruta del archivo: ")
		fmt.Scanln(&ruta)
		ColaPrioridad.LeerArchivoTutores(ruta)
		utilities.MensajeConsola("Carga de Estudiantes Tutores exitosa", "verde")
	case 2:
		titulos("Carga de Estudiantes")
		ruta := ""
		fmt.Print("Ingrese la ruta del archivo: ")
		fmt.Scanln(&ruta)
		listaE.LeerArchivo(ruta)
		utilities.MensajeConsola("Carga de Estudiantes exitosa", "verde")
	case 3:
		titulos("Carga Cursos al sistema")
	case 4:
		titulos("Control de Estudiantes tutores")
		MenuAceptar()
	case 5:
		titulos("Reportes Estructuras")
		fmt.Println(listaTutores.Longitud)
		listaE.Reporte()
		listaTutores.Reporte()
	default:
		utilities.MensajeConsola("Opcion no valida", "rojo")
	}
}

func opcionesUsuario(opcion int) {
	switch opcion {
	case 1:
		titulos("Ver Tutores Disponibles")
		listaTutores.Recorrer()
	case 2:
		titulos("Asignarse a Tutores")
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
