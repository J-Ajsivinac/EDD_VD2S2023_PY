package tabla

type Persona struct {
	Carnet   int
	Nombre   string
	Password string
	Cursos   [3]string
}

type NodoHash struct {
	Llave   int
	Persona *Persona
}
