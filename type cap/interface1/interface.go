package main

import (

	"fmt"
)

type show interface{
	toString() string
}

type person struct{
	name 	 string
	lastName string
}

type product struct{
	name 	string
	price 	float64
}

func (p person) toString() string{
	return p.name + " " + p.lastName
}

func (p product) toString() string{
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x show) {
	fmt.Println(x.toString())
}

func main(){

	person1 := person{
		name:		"person",
		lastName:	"one",
	}

	product1 := product{
		name: 		"product",
		price:		10.0,
	}

	var thingPerson show = person(person1)

	var thingProduct show = product(product1)

	print(thingPerson)
	print(thingProduct)

}