package pkg

import "fmt"

type ListaCircularDoble struct {
	Inicio   *NodoCircularE
	fin      *NodoCircularE
	Longitud int
}

func (l *ListaCircularDoble) InsertarOrdenado(nuevoE *EstudianteTutor) {
	nuevoNodo := &NodoCircularE{Estudiante: nuevoE}
	if l.Inicio == nil {
		l.Inicio = nuevoNodo
		l.fin = nuevoNodo
		l.Inicio.Siguiente = l.Inicio
		l.Inicio.Anterior = l.Inicio
	} else {
		if nuevoNodo.Estudiante.Carnet < l.Inicio.Estudiante.Carnet {
			nuevoNodo.Siguiente = l.Inicio
			nuevoNodo.Anterior = l.fin
			l.Inicio.Anterior = nuevoNodo
			l.Inicio = nuevoNodo
			l.fin.Siguiente = l.Inicio
		} else if nuevoNodo.Estudiante.Carnet > l.fin.Estudiante.Carnet {
			nuevoNodo.Siguiente = l.Inicio
			nuevoNodo.Anterior = l.fin
			l.fin.Siguiente = nuevoNodo
			l.fin = nuevoNodo
			l.Inicio.Anterior = l.fin
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
