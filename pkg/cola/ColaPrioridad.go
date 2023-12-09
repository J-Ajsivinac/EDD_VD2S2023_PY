package cola

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cola struct {
	Inicio   *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &EstudianteTutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) EncolarPrioridad(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &EstudianteTutor{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if nota >= 90 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 75 && nota <= 89 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 65 && nota <= 74 {
		nuevoNodo.Prioridad = 3
	} else if nota >= 61 && nota <= 64 {
		nuevoNodo.Prioridad = 4
	} else {
		return
	}

	if c.Longitud == 0 {
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Inicio
		for aux.Siguiente != nil {
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad == nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad < nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else {
				aux = aux.Siguiente
			}
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.Inicio = c.Inicio.Siguiente
		c.Longitud--
	}
}

func (c *Cola) LeerArchivoTutores(ruta string) {
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
		nota, _ := strconv.Atoi(linea[3])
		c.EncolarPrioridad(valor, linea[1], linea[2], nota)
	}
}

func (c *Cola) Primero() {
	if c.Longitud == 0 {
		fmt.Println("No hay mas Tutores")
		return
	}
	fmt.Println("")
	fmt.Println(" ╔════════════════════════════════════════════════╗")
	fmt.Println(" ║                      Tutor                     ║")
	fmt.Println(" ╠════════════════════════════════════════════════╣")
	fmt.Println(" ║                                                ║")
	fmt.Printf(" ║ %-49s ║\n", "Actual: "+strconv.Itoa(c.Inicio.Tutor.Carnet))
	fmt.Printf(" ║ %-49s ║\n", "Nombre: "+c.Inicio.Tutor.Nombre)
	fmt.Printf(" ║ %-49s ║\n", "Curso: "+c.Inicio.Tutor.Curso)
	fmt.Printf(" ║ %-49s ║\n", "Nota: "+strconv.Itoa(c.Inicio.Tutor.Nota))
	fmt.Printf(" ║ %-49s ║\n", "Prioridad: "+strconv.Itoa(c.Inicio.Prioridad))
	if c.Inicio.Siguiente != nil {
		fmt.Printf(" ║ %-49s ║\n", "Siguiente: "+strconv.Itoa(c.Inicio.Siguiente.Tutor.Carnet))
	} else {
		fmt.Printf(" ║ %-49s ║\n", "Siguiente: No hay mas tutores por evaluar")
	}
	fmt.Println(" ║                                                ║")
	fmt.Println(" ╚════════════════════════════════════════════════╝")
	fmt.Println("")
	fmt.Println("1. Aceptar")
	fmt.Println("2. Rechazar")
	fmt.Println("Presione otra tecla para salir")
	fmt.Print("Eliga una opcion:")
}
