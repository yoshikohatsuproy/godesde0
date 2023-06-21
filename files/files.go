package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yoshio/godesde0/ejercicios"
)

var fileName string = "./files/archivos/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el Write String " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al abrir el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

}
