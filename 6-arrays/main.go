package main

import "fmt"

// Desafio
func numerosSorte(){
	numSorte := [3]int{14, 28, 7}
	fmt.Println("Numeros da Sorte: ", numSorte)
}

func main() {
	var mandamentos [2]string
	mandamentos[0] = "Amor"
	mandamentos[1] = "Paz"
// // Declarar em uma linha o array com os valores
// 	mandamentos := [2]string{"Amor", "Paz"}
	fmt.Println("Mandamentos: ", mandamentos)
	numerosSorte()
}
