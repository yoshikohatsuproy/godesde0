package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	var numero int
	var err error
	var texto string
	for {

		fmt.Println("Ingresar el número para calcular su tabla")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", i, numero, numero*i)
		//fmt.Printf("%d x %d = %d \n", i, numero, numero*i)
	}
	return texto
}
