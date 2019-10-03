package main

import "fmt"

//without parameters and return
func f1() {
	fmt.Println("F1")
}

//with parameters without return
func f2(param1 string, param2 string) {
	fmt.Printf("F2 parameters: %s %s\n", param1, param2)
}

//without parameters with return
func f3() string {
	return "F3"
}

//parameters with same type with return
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

//without parameters, with 2 return
func f5() (string, string) {
	return "R1", "R2"
}

func main() {

	f1()
	f2("test1", "test2")
	fmt.Println(f3())
	fmt.Println(f4("test1", "test2"))

	r51, r52 := f5()
	fmt.Println(r51, r52)

	r551, _ := f5()
	fmt.Println(r551)
}
