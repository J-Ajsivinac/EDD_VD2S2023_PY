package schemas

type AceptarLibros struct {
	Carnet int    `json:"carnet"`
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}

type AgregarLibro struct {
	Carnet    int    `json:"carnet"`
	Contenido string `json:"contenido"`
	Nombre    string `json:"nombre"`
}
