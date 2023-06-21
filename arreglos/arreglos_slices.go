package arreglos

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}

var tablaS []int = []int{2, 5, 4}
var arregloS [10]int = [10]int{2, 5, 4, 6, 12, 15, 678, 78, 9, 12}

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arregloS[3:]
	porcion1 := arregloS[:5]

	fmt.Println(porcion, porcion1)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
