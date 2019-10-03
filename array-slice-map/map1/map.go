package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123] = "teste1"
	aprovados[456] = "teste2"
	aprovados[789] = "teste3"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 123)

	fmt.Println(aprovados)
}
