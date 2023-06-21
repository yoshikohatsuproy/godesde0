package main

import (
	"fmt"

	"github.com/yoshio/godesde0/variables"
)

func main() {
	//variables.MostrarEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoTexto(1500)
	fmt.Println(estado, texto)
}
