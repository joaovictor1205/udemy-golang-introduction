package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3} //array

	s := []int{1, 2, 3} //slice

	fmt.Println(a, s)

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := s2[:1]
	fmt.Println(s3)

}
