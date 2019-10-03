package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado:", nota)
	} else {
		fmt.Println("Reprovado:", nota)
	}
}

func main() {

	var x float64

	fmt.Print("Insira a nota: ")
	fmt.Scan(&x)

	imprimirResultado(x)

}
