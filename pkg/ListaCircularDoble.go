package pkg

import (
	"fmt"
	"strconv"
)

type ListaCircularDoble struct {
	Inicio   *NodoCircularE
	Fin      *NodoCircularE
	Longitud int
}

func (l *ListaCircularDoble) InsertarOrdenado(carnet int, nombre string, curso string, nota int) {
	nuevoAlumno := &EstudianteTutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCircularE{Estudiante: nuevoAlumno}
	if l.Inicio == nil {
		l.Inicio = nuevoNodo
		l.Fin = nuevoNodo
		l.Inicio.Siguiente = l.Inicio
		l.Inicio.Anterior = l.Inicio
	} else {
		if nuevoNodo.Estudiante.Carnet < l.Inicio.Estudiante.Carnet {
			nuevoNodo.Siguiente = l.Inicio
			nuevoNodo.Anterior = l.Fin
			l.Inicio.Anterior = nuevoNodo
			l.Inicio = nuevoNodo
			l.Fin.Siguiente = l.Inicio
		} else if nuevoNodo.Estudiante.Carnet > l.Fin.Estudiante.Carnet {
			nuevoNodo.Siguiente = l.Inicio
			nuevoNodo.Anterior = l.Fin
			l.Fin.Siguiente = nuevoNodo
			l.Fin = nuevoNodo
			l.Inicio.Anterior = l.Fin
		} else {
			aux := l.Inicio
			for aux.Siguiente != l.Inicio {
				if nuevoNodo.Estudiante.Carnet < aux.Siguiente.Estudiante.Carnet {
					break
				}
				aux = aux.Siguiente
			}
			nuevoNodo.Siguiente = aux.Siguiente
			nuevoNodo.Anterior = aux
			aux.Siguiente.Anterior = nuevoNodo
			aux.Siguiente = nuevoNodo
		}
	}
	l.Longitud++
}

func (l *ListaCircularDoble) Recorrer() {
	aux := l.Inicio
	for aux.Siguiente != l.Inicio {
		fmt.Println(aux.Estudiante.Carnet, "->", aux.Estudiante.Nombre, "->")
		aux = aux.Siguiente
	}
}

func (l *ListaCircularDoble) Reporte() {
	if l.Longitud == 0 {
		fmt.Println("No hay mas Tutores para graficar")
		return
	}
	nombreArchivo := "./listadoblecircular.dot"
	nombreImagen := "./listadoblecircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + strconv.Itoa(aux.Estudiante.Carnet) + "\"];\n"
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
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
