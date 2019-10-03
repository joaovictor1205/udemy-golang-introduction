package main

import(
	"fmt"
)

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados{
		fmt.Printf("%d) %s\n", i, aprovado)
	}

}

func main(){

	aprovados := []string{"Test1", "Test2", "Test3", "Test4", "Test5"}
	imprimirAprovados(aprovados...)

}