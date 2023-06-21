package ejercicios

import "strconv"

func Ejercicio1(texto string) (int, string) {

	entero, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Hubo error" + err.Error()
	}

	if entero > 100 {
		return entero, "Es mayor a 100"
	} else {
		return entero, "Es menor a 100"
	}
}
