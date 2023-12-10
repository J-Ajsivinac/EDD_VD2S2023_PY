package listaDCircular

import (
	"Proyecto/pkg"
	"Proyecto/pkg/utilities"
	"fmt"
	"strconv"
)

type ListaCircularDoble struct {
	Inicio   *NodoCircularE
	Longitud int
}

func (l *ListaCircularDoble) InsertarOrdenado(carnet int, nombre string, curso string, nota int) {
	nuevoAlumno := &EstudianteTutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCircularE{Estudiante: nuevoAlumno}

	var existe, esMayor = l.Validar(nuevoAlumno)
	if existe && !esMayor {
		utilities.MensajeConsola("El curso ya tiene un tutor con una nota mayor o igual a la que se desea ingresar", "amarillo")
		return
	}

	if l.Longitud == 0 {
		nuevoNodo.Siguiente = nuevoNodo
		nuevoNodo.Anterior = nuevoNodo
		l.Inicio = nuevoNodo
		l.Longitud++
		if existe && esMayor {
			utilities.MensajeConsola("Se sustituyo el tutor del curso, "+curso, "amarillo")
		} else {
			utilities.MensajeConsola("Se agrego un nuevo tutor al curso, "+curso, "verde")
		}
		return
	}

	if carnet < l.Inicio.Estudiante.Carnet {
		nuevoNodo.Siguiente = l.Inicio
		nuevoNodo.Anterior = l.Inicio.Anterior
		l.Inicio.Anterior.Siguiente = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Inicio = nuevoNodo
		l.Longitud++
		if existe && esMayor {
			utilities.MensajeConsola("Se sustituyo el tutor del curso, "+curso, "amarillo")
		} else {
			utilities.MensajeConsola("Se agrego un nuevo tutor al curso, "+curso, "verde")
		}
		return
	}

	actual := l.Inicio
	for actual.Siguiente != l.Inicio && nuevoAlumno.Carnet > actual.Siguiente.Estudiante.Carnet {
		actual = actual.Siguiente
	}
	nuevoNodo.Siguiente = actual.Siguiente
	nuevoNodo.Anterior = actual
	actual.Siguiente.Anterior = nuevoNodo
	actual.Siguiente = nuevoNodo
	l.Longitud++
	if existe && esMayor {
		utilities.MensajeConsola("Se sustituyo el tutor del curso, "+curso, "amarillo")
	} else {
		utilities.MensajeConsola("Se agrego un nuevo tutor al curso, "+curso, "verde")
	}
}

func (lista *ListaCircularDoble) Recorrer() {
	if lista.Inicio == nil {
		utilities.MensajeConsola("No hay mas Tutores para mostrar", "amarillo")
		return
	}
	formato := " ║ %-10s ║ %-32s ║ %-6s ║ %-34s ║\n"
	actual := lista.Inicio
	fmt.Println(" ╔═════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf(formato, "Carnet", "Nombre", "Código", "Nombre del Curso")
	println(" ╠═════════════════════════════════════════════════════════════════════════════════════════════╣")
	for {
		fmt.Printf(formato, strconv.Itoa(actual.Estudiante.Carnet), actual.Estudiante.Nombre, actual.Estudiante.Curso, "Curso Nombre")
		actual = actual.Siguiente
		if actual == lista.Inicio {
			break
		}
	}
	println(" ╚═════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// existencia, nota
func (l *ListaCircularDoble) Validar(estudiante *EstudianteTutor) (bool, bool) {
	if l.Inicio == nil {
		return false, false
	}

	existe := false
	aux := l.Inicio

	for {
		if aux.Estudiante.Curso == estudiante.Curso {
			existe = true
			if estudiante.Nota >= aux.Estudiante.Nota {
				aux.Anterior.Siguiente = aux.Siguiente
				aux.Siguiente.Anterior = aux.Anterior

				// Si el nodo eliminado es el inicio, actualizamos el puntero de inicio
				if l.Inicio == aux {
					l.Inicio = aux.Siguiente
				}
				l.Longitud--
				return existe, true

			}
		}
		aux = aux.Siguiente
		if aux == l.Inicio {
			break
		}
	}

	return existe, false
}

func (l *ListaCircularDoble) Reporte() {
	if l.Longitud == 0 {
		fmt.Println("No hay mas Tutores para graficar")
		return
	}
	nombreArchivo := "./reportes/listadoblecircular.dot"
	nombreImagen := "./reportes/listadoblecircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record fontname=Verdana];\n"
	aux := l.Inicio
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "Nombre: " + aux.Estudiante.Nombre + " \\n Carnet: " + strconv.Itoa(aux.Estudiante.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0 \n"
	texto += "nodo0 -> " + "nodo" + strconv.Itoa(contador) + "\n"
	texto += "} "
	pkg.CrearArchivo(nombreArchivo)
	pkg.EscribirArchivo(texto, nombreArchivo)
	pkg.Ejecutar(nombreImagen, nombreArchivo)
}
