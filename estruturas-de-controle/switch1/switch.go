package main

import (
	"fmt"
)

func nota(n int) string{
	var nota = int(n)

	switch nota{

	case 10:
		fallthrough

	case 9:
		return "A"

	case 8,7:
		return "B"

	case 6,5:
		return "C"
	
	case 4,3:
		return "D"

	case 2,1,0:
		return "E"

	default:
		return "Nota Invalida!"
	}
	
}

func main()  {

	var valor int

	fmt.Print("Insira a nota: ")
	fmt.Scan(&valor)

	fmt.Println("Conceito: ", nota(valor))
	
}