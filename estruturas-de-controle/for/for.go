package main

import (
	"fmt"
)

func main()  {

	fmt.Print("Pares: ")
	for i := 0; i <= 9; i++{
		if i % 2 == 0{
			fmt.Print(i)
		}

	}

	fmt.Println("")

	fmt.Print("Impares: ")
	for i := 0; i <= 9; i++{
		if i % 2 != 0{
			fmt.Print(i)
		}
	}
	
}