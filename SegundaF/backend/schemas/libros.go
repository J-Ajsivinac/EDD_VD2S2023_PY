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

type ContenidoPub struct {
	Contenido string `json:"contenido"`
	Carnet    int    `json:"carnet"`
}

type BuscarCurso struct {
	Codigo []string `json:"codigo"`
}
