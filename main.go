package main

import (
	//"fmt"
	//"runtime"
	//"github.com/yoshio/godesde0/variables"
	"fmt"

	"github.com/yoshio/godesde0/ejercicios"
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

	numero, texto := ejercicios.Ejercicio1("1000")

	fmt.Println(numero, texto)
}
