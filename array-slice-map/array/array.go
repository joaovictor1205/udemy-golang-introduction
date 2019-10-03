package main

import (
	"fmt"
)

func main() {

	var notas [3]float64
	var total float64

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Insira o %dº valor: ", i)
		fmt.Scan(&notas[i])
		total += notas[i]

	}

	fmt.Println(notas)

	media := total / float64(len(notas))
	fmt.Println(media)

}
