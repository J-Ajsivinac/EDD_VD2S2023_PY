package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

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
