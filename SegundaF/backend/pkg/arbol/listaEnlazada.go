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

func (l *ListaSimple) ConverirLibros() []Usuario {
	var usuarios []Usuario
	aux := l.Inicio
	for aux != nil {
		user := &Usuario{Nombre: aux.Tutor.Usuario.Nombre, Carnet: aux.Tutor.Usuario.Carnet, Password: aux.Tutor.Usuario.Password, Curso: aux.Tutor.Usuario.Curso}
		for _, libro := range aux.Tutor.Usuario.Libros {
			if libro.Estado == "Aceptado" {
				user.Libros = append(user.Libros, libro)
			}
		}
		usuarios = append(usuarios, *user)
		aux = aux.Siguiente
	}
	return usuarios
}
