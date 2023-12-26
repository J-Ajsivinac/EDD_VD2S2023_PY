package arbol

type RamaB struct {
	Primero  *NodoB
	Hoja     bool
	Contador int
}

func (r *RamaB) Insertar(nuevoNodo *NodoB) {
	if r.Primero == nil { // 10
		r.Primero = nuevoNodo
		r.Contador++
	} else {
		if nuevoNodo.Usuario.Carnet < r.Primero.Usuario.Carnet { // 7 | 10
			nuevoNodo.Siguiente = r.Primero
			r.Primero.Izquierdo = nuevoNodo.Derecho //*********** Hice cambio -> nuevoNodo.Derecho = r.Primero.Izquierdo
			r.Primero.Anterior = nuevoNodo
			r.Primero = nuevoNodo
			r.Contador++
		} else if r.Primero.Siguiente != nil { // 7 | 9 | 10 -> 9
			if r.Primero.Siguiente.Usuario.Carnet > nuevoNodo.Usuario.Carnet {
				nuevoNodo.Siguiente = r.Primero.Siguiente
				nuevoNodo.Anterior = r.Primero
				r.Primero.Siguiente.Izquierdo = nuevoNodo.Derecho //*********** Hice cambio
				r.Primero.Derecho = nuevoNodo.Izquierdo
				r.Primero.Siguiente.Anterior = nuevoNodo
				r.Primero.Siguiente = nuevoNodo
				r.Contador++
			} else { // 7 | 10 -> 15
				aux := r.Primero.Siguiente
				nuevoNodo.Anterior = aux
				aux.Derecho = nuevoNodo.Izquierdo //*********** Hice cambio
				aux.Siguiente = nuevoNodo
				r.Contador++
			}
		} else if r.Primero.Siguiente == nil {
			nuevoNodo.Anterior = r.Primero
			r.Primero.Derecho = nuevoNodo.Izquierdo //*********** Hice cambio
			r.Primero.Siguiente = nuevoNodo
			r.Contador++
		}
	}
}