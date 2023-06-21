package main

import (
	//"runtime"
	"fmt"
	//"github.com/yoshio/godesde0/variables"
	"github.com/yoshio/godesde0/ejercicios"
	//"github.com/yoshio/godesde0/teclado"
	//"github.com/yoshio/godesde0/iteraciones"
)

func main() {
	//variables.MostrarEnteros()
	//variables.RestoVariables()

	/*
		estado, texto := variables.ConviertoTexto(1500)
		fmt.Println(estado, texto)

		os := runtime.GOOS

		if os == "Linux." || os == "OS X." {
			fmt.Println("Esto no es Windows")
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	/*
		numero, texto := ejercicios.Ejercicio1("ffff")
		//teclado.IngresoNumeros()

		iteraciones.Iterar()
		fmt.Println(numero, texto)*/

	fmt.Println(ejercicios.TablaMultiplicar())
}
