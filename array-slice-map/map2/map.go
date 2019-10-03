package main

import "fmt"

func main() {

	funcsEsalarios := map[string]float64{
		"Teste":   123,
		"Teste 2": 234,
		"Teste 3": 345,
	}

	fmt.Println(funcsEsalarios)

	funcsEsalarios["Teste 4"] = 456

	for nome, salario := range funcsEsalarios {
		fmt.Println(nome, salario)

	}

}
