package main

import (
	"fmt"
)

func main(){

	i := 1
	var p *int = nil
	p = &i


	fmt.Println(p)
	*p++
	fmt.Println(p, *p, i, &i)

}