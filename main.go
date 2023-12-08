package main

import (
	"Proyecto/pkg"
	"fmt"
	"os"
	"strconv"
)

var listaE pkg.ListaDobleE

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

func login(usuario string, contra string) {
	if usuario == "admin" && contra == "admin" {
		for {
			menuAdmin()
			var opcionL int
			fmt.Scanln(&opcionL)
			if opcionL == 6 {
				fmt.Println("Saliendo del administrador")
				break
			}
			opcionesAdmin(opcionL)
		}
	} else {
		valor, _ := strconv.Atoi(contra)
		if listaE.LoginUser(valor) {
			fmt.Println("Login exitoso")
			for {
				menuUsuario()
				var opcionL int
				fmt.Scanln(&opcionL)
				if opcionL == 3 {
					fmt.Println("Saliendo del administrador")
					break
				}
				opcionesUsuario(opcionL)
			}
		}
		fmt.Println("Usuario o contraseña incorrecta")
	}
}

func opcionesLogin(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Login")
		fmt.Print("Usuario:")
		var usuario string
		fmt.Scanln(&usuario)
		fmt.Print("Password:")
		var contra string
		fmt.Scanln(&contra)
		login(usuario, contra)
	case 2:
		fmt.Println("Salir")
		os.Exit(0)
	default:
		fmt.Println("Opcion no valida")
	}
}

func opcionesAdmin(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Carga de Estudiantes Tutores")
	case 2:
		fmt.Println("Carga de Estudiantes")
		listaE.LeerArchivo("inputE.csv")
		listaE.Recorrer()
	case 3:
		fmt.Println("Carga Cursos al sistema")
	case 4:
		fmt.Println("Carga Cursos al sistema")
	case 5:
		fmt.Println("Carga Cursos al sistema")
	default:
		fmt.Println("Opcion no valida")
	}
}

func opcionesUsuario(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Ver tutores Disponibles")
	case 2:
		fmt.Println("Asignarse a Tutores")
	default:
		fmt.Println("Opcion no valida")
	}
}

func main() {
	// listaE := pkg.ListaEstu{Primero: nil, Ultimo: nil, Longitud: 0}
	// listaE.Insertar(&pkg.Estudiante{Carnet: 201801021, Nombre: "Juan"})
	// listaE.Insertar(&pkg.Estudiante{Carnet: 201801022, Nombre: "Juan1"})
	// listaE.Insertar(&pkg.Estudiante{Carnet: 201801023, Nombre: "Juan2"})
	// listaE.Recorrer()
	// fmt.Println("***************")
	// listaE.Recorrer2()
	for {
		menu()
		var opcion int
		fmt.Scanln(&opcion)
		fmt.Println(&opcion)
		opcionesLogin(opcion)
	}
}
