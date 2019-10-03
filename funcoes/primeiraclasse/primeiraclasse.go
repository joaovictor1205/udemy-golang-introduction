package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

var sub = func(a, b int) int {
	return a - b
}

func main() {

	var a, b int

	fmt.Println("Insira os valores: ")
	fmt.Scan(&a, &b)

	fmt.Printf("Soma = %d\n", soma(a, b))
	fmt.Printf("Subtracao = %d\n", sub(a, b))
}
