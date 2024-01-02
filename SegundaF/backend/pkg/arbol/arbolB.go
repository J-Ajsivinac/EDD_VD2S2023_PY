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

func (a *ArbolB) insertar_rama(nodo *NodoB, rama *RamaB) *NodoB { // 20,
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
			if nodo.Usuario.Curso == temp.Usuario.Curso {
				return nil
			} else if nodo.Usuario.Curso < temp.Usuario.Curso {
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
	tutor := &Usuario{Carnet: 0, Nombre: "", Curso: "", Password: ""}
	val := &NodoB{Usuario: tutor}
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
				rizquierda.Hoja = false
			}
			rizquierda.Insertar(temp)
		} else if contador == 2 {
			val.Usuario = aux.Usuario
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
	contrasenaE := encriptarPassword(password)
	tutor := &Usuario{Carnet: carnet, Nombre: nombre, Curso: curso, Password: contrasenaE}
	fmt.Println(carnet, nombre, curso, password)
	nuevoNodo := &NodoB{Usuario: tutor}
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

func (a *ArbolB) Graficar() string {
	if a.Raiz == nil {
		return ""
	}

	cadena := ""
	nombre_archivo := "./reportes/tutores.dot"
	nombre_imagen := "./reportes/tutores.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record, color=white, fontcolor=white ];\n"
		cadena += "edge[color=white];\n"
		cadena += "bgcolor=\"#1e1f23\";\n"
		cadena += a.grafo(a.Raiz.Primero)
		cadena += a.conexionRamas(a.Raiz.Primero)
		cadena += "}"
	}
	pkg.CrearArchivo(nombre_archivo)
	pkg.EscribirArchivo(cadena, nombre_archivo)
	pkg.Ejecutar(nombre_imagen, nombre_archivo)

	return nombre_imagen

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

func (a *ArbolB) Buscar(numero int, listaSimple *ListaSimple) {
	if a.Raiz == nil {
		return
	}
	a.buscarArbol(a.Raiz.Primero, numero, listaSimple)
	if listaSimple.Longitud > 0 {
		// fmt.Println("Se encontro el elemento", listaSimple.Longitud)
		fmt.Println(numero, listaSimple.Inicio.Tutor.Usuario.Carnet, listaSimple.Inicio.Tutor.Usuario.Password)
	} else {
		fmt.Println("No se encontro")
	}
}

func (a *ArbolB) buscarArbol(raiz *NodoB, numero int, listaSimple *ListaSimple) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.buscarArbol(aux.Izquierdo.Primero, numero, listaSimple)
			}
			if aux.Usuario != nil && aux.Usuario.Carnet == numero {
				fmt.Println(aux.Usuario, "----")
				listaSimple.Insertar(aux)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.buscarArbol(aux.Derecho.Primero, numero, listaSimple)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) ObtenerLibros(raiz *NodoB, listaSimple *ListaSimple) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.ObtenerLibros(aux.Izquierdo.Primero, listaSimple)
			}
			//listaSimple.Insertar(aux)
			if aux.Usuario.Libros != nil {
				listaSimple.Insertar(aux)
			}

			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.ObtenerLibros(aux.Derecho.Primero, listaSimple)
				}
			}
			aux = aux.Siguiente
		}
	}
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
			if aux.Usuario != nil && aux.Usuario.Carnet == carnet {
				fmt.Println("Registre el libro", carnet, nombre)
				aux.Usuario.Libros = append(aux.Usuario.Libros, &Libro{Nombre: nombre, Contenido: contenido, Estado: "Pendiente"})
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

func (a *ArbolB) CambiarEstadoLibro(raiz *NodoB, carnet int, nombre string, estado string) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.CambiarEstadoLibro(aux.Izquierdo.Primero, carnet, nombre, estado)
			}
			if aux.Usuario.Carnet == carnet {
				for _, libro := range aux.Usuario.Libros {
					if libro.Nombre == nombre {
						libro.Estado = estado
						return
					}
				}
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.CambiarEstadoLibro(aux.Derecho.Primero, carnet, nombre, estado)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) LlamarBuscarLibros(numero string, listaSimple *ListaSimple) {
	if a.Raiz == nil {
		return
	}
	a.BuscarLibrosC(a.Raiz.Primero, numero, listaSimple)
	if listaSimple.Longitud > 0 {
		fmt.Println("Se encontro el elemento", listaSimple.Longitud)
	} else {
		fmt.Println("No se encontro")
	}
}

func (a *ArbolB) BuscarLibrosC(raiz *NodoB, codigo string, listaSimple *ListaSimple) {

	if raiz != nil {
		aux := raiz
		for aux != nil {

			if aux.Izquierdo != nil {
				a.BuscarLibrosC(aux.Izquierdo.Primero, codigo, listaSimple)
			}
			if aux.Usuario != nil && aux.Usuario.Curso == codigo {
				listaSimple.Insertar(aux)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.BuscarLibrosC(aux.Derecho.Primero, codigo, listaSimple)
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
				aux.Usuario.Publicaciones = append(aux.Usuario.Publicaciones, contenido)
				// fmt.Println("Registre el libro")
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
