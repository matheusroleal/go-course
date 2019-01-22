package main

import "fmt"

// exemplo de função privada
func Add(x int, y int) (int, string) {
	return x + y, "Foi"
}

func main() {
	result, other := Add(10, 20)
	fmt.Println("O resultado de 10 + 20 é", result)
	fmt.Printf("O parametro extra é %s \n", other)
}
