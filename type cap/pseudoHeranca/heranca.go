package main

import "fmt"

type car struct {
	name string
	vel  int
}

type ferrari struct {
	car
	on bool
}

func main() {

	f := ferrari{}
	f.name = "Ferrari"
	f.vel = 100
	f.on = true

	fmt.Println(f)
}
