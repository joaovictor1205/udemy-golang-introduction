package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number %d!", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {

	var num int

	fmt.Print("Value: ")
	fmt.Scan(&num)

	x, error := fatorial(num)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("Fat: %d\n", x)
	}

}
