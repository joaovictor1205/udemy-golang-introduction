package main

import "fmt"

func obterValor(num int) int {

	defer fmt.Println("Fim!")

	if num > 5000 {
		fmt.Println("Valor alto")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return num
	}
}

func semDefer(num int) int {

	fmt.Println("Fim!")

	if num > 5000 {
		fmt.Println("Valor alto")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return num
	}
}

func main() {

	obterValor(5500)
	fmt.Println("")
	semDefer(2000)

}
