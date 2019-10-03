package main

import "fmt"

type sport interface{
	turnOn()
}

type car struct{
	model string
	on bool
	vel int
}

func (c *car) turnOn() {
	c.on = true
} 

func main(){
	car1 := car{
		model:	"test",
		on: 	 false,
		vel: 	 100,
	}
	car1.turnOn()

	var car2 sport = &car{
		model:	"test2",
		on: 	 false,
		vel: 	 80,
	}
	car2.turnOn()

	fmt.Println(car1, car2)
}