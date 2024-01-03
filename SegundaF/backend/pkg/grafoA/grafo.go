package grafoA

import (
	"backend/pkg"
	"encoding/json"
	"io"
)

type Grafo struct {
	Principal *NodoListaAdyacencia
}

type Cursos struct {
	Codigo        string   `json:"Codigo"`
	PostRequisito []string `json:"Post"`
}

type DatosCurso struct {
	Curso []Cursos `json:"Cursos"`
}

func (g *Grafo) insertarColumna(curso string, post string) {
	nuevoNodo := &NodoListaAdyacencia{Valor: post}
	if g.Principal != nil && curso == g.Principal.Valor {
		g.insertarFila(post)
		aux := g.Principal
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
	} else {
		g.insertarFila(curso)
		aux := g.Principal
		for aux != nil {
			if aux.Valor == curso {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			for aux.Siguiente != nil {
				aux = aux.Siguiente
			}
			aux.Siguiente = nuevoNodo
		}
	}
}

func (g *Grafo) insertarFila(curso string) {
	nuevoNodo := &NodoListaAdyacencia{Valor: curso}
	if g.Principal == nil {
		g.Principal = nuevoNodo
	} else {
		aux := g.Principal
		for aux.Abajo != nil {
			if aux.Valor == curso {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

func (g *Grafo) InsertarValores(curso string, post string) {
	if g.Principal == nil {
		//insertar Fila
		g.insertarFila(curso)
		//insertar Columna
		g.insertarColumna(curso, post)
	} else {
		g.insertarColumna(curso, post)
	}
}

func (g *Grafo) Graficar() string {

	if g.Principal == nil {
		return ""
	}

	cadena := ""
	nombre_archivo := "./reportes/cursos.dot"
	nombre_imagen := "./reportes/cursos.jpg"
	if g.Principal != nil {
		cadena += "digraph grafoDirigido{ \n rankdir=LR; \n node [shape=box, color=white, fontcolor=white]; layout=neato; \n nodo" + g.Principal.Valor + "[label=\"" + g.Principal.Valor + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += "bgcolor=\"#1e1f23\";\n"
		cadena += "edge[color=white];\n"
		cadena += g.retornarValoresMatriz()
		cadena += "\n}"
	}
	pkg.CrearArchivo(nombre_archivo)
	pkg.EscribirArchivo(cadena, nombre_archivo)
	pkg.Ejecutar(nombre_imagen, nombre_archivo)

	return nombre_imagen

}

func (g *Grafo) retornarValoresMatriz() string {
	cadena := ""

	aux := g.Principal.Abajo
	aux1 := aux

	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux1.Valor + "[label=\"" + aux1.Valor + "\" ]; \n"
			aux1 = aux1.Siguiente
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
	}

	aux = g.Principal
	aux1 = aux.Siguiente

	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux.Valor + " -> "
			cadena += "nodo" + aux1.Valor + "[len=1.00]; \n"
			aux1 = aux1.Siguiente
		}
		if aux.Abajo != nil {
			aux = aux.Abajo
			aux1 = aux.Siguiente
		} else {
			aux = aux.Abajo
		}
	}

	return cadena
}

func (matriz1 *Grafo) Lectura(reader io.Reader) (bool, string) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return false, "Error al leer el json: " + err.Error()
	}

	var datos DatosCurso
	err = json.Unmarshal(data, &datos)
	if err != nil {
		return false, "Error al asignar el json: " + err.Error()
	}
	for _, curso := range datos.Curso {
		if len(curso.PostRequisito) > 0 {
			for j := 0; j < len(curso.PostRequisito); j++ {
				matriz1.InsertarValores(curso.Codigo, curso.PostRequisito[j])
			}
		} else {
			matriz1.InsertarValores("ECYS", curso.Codigo)
		}
	}
	return true, "Se ha cargado el archivo correctamente"
}

func (g *Grafo) Buscar(curso string) bool {
	aux := g.Principal
	for aux != nil {
		if aux.Valor == curso {
			return true
		}
		aux = aux.Abajo
	}
	return false
}
