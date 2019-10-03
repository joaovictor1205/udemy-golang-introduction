package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//ascii
	fmt.Println("Teste" + string(97))

	//int for string
	fmt.Println("Teste" + strconv.Itoa(123))

	//string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)
	fmt.Println("Num is:", reflect.TypeOf(num))

	b, _ := strconv.ParseBool("True")
	if b {
		fmt.Println("True OK")
	}

}
