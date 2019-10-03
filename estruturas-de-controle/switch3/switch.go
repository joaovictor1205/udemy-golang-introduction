package main

import (
	"fmt"
	"time"
)

func typeOf(x interface{}) string{

	switch x.(type) {
	case int:
		return "Int"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Err!"
	}

}


func main(){
	
	fmt.Println("Tipo: ", typeOf(32))
	fmt.Println("Tipo: ", typeOf("teste"))
	fmt.Println("Tipo: ", typeOf(32.1))
	fmt.Println("Tipo: ", typeOf(func(){}))
	fmt.Println("Tipo: ", typeOf(time.Now()))

}