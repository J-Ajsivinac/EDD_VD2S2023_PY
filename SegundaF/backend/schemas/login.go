package schemas

type Login struct {
	Carnet     string `json:"carnet"`
	Contrasena string `json:"contrasena"`
	Estutor    bool   `json:"tutor"`
}

type UserData struct {
	Indice   int    `json:"indice"`
	Carnet   int    `json:"carnet"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
	Cursos   string `json:"cursos"`
}
