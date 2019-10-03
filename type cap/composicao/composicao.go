package main

import "fmt"

type sportive interface {
	turboON()
}

type lux interface {
	drive()
}

type sportiveLux interface {
	sportive
	lux

	//...
}

type car struct{}

func (c car) turboON() {
	fmt.Println("ON")
}

func (c car) drive() {
	fmt.Println("DRIVING")
}

func main() {

	var car1 sportiveLux = car{}
	car1.turboON()
	car1.drive()

}
