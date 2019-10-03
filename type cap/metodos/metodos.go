package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

func (p *person) setFullName(fullName string) {
	part := strings.Split(fullName, " ")
	p.name = part[0]
	p.lastName = part[1]
}

func main() {
	p1 := person{
		name:     "Test",
		lastName: "One",
	}

	p2 := person{}

	fmt.Println(p1.getFullName())

	p2.setFullName("Test Two")
	fmt.Println(p2.getFullName())

}
