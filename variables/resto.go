package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1400
	Fecha = time.Now()

	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado)
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)
}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
