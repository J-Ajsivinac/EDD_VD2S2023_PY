package tabla

type Persona struct {
	Carnet   int
	Nombre   string
	Password [32]byte
	Cursos   [3]string
}

type NodoHash struct {
	Llave   int
	Persona *Persona
}
