package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"T": {
			"Teste":  123.123,
			"Teste2": 234.123,
		},
		"T2": {
			"Teste 3": 323.123,
			"Teste 4": 433.123,
		},
	}

	//delete(funcsPorLetra, "T2")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
