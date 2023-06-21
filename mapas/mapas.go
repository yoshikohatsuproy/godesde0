package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	fmt.Println(paises)

	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   30,
		"Real Madrid": 25,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)
}
