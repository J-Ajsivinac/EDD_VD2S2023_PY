package pkg

type EstudianteTutor struct {
	Carnet int
	Nombre string
	Curso  string
	nota   int
}

type NodoCircularE struct {
	Estudiante *EstudianteTutor
	Siguiente  *NodoCircularE
	Anterior   *NodoCircularE
}
