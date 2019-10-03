package main

import (
	"fmt"
)

func main(){

	var a, b int

	fmt.Printf("Digite o primeiro valor:")
	fmt.Scan(&a)
	fmt.Printf("Digite o segundo valor:")
	fmt.Scan(&b)

	soma := somar(a,b)
	imprimir_soma(soma)
}