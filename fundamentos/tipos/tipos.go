package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println(1, 2, 1000)
	fmt.Println(reflect.TypeOf(1))

	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println(i1)

	char := 'a'
	fmt.Println("tipo char é", reflect.TypeOf(char))
}
