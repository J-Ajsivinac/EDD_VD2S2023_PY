package grafoA

type NodoListaAdyacencia struct {
	Siguiente *NodoListaAdyacencia
	Abajo     *NodoListaAdyacencia
	Valor     string
}
