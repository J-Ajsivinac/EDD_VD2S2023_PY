package arbol

type Libro struct {
	Contenido string
	Estado    string
}

type Usuario struct {
	Carnet      int
	Nombre      string
	Curso       string
	Password    string
	Libros      []*Libro
	Publicacion []string
}

type NodoB struct {
	Usuario   *Usuario
	Siguiente *NodoB
	Anterior  *NodoB
	Izquierdo *RamaB
	Derecho   *RamaB
}
