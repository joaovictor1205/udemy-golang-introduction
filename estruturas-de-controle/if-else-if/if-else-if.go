package main

import (
	"fmt"
)

func nota(n float64) string {

	if n >= 9 && n < 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {

	var x float64

	fmt.Print("Nota: ")
	fmt.Scan(&x)

	fmt.Println("Conceito:", nota(x))

}
