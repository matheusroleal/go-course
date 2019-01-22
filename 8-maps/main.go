package main

import "fmt"

func main() {
	people := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	people["junior"] = 10

	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %d\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("Posição não alocada: %d\n", people["wilson"])

	// Desafio
	fmt.Println("#### Desafio de Map ####")
	desafio := map[int]float64{
		1: 3.14,
		2: 6.28,
	}

	for cod, val := range desafio {
		fmt.Printf("Chave %d do mapa valor: %.2f\n", cod, val)
	}

}
