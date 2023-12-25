package schemas

type Login struct {
	Carnet     string `json:"carnet"`
	Contrasena string `json:"contrasena"`
	Estutor    bool   `json:"tutor"`
}
