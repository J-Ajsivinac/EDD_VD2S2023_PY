package cola

import (
	"Proyecto/pkg/utilities"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cola struct {
	Inicio   *NodoCola
	Longitud int
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
		if aux.Prioridad > nuevoNodo.Prioridad {
			nuevoNodo.Siguiente = aux
			c.Inicio = nuevoNodo
			c.Longitud++
			return
		}

		for aux.Siguiente != nil && aux.Siguiente.Prioridad <= nuevoNodo.Prioridad {
			aux = aux.Siguiente
		}

		nuevoNodo.Siguiente = aux.Siguiente
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		utilities.MensajeConsola("No hay tutores en la cola", "rojo")
	} else {
		c.Inicio = c.Inicio.Siguiente
		c.Longitud--
	}
}

func (c *Cola) LeerArchivoTutores(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		utilities.MensajeConsola("No pude abrir el archivo", "rojo")
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
			utilities.MensajeConsola("No se puede leer la linea del csv", "rojo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		nota, _ := strconv.Atoi(linea[3])
		c.EncolarPrioridad(valor, strings.TrimSpace(linea[1]), strings.TrimSpace("0"+linea[2]), nota)
	}
	utilities.MensajeConsola("Carga de Estudiantes Tutores exitosa", "verde")
}

func (c *Cola) ImprimirCola() {
	aux := c.Inicio
	for aux != nil {
		fmt.Println(aux.Tutor.Carnet, aux.Tutor.Nombre, aux.Tutor.Curso, aux.Tutor.Nota, aux.Prioridad)
		aux = aux.Siguiente
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
	fmt.Printf(" ║ %-46s ║\n", "Actual: "+strconv.Itoa(c.Inicio.Tutor.Carnet))
	fmt.Printf(" ║ %-46s ║\n", "Nombre: "+c.Inicio.Tutor.Nombre)
	fmt.Printf(" ║ %-46s ║\n", "Curso: "+c.Inicio.Tutor.Curso)
	fmt.Printf(" ║ %-46s ║\n", "Nota: "+strconv.Itoa(c.Inicio.Tutor.Nota))
	fmt.Printf(" ║ %-46s ║\n", "Prioridad: "+strconv.Itoa(c.Inicio.Prioridad))
	if c.Inicio.Siguiente != nil {
		fmt.Printf(" ║ %-46s ║\n", "Siguiente: "+strconv.Itoa(c.Inicio.Siguiente.Tutor.Carnet))
	} else {
		fmt.Printf(" ║ %-46s ║\n", "Siguiente: No hay mas tutores por evaluar")
	}
	fmt.Println(" ║                                                ║")
	fmt.Println(" ╚════════════════════════════════════════════════╝")
	fmt.Println("")
	fmt.Println("1. Aceptar")
	fmt.Println("2. Rechazar")
	fmt.Println("Presione otro numero para salir")
	fmt.Print(" -> Eliga una opcion: ")
}
