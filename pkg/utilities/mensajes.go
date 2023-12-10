package utilities

import "fmt"

// colores de texto
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

// colores de fondo
const (
	BgReset  = "\033[49m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
)

func MensajeConsola(mensaje string, color string) {
	mensaje = " " + mensaje + " "
	typeColor := ""
	msjColor := ""
	switch color {
	case "rojo":
		typeColor = BgRed + " ✖ Error " + BgReset
		msjColor = ColorRed + mensaje + ColorReset
	case "verde":
		typeColor = BgGreen + " ✔ Operación Exitosa " + BgReset
		msjColor = ColorGreen + mensaje + ColorReset
	case "amarillo":
		typeColor = BgYellow + " ⚠ Advertencia " + BgReset
		msjColor = ColorYellow + mensaje + ColorReset
	case "azul":
		typeColor = BgBlue + " ♦ Procesando " + BgReset
		msjColor = ColorBlue + mensaje + ColorReset
	default:
		typeColor = "?"
		msjColor = mensaje
	}
	fmt.Println(typeColor + msjColor)
}
