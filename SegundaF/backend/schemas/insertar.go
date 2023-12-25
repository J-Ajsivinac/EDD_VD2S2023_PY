package schemas

type AgregarArbol struct {
	Carnet   int    `json:"carnet"`
	Nombre   string `json:"nombre"`
	Curso    string `json:"curso"`
	Password string `json:"password"`
}
