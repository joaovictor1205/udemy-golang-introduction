package main

import (
	"fmt"
)

type value float64

func (v value) between(start, end float64) bool{
	return float64(v) >= start && float64(v) <= end
}

func note(v value) string{
	if v.between(9.0, 10.0){
		return "A"
	}else if v.between(7.0, 8.99) {
		return "B"
	} else if v.between(5.0, 7.99){
		return "C"
	}else if v.between(3.0, 4.99){
		return "D"
	} else{
		return "E"
	}
}

func main(){

	var x value

	fmt.Printf("Value:")
	fmt.Scan(&x)

	fmt.Println(note(x))

}