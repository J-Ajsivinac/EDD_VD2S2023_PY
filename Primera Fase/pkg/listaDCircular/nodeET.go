package listaDCircular

type EstudianteTutor struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoCircularE struct {
	Estudiante *EstudianteTutor
	Siguiente  *NodoCircularE
	Anterior   *NodoCircularE
}
