package pkg

import (
	"Proyecto/pkg/utilities"
	"os"
	"os/exec"
)

func CrearArchivo(nombre_archivo string) {
	var _, err = os.Stat(nombre_archivo)

	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
	utilities.MensajeConsola("Archivo generado exitosamente", "verde")
}

func EscribirArchivo(contenido string, nombre_archivo string) {
	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	err = file.Sync()
	if err != nil {
		return
	}
	utilities.MensajeConsola("Archivo guardado correctamente", "verde")
}

func Ejecutar(nombre_imagen string, archivo string) {
	path, _ := exec.LookPath("dot")
	dpi := "900"
	cmd, _ := exec.Command(path, "-Gscale=4 -Gantialias=true", "-Gdpi"+dpi, "-Tjpg", archivo).Output()
	mode := 0777
	_ = os.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
	//dot lista.dot -Tjpg lista.jpg -o
}
