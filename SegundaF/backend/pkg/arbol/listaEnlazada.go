package arbol

type NodoLista struct {
	Tutor     *NodoB
	Siguiente *NodoLista
}

type ListaSimple struct {
	Inicio   *NodoLista
	Longitud int
}

func (l *ListaSimple) Insertar(tutor *NodoB) {
	if l.Longitud == 0 {
		nuevo := &NodoLista{Tutor: tutor, Siguiente: nil}
		l.Inicio = nuevo
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoLista{Tutor: tutor, Siguiente: nil}
		l.Longitud++
	}
}

func (l *ListaSimple) ConverirUsuarios() []Usuario {
	var usuarios []Usuario
	aux := l.Inicio
	for aux != nil {
		usuarios = append(usuarios, *aux.Tutor.Usuario)
		aux = aux.Siguiente
	}
	return usuarios
}

func (l *ListaSimple) ConverirLibros() []Libro {
	var libros []Libro
	aux := l.Inicio
	for aux != nil {
		for _, libro := range aux.Tutor.Usuario.Libros {
			libros = append(libros, *libro)
		}
		aux = aux.Siguiente
	}
	return libros
}
