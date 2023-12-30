package arbol

import (
	"backend/pkg"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}

func (a *ArbolB) insertar_rama(nodo *NodoB, rama *RamaB) *NodoB {
	if rama.Hoja {
		rama.Insertar(nodo)
		if rama.Contador == a.Orden {
			return a.dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.Primero
		for temp != nil {
			if nodo.Usuario.Carnet == temp.Usuario.Carnet {
				return nil
			} else if nodo.Usuario.Carnet < temp.Usuario.Carnet {
				obj := a.insertar_rama(nodo, temp.Izquierdo)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			} else if temp.Siguiente == nil {
				obj := a.insertar_rama(nodo, temp.Derecho)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Siguiente
		}
	}
	return nil
}

func (a *ArbolB) dividir(rama *RamaB) *NodoB {
	val := &NodoB{Usuario: &Usuario{Carnet: -99999, Nombre: "", Curso: "", Password: ""}}
	aux := rama.Primero
	rderecha := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	rizquierda := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	contador := 0
	for aux != nil {
		contador++
		if contador < 2 {
			temp := &NodoB{Usuario: aux.Usuario}
			temp.Izquierdo = aux.Izquierdo
			if contador == 1 {
				temp.Derecho = aux.Siguiente.Izquierdo
			}
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rizquierda.Hoja = false //Noo es nodo hoja, sino raiz
			}
			rizquierda.Insertar(temp)
		} else if contador == 2 {
			val.Usuario.Carnet = aux.Usuario.Carnet
		} else {
			temp := &NodoB{Usuario: aux.Usuario}
			temp.Izquierdo = aux.Izquierdo
			temp.Derecho = aux.Derecho
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rderecha.Hoja = false
			}
			rderecha.Insertar(temp)
		}
		aux = aux.Siguiente
	}
	nuevo := &NodoB{Usuario: val.Usuario}
	nuevo.Derecho = rderecha
	nuevo.Izquierdo = rizquierda

	return nuevo
}

func (a *ArbolB) Insertar(carnet int, nombre string, curso string, password string) { //15
	nuevoNodo := &NodoB{Usuario: &Usuario{Carnet: carnet, Nombre: nombre, Curso: curso, Password: encriptarPassword(password)}}
	if a.Raiz == nil {
		a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
		a.Raiz.Insertar(nuevoNodo)
	} else {
		obj := a.insertar_rama(nuevoNodo, a.Raiz)
		if obj != nil {
			a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
			a.Raiz.Insertar(obj)
			a.Raiz.Hoja = false
		}
	}
}

func (a *ArbolB) Graficar() {
	cadena := ""
	nombre_archivo := "./reportes/tutores.jpg"
	nombre_imagen := "./reportes/tutores.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record]\n"
		cadena += a.grafo(a.Raiz.Primero)
		cadena += a.conexionRamas(a.Raiz.Primero)
		cadena += "}"
	}
	pkg.CrearArchivo(nombre_archivo)
	pkg.EscribirArchivo(cadena, nombre_archivo)
	pkg.Ejecutar(nombre_imagen, nombre_archivo)

}

func (a *ArbolB) grafo(rama *NodoB) string {
	dot := ""
	if rama != nil {
		dot += a.grafoRamas(rama)
		aux := rama
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += a.grafo(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += a.grafo(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

func (a *ArbolB) grafoRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		dot = dot + "R" + strconv.Itoa(rama.Usuario.Carnet) + "[label=\""
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot = dot + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Siguiente != nil {
				dot = dot + strconv.Itoa(aux.Usuario.Carnet) + "|"
			} else {
				dot = dot + strconv.Itoa(aux.Usuario.Carnet)
				if aux.Derecho != nil {
					dot = dot + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Siguiente
		}
		dot = dot + "\"];\n"
	}
	return dot
}

func (a *ArbolB) conexionRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		actual := "R" + strconv.Itoa(rama.Usuario.Carnet)
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Izquierdo.Primero.Usuario.Carnet) + ";\n"
				r++
				dot += a.conexionRamas(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Derecho.Primero.Usuario.Carnet) + ";\n"
					r++
					dot += a.conexionRamas(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

func (a *ArbolB) Buscar(numero int) *Usuario {
	buscarElemento := a.buscarArbol(a.Raiz.Primero, numero)
	if buscarElemento != nil {
		return buscarElemento
	} else {
		return buscarElemento
	}
}

func (a *ArbolB) buscarArbol(raiz *NodoB, numero int) *Usuario {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				usuarioEncontrado := a.buscarArbol(aux.Izquierdo.Primero, numero)
				if usuarioEncontrado != nil {
					return usuarioEncontrado
				}
			}
			if aux.Usuario != nil && aux.Usuario.Carnet == numero {
				return aux.Usuario
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					usuarioEncontrado := a.buscarArbol(aux.Derecho.Primero, numero)
					if usuarioEncontrado != nil {
						return usuarioEncontrado
					}
				}
			}
			aux = aux.Siguiente
		}
	}
	return nil
}

func (t *ArbolB) LeerCSV(reader io.Reader) {
	lectura := csv.NewReader(reader)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}

		valor, _ := strconv.Atoi(linea[0])
		t.Insertar(valor, linea[1], linea[2], linea[3])
	}
}

func (a *ArbolB) GuardarLibro(raiz *NodoB, nombre string, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.GuardarLibro(aux.Izquierdo.Primero, nombre, contenido, carnet)
			}
			if aux.Usuario.Carnet == carnet {
				raiz.Usuario.Libros = append(raiz.Usuario.Libros, &Libro{Nombre: nombre, Contenido: contenido, Estado: 1})
				fmt.Println("Registre el libro")
				return
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.GuardarLibro(aux.Derecho.Primero, nombre, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) GuardarPublicacion(raiz *NodoB, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.GuardarPublicacion(aux.Izquierdo.Primero, contenido, carnet)
			}
			if aux.Usuario.Carnet == carnet {
				raiz.Usuario.Publicaciones = append(raiz.Usuario.Publicaciones, contenido)
				fmt.Println("Registre el libro")
				return
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.GuardarPublicacion(aux.Derecho.Primero, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func encriptarPassword(password string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(password))
	hexaString = hex.EncodeToString(h.Sum(nil))
	return hexaString
}
