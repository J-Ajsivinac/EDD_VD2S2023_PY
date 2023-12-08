package pkg

import (
	"fmt"
	"strconv"
)

type ListaDobleE struct {
	Primero  *NodoDobleE
	Ultimo   *NodoDobleE
	Longitud int
}

// funcion insertar para la lista doble
func (l *ListaDobleE) Insertar(estudiante *Estudiante) {
	nuevo := &NodoDobleE{Estudiante: estudiante}
	if l.Primero == nil {
		l.Primero = nuevo
		l.Ultimo = nuevo
	} else {
		l.Ultimo.Siguiente = nuevo
		nuevo.Anterior = l.Ultimo
		l.Ultimo = nuevo
	}
	l.Longitud++
}

// funcion para recorrer la lista doble
func (l *ListaDobleE) Recorrer() {
	aux := l.Primero
	for aux != nil {
		fmt.Println(aux.Estudiante.Carnet, "->", aux.Estudiante.Nombre, "->")
		aux = aux.Siguiente
	}
}

func (l *ListaDobleE) Buscar(carnet int) *NodoDobleE {
	aux := l.Primero
	for aux != nil {
		if aux.Estudiante.Carnet == carnet {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (l *ListaDobleE) Recorrer2() {
	aux := l.Ultimo
	for aux != nil {
		fmt.Println(aux.Estudiante.Carnet, "->", aux.Estudiante.Nombre, "->")
		aux = aux.Anterior
	}
}

func (l *ListaDobleE) LoginUser(contra int) bool {
	aux := l.Primero
	for aux != nil {
		if aux.Estudiante.Carnet == contra {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

func (l *ListaDobleE) Reporte() {
	if l.Longitud == 0 {
		fmt.Println("No hay mas Alumnos para graficar")
		return
	}
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Primero
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
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
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
