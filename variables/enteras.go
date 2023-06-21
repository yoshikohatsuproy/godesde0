package variables

import "fmt"

func MostrarEnteros() {
	var intComun int

	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("int comun = ", intComun)
	fmt.Println("int 32 = ", intde32)
	fmt.Println("int 64 = ", intde64)
}
