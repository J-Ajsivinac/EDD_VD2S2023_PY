package pkg

import "fmt"

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
