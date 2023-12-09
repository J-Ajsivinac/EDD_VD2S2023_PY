package cola

type EstudianteTutor struct {
	Carnet int
	Nombre string
	Curso  string
	Nota   int
}

type NodoCola struct {
	Tutor     *EstudianteTutor
	Prioridad int
	Siguiente *NodoCola
}
