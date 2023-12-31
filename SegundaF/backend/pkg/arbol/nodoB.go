package arbol

type Libro struct {
	Nombre    string
	Contenido string
	Estado    string
}

type Usuario struct {
	Carnet        int
	Nombre        string
	Curso         string
	Password      string
	Libros        []*Libro
	Publicaciones []string
}

type NodoB struct {
	Usuario   *Usuario
	Siguiente *NodoB
	Anterior  *NodoB
	Izquierdo *RamaB
	Derecho   *RamaB
}
