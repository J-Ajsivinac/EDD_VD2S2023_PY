package listaD

import (
	"Proyecto/pkg"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type ListaDobleE struct {
	Primero  *NodoDobleE
	Ultimo   *NodoDobleE
	Longitud int
}

// funcion insertar para la lista doble
func (l *ListaDobleE) Insertar(estudiante *Estudiante) {
	nuevo := &NodoDobleE{Estudiante: estudiante}
	if l.Primero == nil {
		l.Primero = nuevo
		l.Ultimo = nuevo
	} else {
		l.Ultimo.Siguiente = nuevo
		nuevo.Anterior = l.Ultimo
		l.Ultimo = nuevo
	}
	l.Longitud++
}

// funcion para recorrer la lista doble
func (l *ListaDobleE) Recorrer() {
	aux := l.Primero
	for aux != nil {
		fmt.Println(aux.Estudiante.Carnet, "->", aux.Estudiante.Nombre, "->")
		aux = aux.Siguiente
	}
}

func (l *ListaDobleE) Buscar(carnet int) *NodoDobleE {
	aux := l.Primero
	for aux != nil {
		if aux.Estudiante.Carnet == carnet {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (l *ListaDobleE) Recorrer2() {
	aux := l.Ultimo
	for aux != nil {
		fmt.Println(aux.Estudiante.Carnet, "->", aux.Estudiante.Nombre, "->")
		aux = aux.Anterior
	}
}

func (l *ListaDobleE) LoginUser(contra int, carnet int) bool {
	aux := l.Primero
	for aux != nil {
		if aux.Estudiante.Carnet == contra && carnet == contra {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}

func (l *ListaDobleE) LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
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
		l.Insertar(&Estudiante{Carnet: valor, Nombre: linea[1]})
	}
}

func (l *ListaDobleE) Reporte() {
	if l.Longitud == 0 {
		fmt.Println("No hay mas Alumnos para graficar")
		return
	}
	nombreArchivo := "./reportes/listadoble.dot"
	nombreImagen := "./reportes/listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record fontname=Verdana];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Primero
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "Nombre: " + aux.Estudiante.Nombre + "\\n Carnet: " + strconv.Itoa(aux.Estudiante.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	pkg.CrearArchivo(nombreArchivo)
	pkg.EscribirArchivo(texto, nombreArchivo)
	pkg.Ejecutar(nombreImagen, nombreArchivo)
}
