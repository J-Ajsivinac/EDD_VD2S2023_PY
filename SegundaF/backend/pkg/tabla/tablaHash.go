package tabla

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calculoIndice(carnet int) int {
	var numeros []int //2017 -> [2 0 1 7]
	// Convertir el numero en array
	for {
		if carnet > 0 {
			digito := carnet % 10
			numeros = append([]int{digito}, numeros...)
			carnet = carnet / 10
		} else {
			break
		}
	}

	//Convertir array de numeros en Codigo ascii
	var numeros_ascii []rune
	for _, numero := range numeros {
		valor := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor)
	}

	//Sumar cada numero ascii
	final := 0
	for _, numero_ascii := range numeros_ascii {
		final += int(numero_ascii)
	}

	indice := final % t.Capacidad
	return indice
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.7
	if t.Utilizacion > int(auxCap) {
		auxAnterior := t.Capacidad
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar(auxAnterior)
	}

}

func (t *TablaHash) nuevaCapacidad() int {
	contador := 0
	a, b := 0, 1
	for contador < 70 {
		contador++
		if a > t.Capacidad {
			return a
		}
		a, b = b, a+b
	}
	return a
}

func (t *TablaHash) reInsertar(capacidadAnterior int) {
	auxTabla := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < capacidadAnterior; i++ {
		if usuario, existe := auxTabla[i]; existe {
			t.Insertar(usuario.Persona.Carnet, usuario.Persona.Nombre, usuario.Persona.Password, usuario.Persona.Cursos)
		}
	}
}

func (t *TablaHash) reCalculoIndice(carnet int, contador int) int {
	nuevoIndice := t.calculoIndice(carnet) + (contador * contador) // 2 + 36 = 38 -> capacidad actual 7
	return t.nuevoIndice(nuevoIndice)
}

func (t *TablaHash) nuevoIndice(nuevoIndice int) int {
	nuevoPosicion := 0
	if nuevoIndice < t.Capacidad {
		nuevoPosicion = nuevoIndice
	} else {
		nuevoPosicion = nuevoIndice - t.Capacidad
		nuevoPosicion = t.nuevoIndice(nuevoPosicion)
	}
	return nuevoPosicion
}

func (t *TablaHash) Insertar(carnet int, nombre string, password string, cursos [3]string) {
	indice := t.calculoIndice(carnet)
	nuevoNodo := &NodoHash{Llave: indice, Persona: &Persona{Carnet: carnet, Nombre: nombre, Password: password, Cursos: cursos}}
	if indice < t.Capacidad {
		// 5
		if _, existe := t.Tabla[indice]; !existe {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if _, existe := t.Tabla[indice]; existe {
					contador++
					indice = t.reCalculoIndice(carnet, contador)
				} else {
					nuevoNodo.Llave = indice
					t.Tabla[indice] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					break
				}
			}
		}
	}
}

func encriptarPassword(password string) string {
	texto := password
	hash := sha256.Sum256([]byte(texto))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}

func (t *TablaHash) BuscarUsuario(carnet int, password string) (*Persona, bool) {
	indice := t.calculoIndice(carnet)
	inputHash := sha256.Sum256([]byte(password))
	inputHashString := fmt.Sprintf("%x", inputHash)
	if usuario, existe := t.Tabla[indice]; existe {
		if usuario.Persona.Carnet == carnet && inputHashString == usuario.Persona.Password {
			return usuario.Persona, true
		} else {
			// Realizar búsqueda con sondaje cuadrático en caso de colisión
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if usuario, existe := t.Tabla[indice]; existe {
					if usuario.Persona.Carnet == carnet && inputHashString == usuario.Persona.Password {
						return usuario.Persona, true
					} else {
						contador++
						indice = t.reCalculoIndice(carnet, contador)
					}
				} else {
					break
				}
			}
		}
	}
	return nil, false
}

func (t *TablaHash) LeerCSVFromReader(reader io.Reader) {
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
		password := encriptarPassword(linea[2])
		var cursos [3]string
		cursos[0] = linea[3]
		cursos[1] = linea[4]
		cursos[2] = linea[5]
		valor, _ := strconv.Atoi(linea[0])
		t.Insertar(valor, linea[1], password, cursos)
	}
}
