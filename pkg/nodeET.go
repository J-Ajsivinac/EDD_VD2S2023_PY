package pkg

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

type NodoCola struct {
	Tutor     *EstudianteTutor
	Prioridad int
	Siguiente *NodoCola
}
