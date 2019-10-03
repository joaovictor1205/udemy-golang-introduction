package main

import (
	"fmt"
	"math"
)

func main(){

	var a, b int

	fmt.Printf("Insira dois valores: ")
	fmt.Scan(&a, &b)

	fmt.Println("Soma: ", a+b)
	fmt.Println("Subtracao: ", a-b)
	fmt.Println("Multiplicacao: ", a*b)
	fmt.Println("Divisao: ", a/b)
	fmt.Println("Mod: ", a%b)

	//bitwise
	fmt.Println("And:", a&b)
	fmt.Println("Or:", a|b)
	fmt.Println("Xor:", a^b)

	//math
	fmt.Println("Maior valor:", math.Max(float64(a), float64(b)))
	fmt.Println("Menor valor:", math.Min(float64(a), float64(b)))
	fmt.Println("Exp:", math.Pow(float64(a), float64(b)))

}