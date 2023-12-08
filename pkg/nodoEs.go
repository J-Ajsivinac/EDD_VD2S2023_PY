package pkg

type Estudiante struct {
	Carnet int
	Nombre string
}

type NodoDobleE struct {
	Estudiante *Estudiante
	Siguiente  *NodoDobleE
	Anterior   *NodoDobleE
}
