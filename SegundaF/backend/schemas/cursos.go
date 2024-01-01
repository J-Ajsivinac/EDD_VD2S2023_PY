package schemas

type Libro struct {
	Nombre    string
	Contenido string
	Estado    string
}

type InfoBook struct {
	Codigo string  `json:"codigo"`
	Libros []Libro `json:"libros"`
}
